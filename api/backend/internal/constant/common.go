package constant

import "github.com/linbaozhong/gentity/pkg/types"

// 授权 / 认证
const (
	Authorization = "Authorization"
)

// 响应码
var (
	CodeSuccess = types.NewError(200, "操作成功")
	CodeError   = types.NewError(500, "操作失败")
)

// 通用异常错误码
var (
	ErrAuthorizationNotFound = types.NewError(400, "访问令牌缺失")

	ErrSysError           = types.NewError(500, "系统繁忙")
	ErrNotExist           = types.NewError(404, "资源不存在")
	ErrDataNotExist       = types.NewError(404, "数据不存在")
	ErrPhoneOrEmailRepeat = types.NewError(500, "手机号或邮箱已注册")
	ErrPhoneError         = types.NewError(500, "手机号格式错误")
	ErrNullCode           = types.NewError(500, "验证码不能为空")
	ErrNoRegister         = types.NewError(500, "账号未注册")
	ErrPassword           = types.NewError(500, "密码错误")
	ErrMethodError        = types.NewError(500, "请求方式错误")
	ErrParamError         = types.NewError(500, "请求参数异常")
	ErrAccessError        = types.NewError(500, "非法访问")
	ErrTokenError         = types.NewError(500, "令牌无效")
	ErrCertificationError = types.NewError(500, "认证失败")
	ErrAuthError          = types.NewError(500, "权限不足")
)

// YesNo 是否标记
type YesNo int8

// 是否（Y/N）
const (
	YesNoNo  YesNo = 0 // N
	YesNoYes YesNo = 1 // Y
)

func (y YesNo) String() string {
	switch y {
	case YesNoNo:
		return "否"
	case YesNoYes:
		return "是"
	default:
		return "否"
	}
}
