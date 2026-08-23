package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ LabelController 请求 ============

// OperateLabelReq 操作客户标签请求
// @request
type OperateLabelReq struct {
	ID     *int64  `json:"id"`                     // id
	TypeID *int64  `json:"type_id"`                // typeId
	Name   *string `json:"name" valid:"required"`  // 名称
	Color  *string `json:"color" valid:"required"` // 颜色
}

// OperateLabelTypeReq 操作客户标签分类请求
// @request
type OperateLabelTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 名称
}

// ============ LabelController 响应 ============

// OperateLabelResp 操作客户标签响应（继承 OperateLabelReq）
// @response
type OperateLabelResp struct {
	OperateLabelReq
}

// OperateLabelTypeResp 操作客户标签分类响应（继承 OperateLabelTypeReq）
// @response
type OperateLabelTypeResp struct {
	OperateLabelTypeReq
}

// AddLabelReq 新增客户标签请求
// @request
type AddLabelReq struct {
	ID     *int64  `json:"id"`                     // id
	TypeID *int64  `json:"type_id"`                // typeId
	Name   *string `json:"name" valid:"required"`  // 名称
	Color  *string `json:"color" valid:"required"` // 颜色
}

// EditLabelReq 编辑客户标签请求
// @request
type EditLabelReq struct {
	ID     *int64  `json:"id"`                     // id
	TypeID *int64  `json:"type_id"`                // typeId
	Name   *string `json:"name" valid:"required"`  // 名称
	Color  *string `json:"color" valid:"required"` // 颜色
}

// PageLabelTypeByKeyReq 分页查询客户标签分类请求
// @request
type PageLabelTypeByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 分类名称
}

// AddLabelTypeReq 新增客户标签分类请求
// @request
type AddLabelTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 名称
}

// EditLabelTypeReq 编辑客户标签分类请求
// @request
type EditLabelTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 名称
}

// PageLabelByKeyResp 分页查询标签响应
// @response
type PageLabelByKeyResp struct {
	ID       types.BigInt `json:"id"`        // 标签编号
	Name     string       `json:"name"`      // 标签名称
	Color    string       `json:"color"`     // 标签颜色
	TypeID   types.BigInt `json:"type_id"`   // 标签分类编号
	TypeName string       `json:"type_name"` // 标签分类名称
}

// PageLabelTypeResp 分页查询标签分类响应
// @response
type PageLabelTypeResp struct {
	ID         types.BigInt `json:"id"`          // 标签分类编号
	Name       string       `json:"name"`        // 标签分类名称
	CreateTime time.Time    `json:"create_time"` // 创建时间
	UpdateTime time.Time    `json:"update_time"` // 更新时间
}

// ListLabelResp 客户标签分类列表响应（定义见 common.go）
