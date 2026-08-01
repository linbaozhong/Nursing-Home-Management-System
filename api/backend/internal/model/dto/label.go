package dto

// ============ LabelController 请求 ============

// @request
// OperateLabelQuery 操作客户标签请求
type OperateLabelQuery struct {
	ID     *int64  `json:"id"` // id
	TypeID *int64  `json:"typeId"` // typeId
	Name   *string `json:"name" valid:"required"` // 名称
	Color  *string `json:"color" valid:"required"` // 颜色
}

// @request
// OperateLabelTypeQuery 操作客户标签分类请求
type OperateLabelTypeQuery struct {
	ID   *int64  `json:"id"` // id
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
	ID     *int64  `json:"id"` // id
	TypeID *int64  `json:"typeId"` // typeId
	Name   *string `json:"name" valid:"required"` // 名称
	Color  *string `json:"color" valid:"required"` // 颜色
}

// @request
// EditLabelQuery 编辑客户标签请求
type EditLabelQuery struct {
	ID     *int64  `json:"id"` // id
	TypeID *int64  `json:"typeId"` // typeId
	Name   *string `json:"name" valid:"required"` // 名称
	Color  *string `json:"color" valid:"required"` // 颜色
}

// @request
// PageLabelTypeByKeyQuery 分页查询客户标签分类请求
type PageLabelTypeByKeyQuery struct {
	PageNum  *int    `json:"pageNum" valid:"required"` // 页码
	PageSize *int    `json:"pageSize" valid:"required"` // 条数
	Name     *string `json:"name"` // 分类名称
}

// @request
// AddLabelTypeQuery 新增客户标签分类请求
type AddLabelTypeQuery struct {
	ID   *int64  `json:"id"` // id
	Name *string `json:"name" valid:"required"` // 名称
}

// @request
// EditLabelTypeQuery 编辑客户标签分类请求
type EditLabelTypeQuery struct {
	ID   *int64  `json:"id"` // id
	Name *string `json:"name" valid:"required"` // 名称
}

// ListLabelVO 客户标签分类列表响应（定义见 common.go）