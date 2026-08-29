package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ ElderAccountLedger 老人资金明细台账 ============

// PageAccountLedgerByKeyReq 分页查询老人资金明细台账请求
// @request
type PageAccountLedgerByKeyReq struct {
	PageNum   *int       `json:"page_num" valid:"required"`  // 页码
	PageSize  *int       `json:"page_size" valid:"required"` // 条数
	ElderName *string    `json:"elder_name"`                 // 老人姓名（模糊）
	Direction *int8      `json:"direction"`                  // 方向：1=入账，2=出账（可空=全部）
	StartTime *time.Time `json:"start_time"`                 // 开始时间
	EndTime   *time.Time `json:"end_time"`                   // 结束时间
}

// GetAccountLedgerByIdReq 按编号查询资金明细请求
// @request
// type GetAccountLedgerByIdReq struct {
// 	ID *int64 `json:"id" valid:"required"` // 明细编号
// }

// ChangeBalanceReq 老人余额增减记账请求（ChangeBalance）
// @request
type ChangeBalanceReq struct {
	ElderID    *int64       `json:"elder_id" valid:"required"`  // 老人编号
	Direction  *int8        `json:"direction" valid:"required"` // 资金方向：1=入账，2=出账
	Amount     *types.Money `json:"amount" valid:"required"`    // 变更金额（对外元，内部存分）
	SourceType *string      `json:"source_type"`                // 来源类型：RECHARGE/FEED/NURSING/REFUND/MANUAL（空则默认 MANUAL）
	SourceID   *int64       `json:"source_id"`                  // 来源业务表主键id（用于幂等）
	BusinessNo *string      `json:"business_no"`                // 业务单号
	Remark     *string      `json:"remark"`                     // 备注
}

// PageAccountLedgerByKeyResp 分页查询老人资金明细台账响应
// @response
type PageAccountLedgerByKeyResp struct {
	ID           types.BigInt `json:"id"`            // 明细编号
	ElderName    string       `json:"elder_name"`    // 老人姓名
	Direction    int8         `json:"direction"`     // 方向：1=入账，2=出账
	Amount       types.Money  `json:"amount"`        // 变更金额（分）
	BalanceAfter types.Money  `json:"balance_after"` // 变动后余额（分）
	SourceType   string       `json:"source_type"`   // 来源类型
	BusinessNo   string       `json:"business_no"`   // 业务单号
	Remark       string       `json:"remark"`        // 备注
	CreateTime   time.Time    `json:"create_time"`   // 记账时间
}

// PageAccountLedgerBalanceResp 老人资金账户汇总响应
// @response
type PageAccountLedgerBalanceResp struct {
	ElderID      types.BigInt `json:"elder_id"`      // 老人编号
	ElderName    string       `json:"elder_name"`    // 老人姓名
	Balance      types.Money  `json:"balance"`       // 当前余额（分）
	TotalIncome  types.Money  `json:"total_income"`  // 累计入账（分）
	TotalOutcome types.Money  `json:"total_outcome"` // 累计出账（分）
}
