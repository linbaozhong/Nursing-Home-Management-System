package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type family struct{}

func init() {
	ack.RegisterRoute(&family{})
}

// RegisterRoute 控制器路由注册
func (a *family) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/family")

	// 发送验证码（无需登录）
	_g.Post("/sendCode", a.sendCode)

	// 注册并绑定老人（无需登录）
	_g.Post("/registerBind", a.registerBind)

	// 登录（无需登录）
	_g.Post("/login", a.login)

	// 以下需要登录
	_g.Use(lib.AuthMiddleware())
	_g.Post("/bindElder", a.bindElder)
	_g.Get("/myElders", a.myElders)
}

// 发送验证码
// @Summary 发送家属注册/绑定验证码
// @Tags 家属
// @Accept application/json
// @Produce application/json
// @Param data body dto.FamilySendCodeReq true "FamilySendCodeReq"
// @Success 200 {object} dto.LoginUserResp
// @Router /family/sendCode [post]
func (a *family) sendCode(ctx ack.Context) {
	ack.Post(ctx, service.Family.SendCode)
}

// 注册并绑定老人
// @Summary 家属注册并绑定一位老人
// @Tags 家属
// @Accept application/json
// @Produce application/json
// @Param data body dto.RegisterBindReq true "RegisterBindReq"
// @Success 200 {object} dto.EmptyResp
// @Router /family/registerBind [post]
func (a *family) registerBind(ctx ack.Context) {
	ack.Post(ctx, service.Family.RegisterBind)
}

// 登录
// @Summary 家属登录
// @Tags 家属
// @Accept application/json
// @Produce application/json
// @Param data body dto.FamilyLoginReq true "FamilyLoginReq"
// @Success 200 {object} dto.FamilyLoginResp
// @Router /family/login [post]
func (a *family) login(ctx ack.Context) {
	ack.Post(ctx, service.Family.Login)
}

// 绑定更多老人
// @Summary 已注册家属绑定更多老人
// @Tags 家属
// @Accept application/json
// @Produce application/json
// @Param data body dto.BindElderReq true "BindElderReq"
// @Success 200 {object} dto.EmptyResp
// @Router /family/bindElder [post]
func (a *family) bindElder(ctx ack.Context) {
	ack.Post(ctx, service.Family.BindElder)
}

// 我的老人
// @Summary 我的老人列表（数据范围）
// @Tags 家属
// @Accept application/json
// @Produce application/json
// @Param data query dto.FamilyMyEldersReq true "FamilyMyEldersReq"
// @Success 200 {object} dto.FamilyMyEldersResp
// @Router /family/myElders [get]
func (a *family) myElders(ctx ack.Context) {
	ack.Get(ctx, service.Family.MyElders)
}
