package dto

// ============ RetreatApplyController 请求 ============

// @request
// PageRetreatAuditQuery 分页查询退住审核请求（RetreatApply 复用）
type PageRetreatAuditQuery struct {
	PageNum   *int    `json:"pageNum" valid:"required"` // 页码
	PageSize  *int    `json:"pageSize" valid:"required"` // 条数
	ElderName *string `json:"elderName"` // 老人姓名
	ElderSex  *string `json:"elderSex"` // 老人性别
	IDNum     *string `json:"idNum"` // 老人身份证号
}

// @request
// PageRetreatApplyQuery 分页查询退住申请请求（继承 PageRetreatAuditQuery）
type PageRetreatApplyQuery struct {
	PageRetreatAuditQuery
	BedName *string `json:"bedName"` // 床位名称
}

// GetReserveByReserveIDAndElderIDQuery 根据预定/老人编号获取请求（定义见 reserve.go）

// @request
// PageRetreatApplyByKeyQuery 分页查询退住申请请求
type PageRetreatApplyByKeyQuery struct {
	PageNum   *int    `json:"pageNum" valid:"required"` // 页码
	PageSize  *int    `json:"pageSize" valid:"required"` // 条数
	ElderName *string `json:"elderName"` // 老人姓名
	ElderSex  *string `json:"elderSex"` // 老人性别
	IDNum     *string `json:"idNum"` // 老人身份证号
}

// @request
// AddRetreatApplyQuery 新增退住申请请求
type AddRetreatApplyQuery struct {
	ElderID     *int64  `json:"elderId" valid:"required"` // 老人编号
	ApplyReason *string `json:"applyReason" valid:"required"` // 退住原因
	ApplyDate   *string `json:"applyDate" valid:"required"` // 申请日期
}

// @request
// EditRetreatApplyQuery 编辑退住申请请求
type EditRetreatApplyQuery struct {
	ID          *int64  `json:"id"` // id
	ElderID     *int64  `json:"elderId" valid:"required"` // 老人编号
	ApplyReason *string `json:"applyReason" valid:"required"` // 退住原因
	ApplyDate   *string `json:"applyDate" valid:"required"` // 申请日期
}

// PageSearchElderByKeyQuery 分页搜索老人请求（定义见 elderrecord.go）

// ============ RetreatApplyController 响应 ============

// @response
// PageRetreatByKeyVO 分页查询退住响应（RetreatApply / RetreatAudit 共用）
type PageRetreatByKeyVO struct {
	Rank
	ApplyID   int64  `json:"applyId"` // 申请编号
	ElderID   int64  `json:"elderId"` // 老人编号
	ElderName string `json:"elderName"` // 老人姓名
	ElderSex  string `json:"elderSex"` // 老人性别
	IDNum     string `json:"idNum"` // 身份证号
	BedName   string `json:"bedName"` // 床位名称
	ApplyFlag string `json:"applyFlag"` // 审核状态
}