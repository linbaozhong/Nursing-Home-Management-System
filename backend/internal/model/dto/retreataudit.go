package dto

// ============ RetreatAuditController 请求 ============

// PageRetreatAuditByKeyReq 分页查询退住审核请求
// @request
type PageRetreatAuditByKeyReq struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	ElderSex  *string `json:"elder_sex"`                  // 老人性别
	IDNum     *string `json:"id_num"`                     // 老人身份证号
}

// AuditRetreatReq 审核退住请求
// @request
type AuditRetreatReq struct {
	ID          *int64  `json:"id" valid:"required"`           // 退住申请编号
	AuditResult *string `json:"audit_result" valid:"required"` // 审核结果
	AuditRemark *string `json:"audit_remark"`                  // 审核备注
}

// PageRetreatAuditReq 分页查询退住审核请求（定义见 retreatapply.go）

// AuditElderFeeReq 审核老人费用详情请求（定义见 elderrecord.go）

// ============ RetreatAuditController 响应 ============

// GetElderFeeByIDResp 根据编号获取老人费用详情响应（定义见 common.go）
