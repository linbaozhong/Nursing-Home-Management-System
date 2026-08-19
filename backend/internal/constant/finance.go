package constant

import "github.com/linbaozhong/gentity/pkg/types"

// ChargeMode 收费方式
type ChargeMode int8

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

// 预定 / 退款
var (
	ErrRefundRepeat  = types.NewError(500, "请勿重复退款")
	ErrDueDateExpire = types.NewError(500, "预定已过期")
)

// 预存充值
var (
	ErrNotEnter = types.NewError(500, "该老人非入住状态，不予充值")
)
