package dto

// ============ RetreatAuditController 请求 ============

// @request
// PageRetreatAuditByKeyQuery 分页查询退住审核请求
type PageRetreatAuditByKeyQuery struct {
	PageNum   *int    `json:"pageNum" valid:"required"` // 页码
	PageSize  *int    `json:"pageSize" valid:"required"` // 条数
	ElderName *string `json:"elderName"` // 老人姓名
	ElderSex  *string `json:"elderSex"` // 老人性别
	IDNum     *string `json:"idNum"` // 老人身份证号
}

// @request
// AuditRetreatQuery 审核退住请求
type AuditRetreatQuery struct {
	ID         *int64  `json:"id" valid:"required"` // 退住申请编号
	AuditResult *string `json:"auditResult" valid:"required"` // 审核结果
	AuditRemark *string `json:"auditRemark"` // 审核备注
}

// PageRetreatAuditQuery 分页查询退住审核请求（定义见 retreatapply.go）

// AuditElderFeeQuery 审核老人费用详情请求（定义见 elderrecord.go）

// ============ RetreatAuditController 响应 ============

// GetElderFeeByIDVO 根据编号获取老人费用详情响应（定义见 common.go）