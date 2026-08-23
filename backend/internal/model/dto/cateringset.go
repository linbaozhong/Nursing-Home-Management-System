package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ CateringSetController 请求 ============

// PageCateringSetByKeyReq 分页查询餐饮套餐请求
// @request
type PageCateringSetByKeyReq struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	SetName    *string `json:"set_name"`                   // 套餐名称
	DishesName *string `json:"dishes_name"`                // 菜品名称
}

// OperateCateringSetReq 操作餐饮套餐请求
// @request
type OperateCateringSetReq struct {
	ID           *int64       `json:"id"`                              // id
	Name         *string      `json:"name" valid:"required"`           // 套餐名称
	MonthPrice   *types.Money `json:"month_price" valid:"required"`    // 月套餐费用
	DishesIDList []int64      `json:"dishes_id_list" valid:"required"` // 菜品编号列表
}

// ============ CateringSetController 响应 ============

// PageCateringSetByKeyResp 分页查询餐饮套餐响应
// @response
type PageCateringSetByKeyResp struct {
	ID         types.BigInt `json:"id"`          // id
	Name       string       `json:"name"`        // 套餐名称
	MonthPrice types.Money  `json:"month_price"` // 月套餐费用
}

// SetDishesResp 套餐菜品明细响应
// @response
type SetDishesResp struct {
	ID    types.BigInt `json:"id"`    // 菜品编号
	Name  string       `json:"name"`  // 菜品名称
	Price types.Money  `json:"price"` // 菜品价格
}

// GetCateringSetByIDResp 根据编号获取餐饮套餐响应（含菜品明细）
// @response
type GetCateringSetByIDResp struct {
	ID         types.BigInt    `json:"id"`          // id
	Name       string          `json:"name"`        // 套餐名称
	MonthPrice types.Money     `json:"month_price"` // 月套餐费用
	SetDishes  []SetDishesResp `json:"set_dishes"`  // 套餐菜品明细
}
