package dto

// ============ 微信登录请求 ============

// WxLoginReq 小程序微信静默登录请求
// @request
type WxLoginReq struct {
	Code *string `json:"code" valid:"required"` // 微信登录 code
}

// ============ 微信登录响应 ============

// WxLoginResp 小程序微信登录响应（智能分流）
// @response
type WxLoginResp struct {
	NeedBind bool           `json:"need_bind"` // 是否需绑定/选择企业
	Token    string         `json:"token"`     // 登录 token（唯一企业时返回）
	User     *LoginUserResp `json:"user"`      // 当前登录用户信息
	Tenants  []TenantResp   `json:"tenants"`   // 已绑定企业列表（多企业时返回）
}
