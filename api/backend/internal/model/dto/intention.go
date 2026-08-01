package dto

// ============ IntentionController 请求 ============

// @request
// PageIntentionByKeyQuery 分页查询意向客户请求
type PageIntentionByKeyQuery struct {
	PageNum    *int    `json:"pageNum" valid:"required"` // 页码
	PageSize   *int    `json:"pageSize" valid:"required"` // 条数
	ElderName  *string `json:"elderName"` // 老人姓名
	ElderPhone *string `json:"elderPhone"` // 老人电话
	LabelID    *int64  `json:"labelId"` // 标签编号
}

// @request
// OperateIntentionQuery 操作意向客户请求
type OperateIntentionQuery struct {
	ID      *int64  `json:"id"` // id
	Name    *string `json:"name" valid:"required"` // 姓名
	IDNum   *string `json:"idNum" valid:"required"` // 身份证号
	Age     *int    `json:"age" valid:"required"` // 年龄
	Sex     *string `json:"sex" valid:"required"` // 性别
	Phone   *string `json:"phone" valid:"required"` // 电话
	Address *string `json:"address" valid:"required"` // 地址
}

// ============ IntentionController 响应 ============

// @response
// PageIntentionByKeyVO 分页查询意向客户响应
type PageIntentionByKeyVO struct {
	Rank
	ID          int64              `json:"id"` // 老人编号
	Name        string             `json:"name"` // 老人姓名
	Phone       string             `json:"phone"` // 老人电话
	Sex         string             `json:"sex"` // 性别
	Age         int                `json:"age"` // 年龄
	LabelVOList []IntentionLabelVO `json:"labelVoList"` // 客户标签
}

// @response
// IntentionLabelVO 意向客户标签（嵌套）
type IntentionLabelVO struct {
	Name  string `json:"name"` // 名称
	Color string `json:"color"` // 颜色
}

// @response
// OperateIntentionVO 操作意向客户响应（继承 OperateIntentionQuery）
type OperateIntentionVO struct {
	OperateIntentionQuery
}

// @request
// AddIntentQuery 新增意向客户请求
type AddIntentQuery struct {
	ID      *int64  `json:"id"` // id
	Name    *string `json:"name" valid:"required"` // 姓名
	IDNum   *string `json:"idNum" valid:"required"` // 身份证号
	Age     *int    `json:"age" valid:"required"` // 年龄
	Sex     *string `json:"sex" valid:"required"` // 性别
	Phone   *string `json:"phone" valid:"required"` // 电话
	Address *string `json:"address" valid:"required"` // 地址
}

// @request
// EditIntentQuery 编辑意向客户请求
type EditIntentQuery struct {
	ID      *int64  `json:"id"` // id
	Name    *string `json:"name" valid:"required"` // 姓名
	IDNum   *string `json:"idNum" valid:"required"` // 身份证号
	Age     *int    `json:"age" valid:"required"` // 年龄
	Sex     *string `json:"sex" valid:"required"` // 性别
	Phone   *string `json:"phone" valid:"required"` // 电话
	Address *string `json:"address" valid:"required"` // 地址
}