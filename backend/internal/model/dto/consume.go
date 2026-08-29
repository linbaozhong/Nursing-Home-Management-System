package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ ConsumeController 请求 ============

// PageConsumeByKeyReq 分页查询消费记录请求
// @request
type PageConsumeByKeyReq struct {
	PageNum   *int       `json:"page_num" valid:"required"`  // 页码
	PageSize  *int       `json:"page_size" valid:"required"` // 条数
	ElderName *string    `json:"elder_name"`                 // 老人姓名
	StartTime *time.Time `json:"start_time"`                 // 开始时间
	EndTime   *time.Time `json:"end_time"`                   // 结束时间
}

// AddConsumeReq 新增消费记录请求
// @request
type AddConsumeReq struct {
	ID            *int64       `json:"id"`                              // id
	ElderID       *int64       `json:"elder_id" valid:"required"`       // 老人编号
	ConsumeType   *string      `json:"consume_type" valid:"required"`   // 消费类别
	ConsumeAmount *types.Money `json:"consume_amount" valid:"required"` // 消费金额
	ConsumeDate   *time.Time   `json:"consume_date" valid:"required"`   // 消费日期
	Remark        *string      `json:"remark"`                          // 备注
	SourceType    *string      `json:"source_type"`                     // 来源类型：ORDER/NURSE_RESERVE/MANUAL
	SourceID      *int64       `json:"source_id"`                       // 来源业务主键id
	OutTradeNo    *string      `json:"out_trade_no"`                    // 外部交易单号（对账）
}

// EditConsumeReq 编辑消费记录请求
// @request
type EditConsumeReq struct {
	ID            *int64       `json:"id"`                              // id
	ElderID       *int64       `json:"elder_id" valid:"required"`       // 老人编号
	ConsumeType   *string      `json:"consume_type" valid:"required"`   // 消费类别
	ConsumeAmount *types.Money `json:"consume_amount" valid:"required"` // 消费金额
	ConsumeDate   *time.Time   `json:"consume_date" valid:"required"`   // 消费日期
	Remark        *string      `json:"remark"`                          // 备注
	SourceType    *string      `json:"source_type"`                     // 来源类型：ORDER/NURSE_RESERVE/MANUAL
	SourceID      *int64       `json:"source_id"`                       // 来源业务主键id
	OutTradeNo    *string      `json:"out_trade_no"`                    // 外部交易单号（对账）
}

// ============ ConsumeController 响应 ============

// PageConsumeByKeyResp 分页查询消费记录响应
// @response
type PageConsumeByKeyResp struct {
	ID            types.BigInt `json:"id"`             // 消费记录编号
	ElderName     string       `json:"elder_name"`     // 老人姓名
	IDNum         string       `json:"id_num"`         // 身份证号
	ConsumeType   string       `json:"consume_type"`   // 消费类别
	ConsumeAmount types.Money  `json:"consume_amount"` // 消费金额
	ConsumeDate   time.Time    `json:"consume_date"`   // 消费日期
}

// GetConsumeByIdResp 根据编号获取消费记录响应
// @response
type GetConsumeByIdResp struct {
	ID            types.BigInt `json:"id"`             // 消费记录编号
	ElderID       types.BigInt `json:"elder_id"`       // 老人编号
	ElderName     string       `json:"elder_name"`     // 老人姓名
	ConsumeType   string       `json:"consume_type"`   // 消费类别
	ConsumeAmount types.Money  `json:"consume_amount"` // 消费金额
	ConsumeDate   time.Time    `json:"consume_date"`   // 消费日期
	Remark        string       `json:"remark"`         // 备注
	SourceType    string       `json:"source_type"`    // 来源类型
	SourceID      types.BigInt `json:"source_id"`      // 来源业务主键id
	OutTradeNo    string       `json:"out_trade_no"`   // 外部交易单号
}
