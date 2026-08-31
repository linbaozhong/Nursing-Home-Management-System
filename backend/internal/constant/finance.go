package constant

import "github.com/linbaozhong/gentity/pkg/types"

// ChargeMode 收费方式
type ChargeMode uint8

// 收费方式
const (
	ChargeAll   ChargeMode = 0 // 所有
	ChargeOnce  ChargeMode = 1 // 按次
	ChargeMonth ChargeMode = 2 // 按月
)

func (c ChargeMode) String() string {
	switch c {
	case ChargeAll:
		return "所有"
	case ChargeOnce:
		return "按次"
	case ChargeMonth:
		return "按月"
	default:
		return "未知"
	}
}

// BillStatus 账单状态（bill.status）
type BillStatus uint8

// 账单状态
const (
	BillUnpaid   BillStatus = 0 // 未支付
	BillPartPaid BillStatus = 1 // 部分支付
	BillPaid     BillStatus = 2 // 已支付
	BillOverdue  BillStatus = 3 // 逾期
)

func (b BillStatus) String() string {
	switch b {
	case BillUnpaid:
		return "未支付"
	case BillPartPaid:
		return "部分支付"
	case BillPaid:
		return "已支付"
	case BillOverdue:
		return "逾期"
	default:
		return "未知"
	}
}

// RechargeStatus 充值状态（family_recharge.status）
type RechargeStatus uint8

// 充值状态
const (
	RechargeUnpaid RechargeStatus = 0 // 待支付
	RechargePaid   RechargeStatus = 1 // 已支付
	RechargeClosed RechargeStatus = 2 // 已关闭
)

func (r RechargeStatus) String() string {
	switch r {
	case RechargeUnpaid:
		return "待支付"
	case RechargePaid:
		return "已支付"
	case RechargeClosed:
		return "已关闭"
	default:
		return "未知"
	}
}

// DepositStatus 缴存状态（drug_deposit.status）
type DepositStatus uint8

// 缴存状态
const (
	DepositNotPaid  DepositStatus = 0 // 未缴
	DepositPartPaid DepositStatus = 1 // 部分缴
	DepositPaid     DepositStatus = 2 // 已缴
)

func (d DepositStatus) String() string {
	switch d {
	case DepositNotPaid:
		return "未缴"
	case DepositPartPaid:
		return "部分缴"
	case DepositPaid:
		return "已缴"
	default:
		return "未知"
	}
}

// ReserveStatus 退款状态（reserve.status）
type ReserveStatus uint8

// 退款状态
const (
	ReserveUnrefunded ReserveStatus = 0 // 未退款
	ReserveRefunded   ReserveStatus = 1 // 已退款
)

func (r ReserveStatus) String() string {
	switch r {
	case ReserveUnrefunded:
		return "未退款"
	case ReserveRefunded:
		return "已退款"
	default:
		return "未知"
	}
}

// 预定 / 退款
var (
	ErrRefundRepeat  = types.NewError(500, "请勿重复退款")
	ErrDueDateExpire = types.NewError(500, "预定已过期")
)

// 预存充值
var (
	ErrNotEnter = types.NewError(500, "该老人非入住状态，不予充值")
)
