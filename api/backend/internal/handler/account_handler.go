package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
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

	// 发送验证码
	_g.Post("/sendCode", a.sendCode)

	// 忘记密码
	_g.Post("/forget", a.forget)

	// 需要登录
	_g.Use(lib.AuthMiddleware())
	_g.Put("/edit", a.edit)
	_g.Get("/logout", a.logout)
}

// login 登录
func (a *account) login(ctx ack.Context) {
	ack.Post(ctx, service.Account.Login)
}

// sendCode 发送验证码
func (a *account) sendCode(ctx ack.Context) {
	ack.Post(ctx, service.Account.SendCode)
}

// forget 忘记密码
func (a *account) forget(ctx ack.Context) {
	ack.Post(ctx, service.Account.Forget)
}

// edit 修改账户
func (a *account) edit(ctx ack.Context) {
	ack.Put(ctx, service.Account.Edit)
}

// logout 退出登录
func (a *account) logout(ctx ack.Context) {
	ack.Get(ctx, service.Account.Logout)
}
