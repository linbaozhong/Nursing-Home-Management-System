package dto

import "time"

// ============ RetreatApplyController 请求 ============

// PageRetreatAuditReq 分页查询退住审核请求（RetreatApply 复用）
// @request
type PageRetreatAuditReq struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	ElderSex  *string `json:"elder_sex"`                  // 老人性别
	IDNum     *string `json:"id_num"`                     // 老人身份证号
}

// PageRetreatApplyReq 分页查询退住申请请求（继承 PageRetreatAuditReq）
// @request
type PageRetreatApplyReq struct {
	PageRetreatAuditReq
	BedName *string `json:"bed_name"` // 床位名称
}

// GetReserveByReserveIDAndElderIDReq 根据预定/老人编号获取请求（定义见 reserve.go）

// PageRetreatApplyByKeyReq 分页查询退住申请请求
// @request
type PageRetreatApplyByKeyReq struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	ElderSex  *string `json:"elder_sex"`                  // 老人性别
	IDNum     *string `json:"id_num"`                     // 老人身份证号
}

// AddRetreatApplyReq 新增退住申请请求
// @request
type AddRetreatApplyReq struct {
	ElderID     *int64     `json:"elder_id" valid:"required"`     // 老人编号
	ApplyReason *string    `json:"apply_reason" valid:"required"` // 退住原因
	ApplyDate   *time.Time `json:"apply_date" valid:"required"`   // 申请日期
}

// EditRetreatApplyReq 编辑退住申请请求
// @request
type EditRetreatApplyReq struct {
	ID          *int64     `json:"id"`                            // id
	ElderID     *int64     `json:"elder_id" valid:"required"`     // 老人编号
	ApplyReason *string    `json:"apply_reason" valid:"required"` // 退住原因
	ApplyDate   *time.Time `json:"apply_date" valid:"required"`   // 申请日期
}

// PageSearchElderByKeyReq 分页搜索老人请求（定义见 elderrecord.go）

// ============ RetreatApplyController 响应 ============

// PageRetreatByKeyResp 分页查询退住响应（RetreatApply / RetreatAudit 共用）
// @response
type PageRetreatByKeyResp struct {
	ApplyID   int64  `json:"apply_id"`   // 申请编号
	ElderID   int64  `json:"elder_id"`   // 老人编号
	ElderName string `json:"elder_name"` // 老人姓名
	ElderSex  string `json:"elder_sex"`  // 老人性别
	IDNum     string `json:"id_num"`     // 身份证号
	BedName   string `json:"bed_name"`   // 床位名称
	ApplyFlag *int8  `json:"apply_flag"` // 审核状态（0-待审核 1-审核中 2-通过 -1-不通过）
}
