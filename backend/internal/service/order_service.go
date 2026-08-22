package service

import (
	"context"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tbldishes"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblorder"
	"api/internal/model/define/table/tblorderdishes"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var _ = (*orderService)(nil)

type orderService struct{}

// orderJoin 接收点餐联表（老人姓名/电话、送餐人姓名）查询结果的中间结构体
type orderJoin struct {
	ID                types.BigInt `json:"id"`
	ElderName         types.String `json:"elder_name"`
	ElderPhone        types.String `json:"elder_phone"`
	DineDate          types.Time   `json:"dine_date"`
	DineType          types.String `json:"dine_type"`
	StaffName         types.String `json:"staff_name"`
	DeliverDishesDate types.Time   `json:"deliver_dishes_date"`
	PayAmount         types.Money  `json:"pay_amount"`
	OrderFlag         types.Int8   `json:"order_flag"`
}

// PageOrderByKey 分页查询点餐
func (s *orderService) PageOrderByKey(ctx context.Context, in *dto.PageOrderByKeyReq, out *[]dto.PageOrderByKeyResp) error {
	clampPage(in.PageNum, in.PageSize)
	q := db.Table(tblorder.TableName).
		LeftJoin(tblorder.ElderId, tblelder.Id).
		LeftJoin(tblorder.StaffId, tblstaff.Id).
		Where(tblorder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.ElderName != nil && *in.ElderName != "" {
		q = q.Where(tblelder.Name.Like(*in.ElderName))
	}
	if in.ElderPhone != nil && *in.ElderPhone != "" {
		q = q.And(tblelder.Phone.Like(*in.ElderPhone))
	}
	var joins []orderJoin
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblorder.Id,
			tblorder.DineDate,
			tblorder.DineType,
			tblorder.DeliverDishesDate,
			tblorder.PayAmount,
			tblorder.OrderFlag,
			tblelder.Name.As("elder_name"),
			tblelder.Phone.As("elder_phone"),
			tblstaff.Name.As("staff_name"),
		).
		Desc(tblorder.Id).
		Select().Gets(ctx, &joins)
	if e != nil {
		return e
	}

	res := make([]dto.PageOrderByKeyResp, 0, len(joins))
	for _, j := range joins {
		res = append(res, dto.PageOrderByKeyResp{
			ID:                int64(j.ID),
			ElderName:         j.ElderName.String(),
			ElderPhone:        j.ElderPhone.String(),
			DineDate:          j.DineDate.Time,
			DineType:          j.DineType.String(),
			StaffName:         j.StaffName.String(),
			DeliverDishesDate: j.DeliverDishesDate.Time,
			PayAmount:         j.PayAmount,
			OrderFlag:         constant.YesNo(j.OrderFlag).String(),
		})
	}
	*out = res
	return nil
}

// GetOrderById 查询点餐详情（含菜品）
func (s *orderService) GetOrderById(ctx context.Context, in *dto.IDReq, out *dto.GetOrderByIDResp) error {
	order, has, e := dao.Order(db).Get(ctx, ace.Where(tblorder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblorder.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.DineDate = order.DineDate.Time
	out.DineType = order.DineType.String()
	out.DeliverDishesDate = order.DeliverDishesDate.Time
	// 老人信息
	if el, eh, ee := dao.Elder(db).Get(ctx, ace.Where(tblelder.Id.Eq(order.ElderId))); ee == nil && eh {
		out.ElderName = el.Name.String()
		out.ElderPhone = el.Phone.String()
	}
	// 送餐人
	if st, sh, se := dao.Staff(db).Get(ctx, ace.Where(tblstaff.Id.Eq(order.StaffId))); se == nil && sh {
		out.StaffName = st.Name.String()
	}
	// 菜品明细
	dishes, _, de := dao.OrderDishes(db).List(ctx, ace.Where(tblorderdishes.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblorderdishes.OrderId.Eq(order.Id)))
	if de == nil {
		for _, d := range dishes {
			out.OrderDishesRespList = append(out.OrderDishesRespList, dto.OrderDishesResp{
				DishesName:   d.DishesName.String(),
				DishesPrice:  d.DishesPrice,
				OrderNum:     int(d.OrderNum),
				SetFlag:      constant.YesNo(d.SetFlag).String(),
				TotalAmount:  d.TotalAmount,
				ReallyAmount: d.ReallyAmount,
			})
		}
	}
	return nil
}

// AddOrder 新增点餐（含菜品明细与金额计算）
func (s *orderService) AddOrder(ctx context.Context, in *dto.AddOrderReq, out *dto.EmptyResp) error {
	if in.ElderID == nil || len(in.OrderDishesList) == 0 {
		return constant.ErrParamInvalid
	}
	order := &do.Order{
		TenantId:  types.BigInt(lib.TenantID(ctx)),
		DineType:  types.String(orEmpty(in.DineType)),
		ElderId:   types.BigInt(*in.ElderID),
		PayAmount: types.Money(0),
		DineDate:  types.Time{*in.DineDate},
		OrderFlag: types.Int8(constant.YesNoNo),
		CreateId:  types.BigInt(*in.ElderID),
	}
	if _, e := dao.Order(db).InsertOne(ctx, order); e != nil {
		return e
	}
	var payAmount types.Money
	for _, d := range in.OrderDishesList {
		if d.DishesID == nil || d.OrderNum == nil {
			continue
		}
		// 套餐菜减价、普通菜原价（Go 端菜品无套餐标记，统一按原价计算）
		var price types.Money
		dishName := ""
		if dish, dh, derr := dao.Dishes(db).Get(ctx, ace.Where(tbldishes.Id.Eq(types.BigInt(*d.DishesID)))); derr == nil && dh {
			price = dish.Price
			dishName = dish.Name.String()
		}
		num := types.Money(*d.OrderNum)
		reallyPrice := price * num
		payAmount += reallyPrice
		od := &do.OrderDishes{
			TenantId:     types.BigInt(lib.TenantID(ctx)),
			OrderId:      order.Id,
			DishesName:   types.String(dishName),
			DishesPrice:  types.Money(price),
			OrderNum:     types.Int32(int32(num)),
			TotalAmount:  types.Money(price * num),
			ReallyAmount: types.Money(reallyPrice),
		}
		if _, e := dao.OrderDishes(db).InsertOne(ctx, od); e != nil {
			return e
		}
	}
	if _, e := dao.Order(db).UpdateById(ctx, order.Id, tblorder.PayAmount.Set(types.Float64(payAmount))); e != nil {
		return e
	}
	return nil
}

// SendOrder 送餐（标记完成并扣费）
func (s *orderService) SendOrder(ctx context.Context, in *dto.SendOrderReq, out *dto.EmptyResp) error {
	if in.ID == nil || in.StaffID == nil {
		return constant.ErrParamInvalid
	}
	order, has, e := dao.Order(db).Get(ctx, ace.Where(tblorder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblorder.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	if _, e = dao.Order(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblorder.OrderFlag.Set(types.Int8(constant.YesNoYes)),
		tblorder.StaffId.Set(types.BigInt(*in.StaffID)),
		tblorder.DeliverDishesDate.Set(types.Time{Time: *in.DeliverDishesDate}),
	); e != nil {
		return e
	}
	// 扣减老人余额并记录消费
	if el, eh, ee := dao.Elder(db).Get(ctx, ace.Where(tblelder.Id.Eq(order.ElderId))); ee == nil && eh {
		newBalance := el.Balance - order.PayAmount
		if newBalance < 0 {
			newBalance = 0
		}
		_, _ = dao.Elder(db).UpdateById(ctx, order.ElderId, tblelder.Balance.Set(types.Float64(newBalance)))
		consume := &do.Consume{
			TenantId:      types.BigInt(lib.TenantID(ctx)),
			ElderId:       order.ElderId,
			ConsumeType:   types.String("点餐"),
			ConsumeAmount: order.PayAmount,
			CreateId:      types.BigInt(*in.StaffID),
		}
		_, _ = dao.Consume(db).InsertOne(ctx, consume)
	}
	return nil
}

func orInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}
