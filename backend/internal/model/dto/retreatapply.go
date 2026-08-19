package dto

import "time"

// ============ RetreatApplyController 请求 ============

// @request
// PageRetreatAuditQuery 分页查询退住审核请求（RetreatApply 复用）
type PageRetreatAuditQuery struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	ElderSex  *string `json:"elder_sex"`                  // 老人性别
	IDNum     *string `json:"id_num"`                     // 老人身份证号
}

// @request
// PageRetreatApplyQuery 分页查询退住申请请求（继承 PageRetreatAuditQuery）
type PageRetreatApplyQuery struct {
	PageRetreatAuditQuery
	BedName *string `json:"bed_name"` // 床位名称
}

// GetReserveByReserveIDAndElderIDQuery 根据预定/老人编号获取请求（定义见 reserve.go）

// @request
// PageRetreatApplyByKeyQuery 分页查询退住申请请求
type PageRetreatApplyByKeyQuery struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	ElderSex  *string `json:"elder_sex"`                  // 老人性别
	IDNum     *string `json:"id_num"`                     // 老人身份证号
}

// @request
// AddRetreatApplyQuery 新增退住申请请求
type AddRetreatApplyQuery struct {
	ElderID     *int64     `json:"elder_id" valid:"required"`     // 老人编号
	ApplyReason *string    `json:"apply_reason" valid:"required"` // 退住原因
	ApplyDate   *time.Time `json:"apply_date" valid:"required"`   // 申请日期
}

// @request
// EditRetreatApplyQuery 编辑退住申请请求
type EditRetreatApplyQuery struct {
	ID          *int64     `json:"id"`                            // id
	ElderID     *int64     `json:"elder_id" valid:"required"`     // 老人编号
	ApplyReason *string    `json:"apply_reason" valid:"required"` // 退住原因
	ApplyDate   *time.Time `json:"apply_date" valid:"required"`   // 申请日期
}

// PageSearchElderByKeyQuery 分页搜索老人请求（定义见 elderrecord.go）

// ============ RetreatApplyController 响应 ============

// @response
// PageRetreatByKeyVO 分页查询退住响应（RetreatApply / RetreatAudit 共用）
type PageRetreatByKeyVO struct {
	ApplyID   int64  `json:"apply_id"`   // 申请编号
	ElderID   int64  `json:"elder_id"`   // 老人编号
	ElderName string `json:"elder_name"` // 老人姓名
	ElderSex  string `json:"elder_sex"`  // 老人性别
	IDNum     string `json:"id_num"`     // 身份证号
	BedName   string `json:"bed_name"`   // 床位名称
	ApplyFlag *int8  `json:"apply_flag"` // 审核状态（0-待审核 1-审核中 2-通过 -1-不通过）
}

// @response
// PageSearchElderByKeyVO 分页搜索老人响应（供退住申请选择老人）
type PageSearchElderByKeyVO struct {
	ElderID   int64  `json:"elder_id"`   // 老人编号
	ElderName string `json:"elder_name"` // 老人姓名
	ElderSex  string `json:"elder_sex"`  // 老人性别
	IDNum     string `json:"id_num"`     // 身份证号
	BedName   string `json:"bed_name"`   // 床位名称
}
