package dto

// ============ HomeController 请求 ============

// @request
// ClientSourceQuery 客户来源渠道请求
type ClientSourceQuery struct {
	StartTime *string `json:"startTime"` // 开始时间
	EndTime   *string `json:"endTime"` // 结束时间
}

// ============ HomeController 响应 ============

// @response
// TodayOverviewVO 今日概览响应
type TodayOverviewVO struct {
	TodayAddConsultNum     int64 `json:"todayAddConsultNum"` // 今日新增咨询
	TodayAddReserveNum     int64 `json:"todayAddReserveNum"` // 今日新增预定
	TodayAddContractNum    int64 `json:"todayAddContractNum"` // 今日新增合同
	TodayContractExpireNum int64 `json:"todayContractExpireNum"` // 合同到期提醒
}

// @response
// AvailableBedVO 可售床位响应
type AvailableBedVO struct {
	IdleRoomNum int64 `json:"idleRoomNum"` // 空闲房间数
	IdleBedNum  int64 `json:"idleBedNum"` // 空闲床位数
	ExitAuditNum int64 `json:"exitAuditNum"` // 退住审核数
}

// @response
// TodaySaleFollowVO 今日销售跟进响应
type TodaySaleFollowVO struct {
	TodayReturnVisitNum    int64 `json:"todayReturnVisitNum"` // 今日待回访数
	TodayReturnedVisitNum  int64 `json:"todayReturnedVisitNum"` // 今日已回访数
	StayReturnedVisitNum   int64 `json:"stayReturnedVisitNum"` // 滞留回访数
}

// @response
// MonthPerformanceRankVO 本月业绩排行响应
type MonthPerformanceRankVO struct {
	ConsultClientNum           int64      `json:"consultClientNum"` // 咨询客户数
	ConsultClientFloatRate     float64    `json:"consultClientFloatRate"` // 咨询客户浮动率
	SignContractNum            int64      `json:"signContractNum"` // 签约合同数
	SignContractFloatRate      float64    `json:"signContractFloatRate"` // 签约合同浮动率
	ConsultConversionRate      float64    `json:"consultConversionRate"` // 咨询转化率
	ConsultConversionFloatRate float64    `json:"consultConversionFloatRate"` // 咨询转化浮动率
	SaleRankList               []SaleRank `json:"saleRankList"` // 销售排行列表
}

// @response
// SaleRank 业务员销售排行（嵌套）
type SaleRank struct {
	Rank        int64  `json:"rank"` // 排行
	Name        string `json:"name"` // 姓名
	ConsultNum  int64  `json:"consultNum"` // 咨询数
	ContractNum int64  `json:"contractNum"` // 合同数
}

// @response
// ClientSourceVO 客户来源渠道响应
type ClientSourceVO struct {
	SourceName string `json:"sourceName"` // 来源渠道名称
	ConsultNum int64  `json:"consultNum"` // 咨询人数
}

// @response
// BusinessTrendVO 业务趋势响应
type BusinessTrendVO struct {
	Month       string `json:"month"` // 月份
	ConsultNum  int64  `json:"consultNum"` // 咨询数
	ContractNum int64  `json:"contractNum"` // 合同数
}