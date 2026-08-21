package constant

import "github.com/linbaozhong/gentity/pkg/types"

// 租户状态
const (
	TenantStatusTrial  int8 = 0 // 试用中
	TenantStatusNormal int8 = 1 // 正式
	TenantStatusLocked int8 = 2 // 锁定
)

// 租户默认试用期（天）
const TrialDays = 30

// 平台管理员租户 ID（存量数据归入的平台租户）
const PlatformTenantID int64 = 1

// 租户成员在职状态
const (
	MemberStatusActive int8 = 0 // 在职
	MemberStatusLeave  int8 = 1 // 离职
)

// 租户 token 角色标识
const TenantRole = "tenant"

// 平台管理员角色标识
const PlatformRole = "platform"

// 租户相关错误
var (
	ErrTenantNameRepeat    = types.NewError(400, "企业名称已存在")
	ErrTenantLocked        = types.NewError(403, "企业已被锁定，请联系平台管理员开通")
	ErrTenantExpired       = types.NewError(403, "企业试用期已结束，已锁定")
	ErrTenantNotExist      = types.NewError(404, "企业不存在")
	ErrTenantNotActive     = types.NewError(403, "企业未开通")
	ErrMemberNotExist      = types.NewError(404, "该成员不存在")
	ErrMemberAlreadyExist  = types.NewError(400, "您已加入该企业")
	ErrMemberInviteInvalid = types.NewError(400, "邀请码无效")
	ErrUserNotExist        = types.NewError(404, "用户不存在")
	ErrUserPhoneUsed       = types.NewError(400, "手机号已注册")
	ErrNotTenantAdmin      = types.NewError(403, "仅租户管理员可操作")
	ErrNotPlatformAdmin    = types.NewError(403, "仅平台管理员可操作")
	ErrWxLoginFailed       = types.NewError(500, "微信登录失败")
	ErrNeedBindTenant      = types.NewError(400, "请先绑定企业")
)
