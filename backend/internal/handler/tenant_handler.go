package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type tenant struct{}

func init() {
	ack.RegisterRoute(&tenant{})
}

// RegisterRoute 租户控制器路由注册
func (t *tenant) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/tenant")

	// 自助注册（无需登录）
	_g.Post("/register", t.register)

	// 需要登录
	_g.Use(lib.AuthMiddleware())
	// 我的企业列表
	_g.Get("/myTenants", t.myTenants)
	// 切换当前租户
	_g.Post("/switchTenant", t.switchTenant)

	// 仅平台管理员
	_g.Post("/open", t.open)
	_g.Post("/lock", t.lock)

	// 租户管理员邀请成员
	_g.Post("/inviteMember", t.inviteMember)
}

// 租户自助注册
// @Summary 租户自助注册
// @Tags 租户
// @Param data body dto.RegisterTenantReq true "RegisterTenantReq"
// @Success 200 {object} dto.LoginUserResp
// @Router /tenant/register [post]
func (t *tenant) register(ctx ack.Context) {
	ack.Post(ctx, service.Tenant.Register)
}

// 我的企业列表
// @Summary 我的企业列表
// @Tags 租户
// @Success 200 {object} dto.UserTenantListResp
// @Router /tenant/myTenants [get]
func (t *tenant) myTenants(ctx ack.Context) {
	ack.Get(ctx, service.Tenant.MyTenants)
}

// 切换当前租户
// @Summary 切换当前租户
// @Tags 租户
// @Param data body dto.SwitchTenantReq true "SwitchTenantReq"
// @Success 200 {object} dto.LoginUserResp
// @Router /tenant/switchTenant [post]
func (t *tenant) switchTenant(ctx ack.Context) {
	ack.Post(ctx, service.Tenant.SwitchTenant)
}

// 平台开通/解锁租户
// @Summary 平台开通/解锁租户
// @Tags 租户
// @Param data body dto.OpenTenantReq true "OpenTenantReq"
// @Success 200 {object} dto.EmptyResp
// @Router /tenant/open [post]
func (t *tenant) open(ctx ack.Context) {
	ack.Post(ctx, service.Tenant.Open)
}

// 试用到期锁定
// @Summary 试用到期锁定
// @Tags 租户
// @Param data body dto.OpenTenantReq true "OpenTenantReq"
// @Success 200 {object} dto.EmptyResp
// @Router /tenant/lock [post]
func (t *tenant) lock(ctx ack.Context) {
	ack.Post(ctx, service.Tenant.Lock)
}

// 邀请成员
// @Summary 邀请成员
// @Tags 租户
// @Param data body dto.InviteMemberReq true "InviteMemberReq"
// @Success 200 {object} dto.MemberResp
// @Router /tenant/inviteMember [post]
func (t *tenant) inviteMember(ctx ack.Context) {
	ack.Post(ctx, service.Tenant.InviteMember)
}
