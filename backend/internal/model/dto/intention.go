package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ IntentionController 请求 ============

// PageIntentionByKeyReq 分页查询意向客户请求
// @request
type PageIntentionByKeyReq struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	ElderName  *string `json:"elder_name"`                 // 老人姓名
	ElderPhone *string `json:"elder_phone"`                // 老人电话
	LabelID    *int64  `json:"label_id"`                   // 标签编号
}

// OperateIntentionReq 操作意向客户请求
// @request
type OperateIntentionReq struct {
	ID      *int64  `json:"id"`                       // id
	Name    *string `json:"name" valid:"required"`    // 姓名
	IDNum   *string `json:"id_num" valid:"required"`  // 身份证号
	Age     *int    `json:"age" valid:"required"`     // 年龄
	Sex     *string `json:"sex" valid:"required"`     // 性别
	Phone   *string `json:"phone" valid:"required"`   // 电话
	Address *string `json:"address" valid:"required"` // 地址
}

// ============ IntentionController 响应 ============

// PageIntentionByKeyResp 分页查询意向客户响应
// @response
type PageIntentionByKeyResp struct {
	ID            types.BigInt         `json:"id"`              // 老人编号
	Name          string               `json:"name"`            // 老人姓名
	Phone         string               `json:"phone"`           // 老人电话
	Sex           string               `json:"sex"`             // 性别
	Age           int                  `json:"age"`             // 年龄
	LabelRespList []IntentionLabelResp `json:"label_resp_list"` // 客户标签
}

// IntentionLabelResp 意向客户标签（嵌套）
// @response
type IntentionLabelResp struct {
	Name  string `json:"name"`  // 名称
	Color string `json:"color"` // 颜色
}

// OperateIntentionResp 操作意向客户响应（继承 OperateIntentionReq）
// @response
type OperateIntentionResp struct {
	OperateIntentionReq
	IntentionLabelRespList []IntentionLabelResp `json:"intention_label_resp_list"` // 客户标签
}

// AddIntentReq 新增意向客户请求
// @request
type AddIntentReq struct {
	ID      *int64  `json:"id"`                       // id
	Name    *string `json:"name" valid:"required"`    // 姓名
	IDNum   *string `json:"id_num" valid:"required"`  // 身份证号
	Age     *int    `json:"age" valid:"required"`     // 年龄
	Sex     *string `json:"sex" valid:"required"`     // 性别
	Phone   *string `json:"phone" valid:"required"`   // 电话
	Address *string `json:"address" valid:"required"` // 地址
}

// EditIntentReq 编辑意向客户请求
// @request
type EditIntentReq struct {
	ID      *int64  `json:"id"`                       // id
	Name    *string `json:"name" valid:"required"`    // 姓名
	IDNum   *string `json:"id_num" valid:"required"`  // 身份证号
	Age     *int    `json:"age" valid:"required"`     // 年龄
	Sex     *string `json:"sex" valid:"required"`     // 性别
	Phone   *string `json:"phone" valid:"required"`   // 电话
	Address *string `json:"address" valid:"required"` // 地址
}
