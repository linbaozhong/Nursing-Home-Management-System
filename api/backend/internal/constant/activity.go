package constant

import "github.com/linbaozhong/gentity/pkg/types"

// 来源渠道
var (
	ErrSourceRepeat = types.NewError(500, "来源渠道已存在")
)

// 活动分类
var (
	ErrActiveTypeRepeat = types.NewError(500, "活动分类已存在")
)

// 活动
var (
	ErrActiveRepeat = types.NewError(500, "活动名称已存在")
)

// ConsumeType 消费类型
type ConsumeType int8

// 消费类型
const (
	ConsumeNurse  ConsumeType = 1 // 护理
	ConsumeDishes ConsumeType = 2 // 点餐
)

func (c ConsumeType) String() string {
	switch c {
	case ConsumeNurse:
		return "护理"
	case ConsumeDishes:
		return "点餐"
	default:
		return "未知"
	}
}

// 回访计划
var (
	ErrVisitPlanRepeat = types.NewError(500, "回访计划标题已存在")
)

// 忘记密码 / 修改密码 / 上传文件
var (
	ErrAccountFormatError = types.NewError(500, "账号格式有误")
	ErrAccountNotRegister = types.NewError(500, "账号未注册")
	ErrPassSame           = types.NewError(500, "新密码与原密码相同")
	ErrRepeatSendCode     = types.NewError(500, "请勿重复发送验证码")
	ErrCodeExpire         = types.NewError(500, "验证码过期")
	ErrCodeError          = types.NewError(500, "验证码错误")
	ErrOldPassError       = types.NewError(500, "原密码错误")
	ErrUploadSuffixError  = types.NewError(500, "不支持的文件后缀")
)
