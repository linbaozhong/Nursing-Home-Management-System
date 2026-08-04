package dto

import "time"

// ============ LabelController 请求 ============

// @request
// OperateLabelQuery 操作客户标签请求
type OperateLabelQuery struct {
	ID     *int64  `json:"id"`                     // id
	TypeID *int64  `json:"type_id"`                // typeId
	Name   *string `json:"name" valid:"required"`  // 名称
	Color  *string `json:"color" valid:"required"` // 颜色
}

// @request
// OperateLabelTypeQuery 操作客户标签分类请求
type OperateLabelTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 名称
}

// ============ LabelController 响应 ============

// @response
// OperateLabelVO 操作客户标签响应（继承 OperateLabelQuery）
type OperateLabelVO struct {
	OperateLabelQuery
}

// @response
// OperateLabelTypeVO 操作客户标签分类响应（继承 OperateLabelTypeQuery）
type OperateLabelTypeVO struct {
	OperateLabelTypeQuery
}

// @request
// AddLabelQuery 新增客户标签请求
type AddLabelQuery struct {
	ID     *int64  `json:"id"`                     // id
	TypeID *int64  `json:"type_id"`                // typeId
	Name   *string `json:"name" valid:"required"`  // 名称
	Color  *string `json:"color" valid:"required"` // 颜色
}

// @request
// EditLabelQuery 编辑客户标签请求
type EditLabelQuery struct {
	ID     *int64  `json:"id"`                     // id
	TypeID *int64  `json:"type_id"`                // typeId
	Name   *string `json:"name" valid:"required"`  // 名称
	Color  *string `json:"color" valid:"required"` // 颜色
}

// @request
// PageLabelTypeByKeyQuery 分页查询客户标签分类请求
type PageLabelTypeByKeyQuery struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 分类名称
}

// @request
// AddLabelTypeQuery 新增客户标签分类请求
type AddLabelTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 名称
}

// @request
// EditLabelTypeQuery 编辑客户标签分类请求
type EditLabelTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 名称
}

// @response
// PageLabelByKeyVO 分页查询标签响应
type PageLabelByKeyVO struct {
	ID       int64  `json:"id"`        // 标签编号
	Name     string `json:"name"`      // 标签名称
	Color    string `json:"color"`     // 标签颜色
	TypeID   int64  `json:"type_id"`   // 标签分类编号
	TypeName string `json:"type_name"` // 标签分类名称
}

// @response
// PageLabelTypeVO 分页查询标签分类响应
type PageLabelTypeVO struct {
	ID         int64     `json:"id"`          // 标签分类编号
	Name       string    `json:"name"`        // 标签分类名称
	CreateTime time.Time `json:"create_time"` // 创建时间
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

// ListLabelVO 客户标签分类列表响应（定义见 common.go）
