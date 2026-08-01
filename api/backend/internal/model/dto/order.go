package dto

// ============ OrderController 请求 ============

// @request
// PageOrderByKeyQuery 分页查询点餐请求
type PageOrderByKeyQuery struct {
	PageNum    *int    `json:"pageNum" valid:"required"` // 页码
	PageSize   *int    `json:"pageSize" valid:"required"` // 条数
	ElderName  *string `json:"elderName"` // 老人姓名
	ElderPhone *string `json:"elderPhone"` // 老人电话
}

// @request
// AddOrderQuery 新增点餐请求
type AddOrderQuery struct {
	ElderID         *int64                  `json:"elderId" valid:"required"` // 老人编号
	DineDate        *string                 `json:"dineDate" valid:"required"` // 就餐时间
	DineType        *string                 `json:"dineType" valid:"required"` // 就餐方式
	OrderDishesList []AddOrderDishesQuery  `json:"orderDishesList" valid:"required"` // 菜品列表
}

// @request
// AddOrderDishesQuery 新增点餐食物请求（嵌套）
type AddOrderDishesQuery struct {
	DishesID *int64 `json:"dishesId" valid:"required"` // 菜品编号
	OrderNum *int   `json:"orderNum" valid:"required"` // 菜品份数
}

// @request
// SendOrderQuery 送餐请求
type SendOrderQuery struct {
	ID               *int64  `json:"id" valid:"required"` // id
	StaffID          *int64  `json:"staffId" valid:"required"` // 送餐人员编号
	DeliverDishesDate *string `json:"deliverDishesDate" valid:"required"` // 送餐时间
}

// ============ OrderController 响应 ============

// @response
// PageOrderByKeyVO 分页查询点餐响应
type PageOrderByKeyVO struct {
	Rank
	ID               int64   `json:"id"` // id
	ElderName        string  `json:"elderName"` // 老人姓名
	ElderPhone       string  `json:"elderPhone"` // 老人电话
	DineDate         string  `json:"dineDate"` // 就餐时间
	DineType         string  `json:"dineType"` // 就餐方式
	StaffName        string  `json:"staffName"` // 送餐人员姓名
	DeliverDishesDate string `json:"deliverDishesDate"` // 送餐时间
	PayAmount        float64 `json:"payAmount"` // 支付总额
	OrderFlag        string  `json:"orderFlag"` // 订单状态
}

// @response
// GetOrderByIDVO 根据编号查询点餐响应
type GetOrderByIDVO struct {
	ElderName         string          `json:"elderName"` // 老人姓名
	ElderPhone        string          `json:"elderPhone"` // 老人电话
	DineDate          string          `json:"dineDate"` // 就餐时间
	DineType          string          `json:"dineType"` // 就餐方式
	StaffName         string          `json:"staffName"` // 送餐人员姓名
	DeliverDishesDate string          `json:"deliverDishesDate"` // 送餐时间
	OrderDishesVOList []OrderDishesVO `json:"orderDishesVoList"` // 订单菜品列表
}

// @response
// OrderDishesVO 菜单菜品响应（嵌套，继承 Rank）
type OrderDishesVO struct {
	Rank
	DishesName  string  `json:"dishesName"` // 菜品名称
	DishesPrice float64 `json:"dishesPrice"` // 菜品价格
	OrderNum    int     `json:"orderNum"` // 菜品份数
	SetFlag     string  `json:"setFlag"` // 套餐标记
	TotalAmount float64 `json:"totalAmount"` // 菜品总额
	ReallyAmount float64 `json:"reallyAmount"` // 实际总额
}