package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ AccountController 请求 ============

// LoginReq 登录请求
// @request
type LoginReq struct {
	Phone *string `json:"phone" valid:"required"` // 电话
	Pass  *string `json:"pass" valid:"required"`  // 密码
	Code  *string `json:"code"`                   // 验证码（短信登录时可选）
}

// SendCodeReq 发送验证码请求
// @request
type SendCodeReq struct {
	Phone *string `json:"phone" valid:"required"` // 电话
}

// SendCodeResp 发送验证码响应
// @response
type SendCodeResp struct {
	Code string `json:"code"`
}

// ForgetReq 忘记密码请求
// @request
type ForgetReq struct {
	Account *string `json:"account" valid:"required"` // 账号
	Pass    *string `json:"pass" valid:"required"`    // 密码
	Code    *string `json:"code" valid:"required"`    // 验证码
}

// EditReq 修改账户请求
// @request
type EditReq struct {
	ID      *int64  `json:"id" valid:"required"`       // 编号
	OldPass *string `json:"old_pass" valid:"required"` // 旧密码
	NewPass *string `json:"new_pass" valid:"required"` // 新密码
}

// ============ AccountController 响应 ============

// LoginUserResp 登录用户响应
// @response
type LoginUserResp struct {
	ID          types.BigInt   `json:"id"`            // 全局用户id
	Name        string         `json:"name"`          // 姓名
	Avator      string         `json:"avator"`        // 头像
	Phone       string         `json:"phone"`         // 电话
	TenantID    types.BigInt   `json:"tenant_id"`     // 当前租户编号
	MemberID    types.BigInt   `json:"member_id"`     // 当前成员编号
	RoleID      types.BigInt   `json:"role_id"`       // 当前角色编号
	NeedBind    bool           `json:"need_bind"`     // 是否需要选择/绑定企业
	AuthIDList  []types.BigInt `json:"auth_id_list"`  // 权限id列表
	AuthUrlList []string       `json:"auth_url_list"` // 权限url列表
	Tenants     []TenantResp   `json:"tenants"`       // 已绑定企业列表（多企业时返回）
	Token       string         `json:"token"`         // token
}
