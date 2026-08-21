package constant

import "github.com/linbaozhong/gentity/pkg/types"

// 家属端相关错误
var (
	ErrFamilyNotBind    = types.NewError(400, "该手机号尚未注册或绑定家属")
	ErrFamilyPhoneUsed  = types.NewError(400, "该手机号已注册")
	ErrFamilyElderMatch = types.NewError(400, "老人信息校验失败（姓名/身份证后4位不匹配或未入住）")
	ErrFamilyNoElder    = types.NewError(400, "您暂未绑定任何老人")
	ErrWxPayFailed      = types.NewError(500, "微信支付下单失败")
)

// 家属 token 角色标识
const FamilyRole = "family"
