package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type familyrecharge struct{}

func init() {
	ack.RegisterRoute(&familyrecharge{})
}

// RegisterRoute 控制器路由注册
func (a *familyrecharge) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/family/recharge")

	// 绑定微信 openid（需登录）
	_g.Use(lib.AuthMiddleware())
	_g.Post("/bindOpenid", a.bindOpenid)
	_g.Post("/unifiedOrder", a.unifiedOrder)

	// 微信支付回调（无需登录、无需鉴权）
	_g2 := ack.NewParty(group, "/family/recharge")
	_g2.Post("/notify", a.notify)
}

// 绑定微信 openid
// @Summary 家属绑定微信 openid（充值前置）
// @Tags 家属充值
// @Accept application/json
// @Produce application/json
// @Param data body dto.BindOpenidQuery true "BindOpenidQuery"
// @Success 200 {object} dto.BindOpenidVO
// @Router /family/recharge/bindOpenid [post]
func (a *familyrecharge) bindOpenid(ctx ack.Context) {
	ack.Post(ctx, service.Family.BindOpenid)
}

// 统一下单
// @Summary 家属充值统一下单
// @Tags 家属充值
// @Accept application/json
// @Produce application/json
// @Param data body dto.RechargeUnifiedOrderQuery true "RechargeUnifiedOrderQuery"
// @Success 200 {object} dto.RechargeUnifiedOrderVO
// @Router /family/recharge/unifiedOrder [post]
func (a *familyrecharge) unifiedOrder(ctx ack.Context) {
	ack.Post(ctx, service.FamilyRechargeSvc.RechargeUnifiedOrder)
}

// 支付结果回调
// @Summary 微信支付结果回调
// @Tags 家属充值
// @Accept application/json
// @Produce application/json
// @Param data body dto.WechatPayNotifyQuery true "WechatPayNotifyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /family/recharge/notify [post]
func (a *familyrecharge) notify(ctx ack.Context) {
	ack.Post(ctx, service.FamilyRechargeSvc.PayNotify)
}
