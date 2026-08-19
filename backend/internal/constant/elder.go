package constant

import "github.com/linbaozhong/gentity/pkg/types"

// CheckStatus 咨询 / 客户状态
type CheckStatus int8

// 咨询 / 客户状态
const (
	CheckConsult   CheckStatus = 0 // 咨询中
	CheckIntention CheckStatus = 1 // 意向跟进
	CheckReserve   CheckStatus = 2 // 预定
	CheckEnter     CheckStatus = 3 // 入住
	CheckExitAudit CheckStatus = 4 // 退住审核
	CheckExit      CheckStatus = 5 // 已退住
	CheckFailure   CheckStatus = 6 // 已失效
)

func (c CheckStatus) String() string {
	switch c {
	case CheckConsult:
		return "咨询中"
	case CheckIntention:
		return "意向跟进"
	case CheckReserve:
		return "预定"
	case CheckEnter:
		return "入住"
	case CheckExitAudit:
		return "退住审核"
	case CheckExit:
		return "已退住"
	case CheckFailure:
		return "已失效"
	default:
		return "未知"
	}
}

// Sex 性别
type Sex int8

const (
	SexMale   Sex = 0 // 男
	SexFemale Sex = 1 // 女
)

func (s Sex) String() string {
	switch s {
	case SexMale:
		return "男"
	case SexFemale:
		return "女"
	default:
		return "未知"
	}
}

// ElderCheckFlag 老人入住状态（tblelder.check_flag 字典）
type ElderCheckFlag int8

const (
	ElderCheckNotEnter ElderCheckFlag = 0 // 未入住
	ElderCheckEntered  ElderCheckFlag = 1 // 入住中
	ElderCheckExited   ElderCheckFlag = 2 // 已离院
)

func (e ElderCheckFlag) String() string {
	switch e {
	case ElderCheckNotEnter:
		return "未入住"
	case ElderCheckEntered:
		return "入住中"
	case ElderCheckExited:
		return "已离院"
	default:
		return "未知"
	}
}

// 老人
var (
	ErrElderAlreadyDelete = types.NewError(500, "该老人已被删除")
	ErrIdNumRepeat        = types.NewError(500, "老人身份证号已存在")
	ErrBalanceDeficiency  = types.NewError(500, "老人余额不足")
	ErrElderNotEnter      = types.NewError(500, "老人暂未入住")
	ErrElderNotExitAudit  = types.NewError(500, "老人暂未申请退住")
	ErrAlreadyIntention   = types.NewError(500, "该老人已是意向客户")
	ErrAlreadyReserve     = types.NewError(500, "该老人已预定")
	ErrAlreadyEnter       = types.NewError(500, "该老人已入住")
)

// 标签
var (
	ErrLabelTypeRepeat = types.NewError(500, "标签分类已存在")
	ErrLabelTypeOut    = types.NewError(500, "标签分类总数超过限制")
	ErrLabelNotNull    = types.NewError(500, "该标签分类存在标签，删除失败")
	ErrLabelRepeat     = types.NewError(500, "标签已存在")
	ErrLabelOut        = types.NewError(500, "该分类标签总数超过限制")
)

// 员工
var (
	ErrPhoneRepeat = types.NewError(500, "该手机号已存在")
	ErrEmailRepeat = types.NewError(500, "该邮箱已存在")
)
