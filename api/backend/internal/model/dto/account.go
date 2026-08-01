package dto

// ============ AccountController 请求 ============

// @request
// LoginQuery 登录请求
type LoginQuery struct {
	Phone *string `json:"phone" valid:"required"` // 电话
	Pass  *string `json:"pass" valid:"required"` // 密码
}

// @request
// SendCodeQuery 发送验证码请求
type SendCodeQuery struct {
	Phone *string `json:"phone" valid:"required"` // 电话
}

// @request
// ForgetQuery 忘记密码请求
type ForgetQuery struct {
	Account *string `json:"account" valid:"required"` // 账号
	Pass    *string `json:"pass" valid:"required"` // 密码
	Code    *string `json:"code" valid:"required"` // 验证码
}

// @request
// EditQuery 修改账户请求
type EditQuery struct {
	ID       *int64  `json:"id" valid:"required"` // 编号
	OldPass  *string `json:"oldPass" valid:"required"` // 旧密码
	NewPass  *string `json:"newPass" valid:"required"` // 新密码
}

// ============ AccountController 响应 ============

// @response
// LoginUserVO 登录用户响应
type LoginUserVO struct {
	ID          int64    `json:"id"` // id
	Name        string   `json:"name"` // 姓名
	Avator      string   `json:"avator"` // 头像
	Phone       string   `json:"phone"` // 电话
	Pass        string   `json:"pass"` // 密码
	AuthIDList  []int64  `json:"authIdList"` // 权限id列表
	AuthUrlList []string `json:"authUrlList"` // 权限url列表
	Token       string   `json:"token"` // token
}