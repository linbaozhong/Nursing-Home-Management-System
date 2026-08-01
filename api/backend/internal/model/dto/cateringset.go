package dto

// ============ CateringSetController 请求 ============

// @request
// PageCateringSetByKeyQuery 分页查询餐饮套餐请求
type PageCateringSetByKeyQuery struct {
	PageNum    *int    `json:"pageNum" valid:"required"` // 页码
	PageSize   *int    `json:"pageSize" valid:"required"` // 条数
	SetName    *string `json:"setName"` // 套餐名称
	DishesName *string `json:"dishesName"` // 菜品名称
}

// @request
// OperateCateringSetQuery 操作餐饮套餐请求
type OperateCateringSetQuery struct {
	ID            *int64   `json:"id"` // id
	Name          *string  `json:"name" valid:"required"` // 套餐名称
	MonthPrice    *float64 `json:"monthPrice" valid:"required"` // 月套餐费用
	DishesIDList  []int64 `json:"dishesIdList" valid:"required"` // 菜品编号列表
}

// ============ CateringSetController 响应 ============

// @response
// PageCateringSetByKeyVO 分页查询餐饮套餐响应
type PageCateringSetByKeyVO struct {
	Rank
	ID        int64   `json:"id"` // id
	Name      string  `json:"name"` // 套餐名称
	MonthPrice float64 `json:"monthPrice"` // 月套餐费用
}