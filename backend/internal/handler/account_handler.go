package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type account struct{}

func init() {
	ack.RegisterRoute(&account{})
}

// RegisterRoute 控制器路由注册
func (a *account) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/account")

	// 登录
	_g.Post("/login", a.login)

	// 微信静默登录
	_g.Post("/wxLogin", a.wxLogin)

	// 发送验证码
	_g.Post("/sendCode", a.sendCode)

	// 忘记密码
	_g.Post("/forget", a.forget)

	// 需要登录
	_g.Use(lib.AuthMiddleware())
	_g.Post("/edit", a.edit)
	_g.Get("/logout", a.logout)
}

// 账户
// @Summary 账户
// @Tags 账户
// @Accept application/json
// @Produce application/json
// @Param data body dto.LoginQuery true "LoginQuery"
// @Success 200 {object} dto.LoginUserVO
// @Router /account/login [post]
func (a *account) login(ctx ack.Context) {
	ack.Post(ctx, service.Account.Login)
}

// 微信静默登录
// @Summary 微信静默登录
// @Tags 账户
// @Accept application/json
// @Produce application/json
// @Param data body dto.WxLoginQuery true "WxLoginQuery"
// @Success 200 {object} dto.WxLoginVO
// @Router /account/wxLogin [post]
func (a *account) wxLogin(ctx ack.Context) {
	ack.Post(ctx, service.Tenant.WxLogin)
}

// 发送验证码
// @Summary 发送验证码
// @Tags 账户
// @Accept application/json
// @Produce application/json
// @Param data body dto.SendCodeQuery true "SendCodeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /account/sendCode [post]
func (a *account) sendCode(ctx ack.Context) {
	ack.Post(ctx, service.Account.SendCode)
}

// 账户
// @Summary 账户
// @Tags 账户
// @Accept application/json
// @Produce application/json
// @Param data body dto.ForgetQuery true "ForgetQuery"
// @Success 200 {object} dto.LoginUserVO
// @Router /account/forget [post]
func (a *account) forget(ctx ack.Context) {
	ack.Post(ctx, service.Account.Forget)
}

// 账户
// @Summary 账户
// @Tags 账户
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditQuery true "EditQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /account/edit [post]
func (a *account) edit(ctx ack.Context) {
	ack.Post(ctx, service.Account.Edit)
}

// 账户
// @Summary 账户
// @Tags 账户
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /account/logout [get]
func (a *account) logout(ctx ack.Context) {
	ack.Get(ctx, service.Account.Logout)
}
