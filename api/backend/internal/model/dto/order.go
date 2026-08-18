package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ OrderController 请求 ============

// @request
// PageOrderByKeyQuery 分页查询点餐请求
type PageOrderByKeyQuery struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	ElderName  *string `json:"elder_name"`                 // 老人姓名
	ElderPhone *string `json:"elder_phone"`                // 老人电话
}

// @request
// AddOrderQuery 新增点餐请求
type AddOrderQuery struct {
	ElderID         *int64                `json:"elder_id" valid:"required"`          // 老人编号
	DineDate        *time.Time            `json:"dine_date" valid:"required"`         // 就餐时间
	DineType        *string               `json:"dine_type" valid:"required"`         // 就餐方式
	OrderDishesList []AddOrderDishesQuery `json:"order_dishes_list" valid:"required"` // 菜品列表
}

// @request
// AddOrderDishesQuery 新增点餐食物请求（嵌套）
type AddOrderDishesQuery struct {
	DishesID *int64 `json:"dishes_id" valid:"required"` // 菜品编号
	OrderNum *int   `json:"order_num" valid:"required"` // 菜品份数
}

// @request
// SendOrderQuery 送餐请求
type SendOrderQuery struct {
	ID                *int64     `json:"id" valid:"required"`                  // id
	StaffID           *int64     `json:"staff_id" valid:"required"`            // 送餐人员编号
	DeliverDishesDate *time.Time `json:"deliver_dishes_date" valid:"required"` // 送餐时间
}

// ============ OrderController 响应 ============

// @response
// PageOrderByKeyVO 分页查询点餐响应
type PageOrderByKeyVO struct {
	ID                int64       `json:"id"`                  // id
	ElderName         string      `json:"elder_name"`          // 老人姓名
	ElderPhone        string      `json:"elder_phone"`         // 老人电话
	DineDate          time.Time   `json:"dine_date"`           // 就餐时间
	DineType          string      `json:"dine_type"`           // 就餐方式
	StaffName         string      `json:"staff_name"`          // 送餐人员姓名
	DeliverDishesDate time.Time   `json:"deliver_dishes_date"` // 送餐时间
	PayAmount         types.Money `json:"pay_amount"`          // 支付总额
	OrderFlag         string      `json:"order_flag"`          // 订单状态
}

// @response
// GetOrderByIDVO 根据编号查询点餐响应
type GetOrderByIDVO struct {
	ElderName         string          `json:"elder_name"`           // 老人姓名
	ElderPhone        string          `json:"elder_phone"`          // 老人电话
	DineDate          time.Time       `json:"dine_date"`            // 就餐时间
	DineType          string          `json:"dine_type"`            // 就餐方式
	StaffName         string          `json:"staff_name"`           // 送餐人员姓名
	DeliverDishesDate time.Time       `json:"deliver_dishes_date"`  // 送餐时间
	OrderDishesVOList []OrderDishesVO `json:"order_dishes_vo_list"` // 订单菜品列表
}

// @response
// OrderDishesVO 菜单菜品响应（嵌套，继承 Rank）
type OrderDishesVO struct {
	DishesName   string      `json:"dishes_name"`   // 菜品名称
	DishesPrice  types.Money `json:"dishes_price"`  // 菜品价格
	OrderNum     int         `json:"order_num"`     // 菜品份数
	SetFlag      string      `json:"set_flag"`      // 套餐标记
	TotalAmount  types.Money `json:"total_amount"`  // 菜品总额
	ReallyAmount types.Money `json:"really_amount"` // 实际总额
}
