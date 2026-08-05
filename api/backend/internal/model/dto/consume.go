package dto

import "time"

// ============ ConsumeController 请求 ============

// @request
// PageConsumeByKeyQuery 分页查询消费记录请求
type PageConsumeByKeyQuery struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	StartTime *string `json:"start_time"`                 // 开始时间
	EndTime   *string `json:"end_time"`                   // 结束时间
}

// @request
// AddConsumeQuery 新增消费记录请求
type AddConsumeQuery struct {
	ID            *int64   `json:"id"`                              // id
	ElderID       *int64   `json:"elder_id" valid:"required"`       // 老人编号
	ConsumeType   *string  `json:"consume_type" valid:"required"`   // 消费类别
	ConsumeAmount *float64 `json:"consume_amount" valid:"required"` // 消费金额
	ConsumeDate   *string  `json:"consume_date" valid:"required"`   // 消费日期
	Remark        *string  `json:"remark"`                          // 备注
}

// @request
// EditConsumeQuery 编辑消费记录请求
type EditConsumeQuery struct {
	ID            *int64   `json:"id"`                              // id
	ElderID       *int64   `json:"elder_id" valid:"required"`       // 老人编号
	ConsumeType   *string  `json:"consume_type" valid:"required"`   // 消费类别
	ConsumeAmount *float64 `json:"consume_amount" valid:"required"` // 消费金额
	ConsumeDate   *string  `json:"consume_date" valid:"required"`   // 消费日期
	Remark        *string  `json:"remark"`                          // 备注
}

// ============ ConsumeController 响应 ============

// @response
// PageConsumeByKeyVO 分页查询消费记录响应
type PageConsumeByKeyVO struct {
	ID            int64     `json:"id"`             // 消费记录编号
	ElderName     string    `json:"elder_name"`     // 老人姓名
	IDNum         string    `json:"id_num"`         // 身份证号
	ConsumeType   string    `json:"consume_type"`   // 消费类别
	ConsumeAmount float64   `json:"consume_amount"` // 消费金额
	ConsumeDate   time.Time `json:"consume_date"`   // 消费日期
}

// @response
// GetConsumeByIdVO 根据编号获取消费记录响应
type GetConsumeByIdVO struct {
	ID            int64     `json:"id"`             // 消费记录编号
	ElderID       int64     `json:"elder_id"`       // 老人编号
	ElderName     string    `json:"elder_name"`     // 老人姓名
	ConsumeType   string    `json:"consume_type"`   // 消费类别
	ConsumeAmount float64   `json:"consume_amount"` // 消费金额
	ConsumeDate   time.Time `json:"consume_date"`   // 消费日期
	Remark        string    `json:"remark"`         // 备注
}
