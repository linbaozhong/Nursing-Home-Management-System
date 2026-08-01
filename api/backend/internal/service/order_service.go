package service

import (
	"context"
	"errors"
	"time"

	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblorder"
	"api/internal/model/define/table/tblorderdishes"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

// parseTime 将 "yyyy-MM-dd HH:mm:ss" 或 "yyyy-MM-dd" 字符串解析为 types.Time
func parseTime(s string) types.Time {
	for _, layout := range []string{"2006-01-02 15:04:05", "2006-01-02"} {
		if t, e := time.Parse(layout, s); e == nil {
			return types.Time(t)
		}
	}
	return types.Time(time.Time{})
}

type order struct{}

var Order = &order{}

// PageOrderByKey 分页查询点餐列表（按老人姓名/电话关键字）
// 对应 Java: OrderServiceImpl.pageOrderByKey -> orderMapper.listOrderByKey
// SQL: SELECT o.*, e.elder_name, e.elder_phone, u.name AS staff_name
//
//	FROM `order` o
//	LEFT JOIN elder e ON e.id = o.elder_id
//	LEFT JOIN user u ON u.id = o.staff_id
//	WHERE (e.elder_name LIKE %key% OR e.elder_phone LIKE %key%) [可选]
//	ORDER BY o.create_time DESC;
//
// 再内存分页，并将 dine_date/deliver_dishes_date 截取为 yyyy-MM-dd，order_flag: NO=待支付 YES=已完成
//
// todo: 1) 用 dao.Order(db) 联表分页查询 2) 日期格式化 3) 状态替换 4) 组装 PageOrderByKeyVO 列表, 赋值 out
func (o *order) PageOrderByKey(ctx context.Context, in *dto.PageOrderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询并格式化
	return nil
}

// AddOrder 新增点餐（含套餐优惠计算 + 批量写入订单菜品）
// 对应 Java: OrderServiceImpl.addOrder @Transactional
// 逻辑：
//  1. 根据老人编号查出其套餐已含菜品列表(listSetDishesByElderId)
//  2. 对点餐菜品按 dishesId 分组，累计每份菜品 orderNum
//  3. 逐菜品：若属于套餐已含 -> reallyAmount = totalAmount - price（免一份），set_flag=YES；否则 reallyAmount = totalAmount，set_flag=NO
//  4. 累加 payAmount
//  5. 插入 order(pay_amount, order_flag=NO)，再批量插入 order_dishes(order_id 回填)
func (o *order) AddOrder(ctx context.Context, in *dto.AddOrderQuery, out *dto.EmptyResp) error {
	// 1) 老人套餐已含菜品列表
	// todo: setDishesList, e := dao.Dishes(db) ... listSetDishesByElderId(ctx, types.BigInt(in.ElderID)); 标记哪些 dishesId 已含
	// 2) 按 dishesId 分组累计份数
	numMap := map[int64]int{}
	for _, d := range in.OrderDishesList {
		numMap[d.DishesID] += d.OrderNum
	}
	// 3) 批量查菜品
	ids := make([]any, 0, len(numMap))
	for id := range numMap {
		ids = append(ids, types.BigInt(id))
	}
	dishesList, _, e := dao.Dishes(db).GetByIds(ctx, ids)
	if e != nil {
		return e
	}
	// 4) 计算金额并封装订单菜品
	var payAmount float64
	orderDishes := make([]*do.OrderDishes, 0, len(numMap))
	for _, dishes := range dishesList {
		num := numMap[dishes.Id]
		total := float64(num) * float64(dishes.Price)
		exist := false // todo: 判断该菜品是否在老人套餐已含列表中
		really := total
		setFlag := "N"
		if exist {
			really = total - float64(dishes.Price)
			setFlag = "Y"
		}
		payAmount += really
		orderDishes = append(orderDishes, &do.OrderDishes{
			DishesName:   dishes.DishesName,
			DishesPrice:  float64(dishes.Price),
			OrderNum:     int32(num),
			SetFlag:      setFlag,
			TotalAmount:  total,
			ReallyAmount: really,
		})
	}
	// 5) 插入订单
	bean := &do.Order{
		ElderId:  types.BigInt(in.ElderID),
		DineDate: parseTime(in.DineDate),
		DineType: in.DineType,
		PayAmount: payAmount,
		OrderFlag: "N",
	}
	_, e = dao.Order(db).InsertOne(ctx, bean)
	if e != nil {
		return e
	}
	// 6) 订单菜品回填 orderId 并批量插入
	for _, od := range orderDishes {
		od.OrderId = bean.Id
	}
	_, e = dao.OrderDishes(db).InsertBatch(ctx, orderDishes)
	if e != nil {
		return e
	}
	return nil
}

// GetOrderById 根据编号获取点餐详情（含订单菜品列表）
// 对应 Java: OrderServiceImpl.getOrderById -> orderMapper.getOrderById + orderDishesFunc.listOrderDishesByOrderId
// todo: 1) dao.Order(db).GetByID 查订单 2) dao.OrderDishes(db).List(Where order_id=id) 查菜品 3) 编序号 4) 组装 GetOrderByIDVO, 赋值 out
func (o *order) GetOrderById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Order(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("订单不存在")
	}
	_ = obj
	return nil
}

// SendOrder 送餐（标记完成 + 老人扣费 + 新增消费记录）
// 对应 Java: OrderServiceImpl.sendOrder @Transactional
// 逻辑：
//  1. 查订单，若 staff_id/deliver_dishes_date 已填 或 order_flag=YES -> 已送餐，报错(ORDER_SUCCESS)
//  2. 更新订单：staff_id/deliver_dishes_date/order_flag=YES
//  3. 老人扣费：elder.deductionFee(elderId, payAmount)  todo: 调 Elder 账户扣费
//  4. 新增消费记录：Consume.AddConsume(elderId, DISHES, payAmount, dineDate)
func (o *order) SendOrder(ctx context.Context, in *dto.SendOrderQuery, out *dto.EmptyResp) error {
	// 1) 查订单
	getOrder, has, e := dao.Order(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("订单不存在")
	}
	// 2) 校验是否已送餐
	if getOrder.StaffId != 0 || !getOrder.DeliverDishesDate.IsZero() || getOrder.OrderFlag == "Y" {
		return errors.New("订单已完成送餐") // 对应 Java ExceptionEnum.ORDER_SUCCESS
	}
	// 3) 更新订单为已送餐
	_, e = dao.Order(db).UpdateById(ctx, types.BigInt(in.ID),
		ace.Where(tblorder.StaffId.Set(types.BigInt(in.StaffID))),
		ace.Where(tblorder.DeliverDishesDate.Set(parseTime(in.DeliverDishesDate))),
		ace.Where(tblorder.OrderFlag.Set("Y")),
	)
	if e != nil {
		return e
	}
	// 4) 老人扣费：从老人账户余额扣除 payAmount
	// todo: elderFunc.deductionFee(getOrder.ElderId, getOrder.PayAmount) 调用 Elder 账户扣费逻辑(查 elder 表减余额)
	// 5) 新增消费记录（DISHES 类型，对应 dineDate）
	_, e = dao.Consume(db).InsertOne(ctx, &do.Consume{
		ElderId:      getOrder.ElderId,
		ConsumeType:  "DISHES",
		ConsumeAmount: getOrder.PayAmount,
		ConsumeDate:  getOrder.DineDate,
	})
	if e != nil {
		return e
	}
	return nil
}
