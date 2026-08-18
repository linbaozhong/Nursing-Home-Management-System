package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ ServiceProjectController 请求 ============

// @request
// PageServiceByKeyQuery 分页查询服务请求
type PageServiceByKeyQuery struct {
	PageNum     *int    `json:"page_num" valid:"required"`  // 页码
	PageSize    *int    `json:"page_size" valid:"required"` // 条数
	ServiceName *string `json:"service_name"`               // 服务名称
	TypeName    *string `json:"type_name"`                  // 服务类型名称
}

// @request
// OperateServiceTypeQuery 操作服务类型请求
type OperateServiceTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 服务类型名称
}

// OperateServiceTypeQuery 操作服务类型请求（定义见 common.go）

// OperateServiceQuery 操作服务请求（定义见 common.go）

// ============ ServiceProjectController 响应 ============

// @response
// PageServiceByKeyVO 分页查询服务响应
type PageServiceByKeyVO struct {
	ID           int64       `json:"id"`            // id
	TypeName     string      `json:"type_name"`     // 服务类型名称
	ServiceName  string      `json:"service_name"`  // 服务名称
	ChargeMethod string      `json:"charge_method"` // 收费方式
	Price        types.Money `json:"price"`         // 服务价格
	NeedDate     int         `json:"need_date"`     // 所需时间
}
