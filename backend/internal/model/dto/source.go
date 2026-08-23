package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ SourceController 请求 ============

// PageSourceByKeyReq 分页查询来源渠道请求
// @request
type PageSourceByKeyReq struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	SourceName *string `json:"source_name"`                // 来源渠道名称
}

// OperateSourceReq 操作来源渠道请求
// @request
type OperateSourceReq struct {
	ID   *int64  `json:"id" valid:"required"`   // id
	Name *string `json:"name" valid:"required"` // 来源渠道名称
}

// ============ SourceController 响应 ============

// PageSourceByKeyResp 分页查询来源渠道响应
// @response
type PageSourceByKeyResp struct {
	ID   types.BigInt `json:"id"`   // id
	Name string       `json:"name"` // 来源渠道名称
}

// OperateSourceResp 操作来源渠道响应
// @response
type OperateSourceResp struct {
	OperateSourceReq
}
