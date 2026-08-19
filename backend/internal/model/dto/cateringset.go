package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ CateringSetController 请求 ============

// @request
// PageCateringSetByKeyQuery 分页查询餐饮套餐请求
type PageCateringSetByKeyQuery struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	SetName    *string `json:"set_name"`                   // 套餐名称
	DishesName *string `json:"dishes_name"`                // 菜品名称
}

// @request
// OperateCateringSetQuery 操作餐饮套餐请求
type OperateCateringSetQuery struct {
	ID           *int64       `json:"id"`                              // id
	Name         *string      `json:"name" valid:"required"`           // 套餐名称
	MonthPrice   *types.Money `json:"month_price" valid:"required"`    // 月套餐费用
	DishesIDList []int64      `json:"dishes_id_list" valid:"required"` // 菜品编号列表
}

// ============ CateringSetController 响应 ============

// @response
// PageCateringSetByKeyVO 分页查询餐饮套餐响应
type PageCateringSetByKeyVO struct {
	ID         int64       `json:"id"`          // id
	Name       string      `json:"name"`        // 套餐名称
	MonthPrice types.Money `json:"month_price"` // 月套餐费用
}

// @response
// SetDishesVO 套餐菜品明细响应
type SetDishesVO struct {
	ID    int64       `json:"id"`    // 菜品编号
	Name  string      `json:"name"`  // 菜品名称
	Price types.Money `json:"price"` // 菜品价格
}

// @response
// GetCateringSetByIDVO 根据编号获取餐饮套餐响应（含菜品明细）
type GetCateringSetByIDVO struct {
	ID         int64         `json:"id"`          // id
	Name       string        `json:"name"`        // 套餐名称
	MonthPrice types.Money   `json:"month_price"` // 月套餐费用
	SetDishes  []SetDishesVO `json:"set_dishes"`  // 套餐菜品明细
}
