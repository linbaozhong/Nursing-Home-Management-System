package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ ServiceProjectController 请求 ============

// PageServiceByKeyReq 分页查询服务请求
// @request
type PageServiceByKeyReq struct {
	PageNum     *int    `json:"page_num" valid:"required"`  // 页码
	PageSize    *int    `json:"page_size" valid:"required"` // 条数
	ServiceName *string `json:"service_name"`               // 服务名称
	TypeName    *string `json:"type_name"`                  // 服务类型名称
}

// OperateServiceTypeReq 操作服务类型请求
// @request
type OperateServiceTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 服务类型名称
}

// OperateServiceTypeReq 操作服务类型请求（定义见 common.go）

// OperateServiceReq 操作服务请求（定义见 common.go）

// ============ ServiceProjectController 响应 ============

// PageServiceByKeyResp 分页查询服务响应
// @response
type PageServiceByKeyResp struct {
	ID           int64       `json:"id"`            // id
	TypeName     string      `json:"type_name"`     // 服务类型名称
	ServiceName  string      `json:"service_name"`  // 服务名称
	ChargeMethod string      `json:"charge_method"` // 收费方式
	Price        types.Money `json:"price"`         // 服务价格
	NeedDate     int         `json:"need_date"`     // 所需时间
}
