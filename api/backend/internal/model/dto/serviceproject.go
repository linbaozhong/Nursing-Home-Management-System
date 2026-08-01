package dto

// ============ ServiceProjectController 请求 ============

// @request
// PageServiceByKeyQuery 分页查询服务请求
type PageServiceByKeyQuery struct {
	PageNum     *int    `json:"pageNum" valid:"required"` // 页码
	PageSize    *int    `json:"pageSize" valid:"required"` // 条数
	ServiceName *string `json:"serviceName"` // 服务名称
	TypeName    *string `json:"typeName"` // 服务类型名称
}

// @request
// OperateServiceTypeQuery 操作服务类型请求
type OperateServiceTypeQuery struct {
	ID   *int64  `json:"id"` // id
	Name *string `json:"name" valid:"required"` // 服务类型名称
}

// OperateServiceTypeQuery 操作服务类型请求（定义见 common.go）

// OperateServiceQuery 操作服务请求（定义见 common.go）

// ============ ServiceProjectController 响应 ============

// @response
// PageServiceByKeyVO 分页查询服务响应
type PageServiceByKeyVO struct {
	Rank
	ID           int64   `json:"id"` // id
	TypeName     string  `json:"typeName"` // 服务类型名称
	ServiceName  string  `json:"serviceName"` // 服务名称
	ChargeMethod string  `json:"chargeMethod"` // 收费方式
	Price        float64 `json:"price"` // 服务价格
	NeedDate     int     `json:"needDate"` // 所需时间
}