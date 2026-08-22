package dto

import "time"

// PageVisitPlanReq 分页查询回访计划（Java: PageVisitPlanReq）
// @request
type PageVisitPlanReq struct {
	PageNum      *int   `json:"page_num"`      // 页码
	PageSize     *int   `json:"page_size"`     // 每页数量
	ElderID      *int64 `json:"elder_id"`      // 老人编号
	CompleteFlag *bool  `json:"complete_flag"` // 是否已完成(null=全部 true=已完成 false=待执行)
}

// AddVisitPlanReq 新增回访计划（Java: AddVisitPlanReq）
// @request
type AddVisitPlanReq struct {
	ElderID  *int64     `json:"elder_id"`  // 老人编号
	Title    *string    `json:"title"`     // 回访计划标题
	PlanDate *time.Time `json:"plan_date"` // 计划回访时间 yyyy-MM-dd HH:mm:ss
}

// CompleteVisitPlanReq 执行回访计划（Java: CompleteVisitPlanReq）
// @request
type CompleteVisitPlanReq struct {
	ID           *int64     `json:"id"`            // 回访计划编号
	Content      *string    `json:"content"`       // 回访内容
	CompleteDate *time.Time `json:"complete_date"` // 完成时间 yyyy-MM-dd HH:mm:ss
}

// DeleteVisitPlanReq 删除回访计划（Java: deleteVisitPlan 入参 visitPlanId）
// @request
type DeleteVisitPlanReq struct {
	ID *int64 `json:"id"` // 回访计划编号
}

// ListLabelReq 标签下拉列表（Java: listLabel 无入参, 用空结构体即可, 此处仅占位）
// @request
type ListLabelReq struct{}

// GetElderLabelByIdReq 按老人编号查询其标签（Java: getElderLabelById）
// @request
type GetElderLabelByIdReq struct {
	ElderID *int64 `json:"elder_id"` // 老人编号
}

// GetEditElderLabelByIdReq 编辑时按老人编号查询标签（按分类分组）
// @request
type GetEditElderLabelByIdReq struct {
	ElderID *int64 `json:"elder_id"` // 老人编号
}

// PageCommunicationRecordReq / AddCommunicationRecordReq / EditCommunicationRecordReq
// 已在 consult.go 中定义（沟通记录被 Consult 与 CheckContract/Intention 共用），此处不再重复。

// DeleteCommunicationRecordReq 删除沟通记录
// @request
type DeleteCommunicationRecordReq struct {
	ID *int64 `json:"id"` // 编号
}

// PageVisitPlanResp 分页查询回访计划响应
// @response
type PageVisitPlanResp struct {
	ID           int64     `json:"id"`            // 回访计划编号
	ElderID      int64     `json:"elder_id"`      // 老人编号
	Title        string    `json:"title"`         // 回访计划标题
	PlanDate     time.Time `json:"plan_date"`     // 计划回访时间
	CompleteDate time.Time `json:"complete_date"` // 完成时间
	Content      string    `json:"content"`       // 回访内容
}
