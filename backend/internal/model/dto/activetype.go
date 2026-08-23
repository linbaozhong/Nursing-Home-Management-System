package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ ActiveTypeController 请求 ============

// PageActiveTypeByKeyReq 分页查询活动分类请求
// @request
type PageActiveTypeByKeyReq struct {
	PageNum        *int    `json:"page_num" valid:"required"`  // 页码
	PageSize       *int    `json:"page_size" valid:"required"` // 条数
	ActiveTypeName *string `json:"active_type_name"`           // 活动分类名称
}

// OperateActiveTypeReq 操作活动分类请求
// @request
type OperateActiveTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 活动分类名称
}

// AddActiveTypeReq 新增活动分类请求
// @request
type AddActiveTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 活动分类名称
}

// ============ ActiveTypeController 响应 ============

// PageActiveTypeByKeyResp 分页查询活动分类响应
// @response
type PageActiveTypeByKeyResp struct {
	ID   types.BigInt `json:"id"`   // id
	Name string       `json:"name"` // 活动分类名称
}

// OperateActiveTypeResp 操作活动分类响应（继承 OperateActiveTypeReq）
// @response
type OperateActiveTypeResp struct {
	OperateActiveTypeReq
}
