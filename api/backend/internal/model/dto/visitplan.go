package dto

// @request
// PageVisitPlanQuery 分页查询回访计划（Java: PageVisitPlanQuery）
type PageVisitPlanQuery struct {
	PageNum      *int   `json:"page_num"`      // 页码
	PageSize     *int   `json:"page_size"`     // 每页数量
	ElderID      *int64 `json:"elder_id"`      // 老人编号
	CompleteFlag *bool  `json:"complete_flag"` // 是否已完成(null=全部 true=已完成 false=待执行)
}

// @request
// AddVisitPlanQuery 新增回访计划（Java: AddVisitPlanQuery）
type AddVisitPlanQuery struct {
	ElderID  *int64  `json:"elder_id"`  // 老人编号
	Title    *string `json:"title"`     // 回访计划标题
	PlanDate *string `json:"plan_date"` // 计划回访时间 yyyy-MM-dd HH:mm:ss
}

// @request
// CompleteVisitPlanQuery 执行回访计划（Java: CompleteVisitPlanQuery）
type CompleteVisitPlanQuery struct {
	ID           *int64  `json:"id"`            // 回访计划编号
	Content      *string `json:"content"`       // 回访内容
	CompleteDate *string `json:"complete_date"` // 完成时间 yyyy-MM-dd HH:mm:ss
}

// @request
// DeleteVisitPlanQuery 删除回访计划（Java: deleteVisitPlan 入参 visitPlanId）
type DeleteVisitPlanQuery struct {
	ID *int64 `json:"id"` // 回访计划编号
}

// @request
// ListLabelQuery 标签下拉列表（Java: listLabel 无入参, 用空结构体即可, 此处仅占位）
type ListLabelQuery struct{}

// @request
// GetElderLabelByIdQuery 按老人编号查询其标签（Java: getElderLabelById）
type GetElderLabelByIdQuery struct {
	ElderID *int64 `json:"elder_id"` // 老人编号
}

// @request
// GetEditElderLabelByIdQuery 编辑时按老人编号查询标签（按分类分组）
type GetEditElderLabelByIdQuery struct {
	ElderID *int64 `json:"elder_id"` // 老人编号
}

// PageCommunicationRecordQuery / AddCommunicationRecordQuery / EditCommunicationRecordQuery
// 已在 consult.go 中定义（沟通记录被 Consult 与 CheckContract/Intention 共用），此处不再重复。

// @request
// DeleteCommunicationRecordQuery 删除沟通记录
type DeleteCommunicationRecordQuery struct {
	ID *int64 `json:"id"` // 编号
}
