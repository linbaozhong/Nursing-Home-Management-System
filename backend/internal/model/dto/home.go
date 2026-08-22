package dto

import "time"

// ============ HomeController 请求 ============

// ClientSourceReq 客户来源渠道请求
// @request
type ClientSourceReq struct {
	StartTime *time.Time `json:"start_time"` // 开始时间
	EndTime   *time.Time `json:"end_time"`   // 结束时间
}

// ============ HomeController 响应 ============

// TodayOverviewResp 今日概览响应
// @response
type TodayOverviewResp struct {
	TodayAddConsultNum     int64 `json:"today_add_consult_num"`     // 今日新增咨询
	TodayAddReserveNum     int64 `json:"today_add_reserve_num"`     // 今日新增预定
	TodayAddContractNum    int64 `json:"today_add_contract_num"`    // 今日新增合同
	TodayContractExpireNum int64 `json:"today_contract_expire_num"` // 合同到期提醒
}

// AvailableBedResp 可售床位响应
// @response
type AvailableBedResp struct {
	IdleRoomNum  int64 `json:"idle_room_num"`  // 空闲房间数
	IdleBedNum   int64 `json:"idle_bed_num"`   // 空闲床位数
	ExitAuditNum int64 `json:"exit_audit_num"` // 退住审核数
}

// TodaySaleFollowResp 今日销售跟进响应
// @response
type TodaySaleFollowResp struct {
	TodayReturnVisitNum   int64 `json:"today_return_visit_num"`   // 今日待回访数
	TodayReturnedVisitNum int64 `json:"today_returned_visit_num"` // 今日已回访数
	StayReturnedVisitNum  int64 `json:"stay_returned_visit_num"`  // 滞留回访数
}

// MonthPerformanceRankResp 本月业绩排行响应
// @response
type MonthPerformanceRankResp struct {
	ConsultClientNum           int64      `json:"consult_client_num"`            // 咨询客户数
	ConsultClientFloatRate     float64    `json:"consult_client_float_rate"`     // 咨询客户浮动率
	SignContractNum            int64      `json:"sign_contract_num"`             // 签约合同数
	SignContractFloatRate      float64    `json:"sign_contract_float_rate"`      // 签约合同浮动率
	ConsultConversionRate      float64    `json:"consult_conversion_rate"`       // 咨询转化率
	ConsultConversionFloatRate float64    `json:"consult_conversion_float_rate"` // 咨询转化浮动率
	SaleRankList               []SaleRank `json:"sale_rank_list"`                // 销售排行列表
}

// SaleRank 业务员销售排行（嵌套）
// @response
type SaleRank struct {
	Rank        int64  `json:"rank"`         // 排行
	Name        string `json:"name"`         // 姓名
	ConsultNum  int64  `json:"consult_num"`  // 咨询数
	ContractNum int64  `json:"contract_num"` // 合同数
}

// ClientSourceResp 客户来源渠道响应
// @response
type ClientSourceResp struct {
	SourceName string `json:"source_name"` // 来源渠道名称
	ConsultNum int64  `json:"consult_num"` // 咨询人数
}

// BusinessTrendResp 业务趋势响应
// @response
type BusinessTrendResp struct {
	Month       string `json:"month"`        // 月份
	ConsultNum  int64  `json:"consult_num"`  // 咨询数
	ContractNum int64  `json:"contract_num"` // 合同数
}
