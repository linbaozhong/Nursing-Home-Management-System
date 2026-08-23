package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ Tenant 请求 ============

// RegisterTenantReq 租户自助注册请求
// @request
type RegisterTenantReq struct {
	Name         *string `json:"name" valid:"required"`          // 企业名称
	Logo         *string `json:"logo"`                           // 企业logo
	ContactName  *string `json:"contact_name" valid:"required"`  // 联系人姓名
	ContactPhone *string `json:"contact_phone" valid:"required"` // 联系电话（即管理员手机号）
	Password     *string `json:"password" valid:"required"`      // 管理员密码
	WxCode       *string `json:"wx_code"`                        // 微信登录code（可选，绑定微信）
}

// OpenTenantReq 平台开通/解锁租户请求
// @request
type OpenTenantReq struct {
	ID *int64 `json:"id" valid:"required"` // 租户编号
}

// SwitchTenantReq 切换当前租户请求
// @request
type SwitchTenantReq struct {
	TenantID *int64 `json:"tenant_id" valid:"required"` // 目标租户编号
}

// InviteMemberReq 邀请成员请求
// @request
type InviteMemberReq struct {
	TenantID *int64  `json:"tenant_id" valid:"required"` // 租户编号
	Phone    *string `json:"phone" valid:"required"`     // 被邀请人手机号
	RoleID   *int64  `json:"role_id" valid:"required"`   // 角色编号
}

// JoinMemberReq 加入已有企业（邀请码）请求
// @request
type JoinMemberReq struct {
	InviteCode *string `json:"invite_code" valid:"required"` // 邀请码
}

// ============ Tenant 响应 ============

// TenantResp 租户信息
// @response
type TenantResp struct {
	ID           types.BigInt `json:"id"`            // 租户编号
	Name         string       `json:"name"`          // 企业名称
	Logo         string       `json:"logo"`          // 企业logo
	ContactName  string       `json:"contact_name"`  // 联系人姓名
	ContactPhone string       `json:"contact_phone"` // 联系电话
	Plan         string       `json:"plan"`          // 套餐
	Status       int8         `json:"status"`        // 状态：0试用中 1正式 2锁定
	TrialStart   time.Time    `json:"trial_start"`   // 试用开始
	TrialEnd     time.Time    `json:"trial_end"`     // 试用结束
}

// UserTenantListResp 用户已绑定企业列表
// @response
type UserTenantListResp struct {
	Tenants []TenantResp `json:"tenants"` // 企业列表
	Current types.BigInt `json:"current"` // 当前企业编号（0表示未选择）
}

// MemberResp 成员信息
// @response
type MemberResp struct {
	ID       types.BigInt `json:"id"`        // 成员编号
	UserID   types.BigInt `json:"user_id"`   // 全局用户编号
	Name     string       `json:"name"`      // 姓名
	Phone    string       `json:"phone"`     // 手机号
	RoleID   types.BigInt `json:"role_id"`   // 角色编号
	Status   int8         `json:"status"`    // 状态：0在职 1离职
	AuthUrls []string     `json:"auth_urls"` // 权限url列表
}
