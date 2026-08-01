package dto

import "time"

// ============ ConsumeController 请求 ============

// @request
// PageConsumeByKeyQuery 分页查询消费记录请求
type PageConsumeByKeyQuery struct {
	PageNum   *int    `json:"pageNum" valid:"required"` // 页码
	PageSize  *int    `json:"pageSize" valid:"required"` // 条数
	ElderName *string `json:"elderName"` // 老人姓名
	StartTime *string `json:"startTime"` // 开始时间
	EndTime   *string `json:"endTime"` // 结束时间
}

// @request
// AddConsumeQuery 新增消费记录请求
type AddConsumeQuery struct {
	ID           *int64   `json:"id"` // id
	ElderID      *int64   `json:"elderId" valid:"required"` // 老人编号
	ConsumeType  *string  `json:"consumeType" valid:"required"` // 消费类别
	ConsumeAmount *float64 `json:"consumeAmount" valid:"required"` // 消费金额
	ConsumeDate  *string  `json:"consumeDate" valid:"required"` // 消费日期
	Remark       *string  `json:"remark"` // 备注
}

// @request
// EditConsumeQuery 编辑消费记录请求
type EditConsumeQuery struct {
	ID           *int64   `json:"id"` // id
	ElderID      *int64   `json:"elderId" valid:"required"` // 老人编号
	ConsumeType  *string  `json:"consumeType" valid:"required"` // 消费类别
	ConsumeAmount *float64 `json:"consumeAmount" valid:"required"` // 消费金额
	ConsumeDate  *string  `json:"consumeDate" valid:"required"` // 消费日期
	Remark       *string  `json:"remark"` // 备注
}

// ============ ConsumeController 响应 ============

// @response
// PageConsumeByKeyVO 分页查询消费记录响应
type PageConsumeByKeyVO struct {
	Rank
	ElderName    string    `json:"elderName"` // 老人姓名
	IDNum        string    `json:"idNum"` // 身份证号
	ConsumeType  string    `json:"consumeType"` // 消费类别
	ConsumeAmount float64  `json:"consumeAmount"` // 消费金额
	ConsumeDate  time.Time `json:"consumeDate"` // 消费日期
}