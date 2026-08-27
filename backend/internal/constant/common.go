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
	ErrParamInvalid       = types.NewError(500, "请求参数异常")
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

// State 管理状态（通用，替代 del_flag 的 Y/N）
type State int8

// 管理状态枚举
const (
	StateDeleted  State = -1 // 删除
	StateDisabled State = 0  // 禁用
	StateEnabled  State = 1  // 启用
)

func (s State) String() string {
	switch s {
	case StateDeleted:
		return "已删除"
	case StateDisabled:
		return "已禁用"
	case StateEnabled:
		return "正常"
	default:
		return "未知"
	}
}

// 审计动作（audit_log.action）
const (
	AuditCreate = "create" // 新增
	AuditUpdate = "update" // 编辑
	AuditDelete = "delete" // 删除
)

// 字段中文名 cache 相关 key
const (
	AuditFieldDictCache = "audit_field_dict" // 字段中文名字典缓存（进程内）
)
