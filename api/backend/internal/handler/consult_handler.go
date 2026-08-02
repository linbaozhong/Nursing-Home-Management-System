package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type consult struct{}

func init() {
	ack.RegisterRoute(&consult{})
}

func (c *consult) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/consult")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageConsultByKey", c.pageConsultByKey)
	_g.Get("/getConsultByConsultIdAndElderId", c.getConsultByConsultIdAndElderId)
	_g.Post("/addConsult", c.addConsult)
	_g.Get("/pageSearchElderByKey", c.pageSearchElderByKey)
	_g.Get("/pageIntentByKey", c.pageIntentByKey)
}

// 分页查询咨询
// @Summary 分页查询咨询
// @Tags 咨询
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageConsultByKeyQuery true "PageConsultByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /consult/pageConsultByKey [get]
func (c *consult) pageConsultByKey(ctx ack.Context) {
	ack.Get(ctx, service.Consult.PageConsultByKey)
}

// 获取咨询咨询And老人
// @Summary 获取咨询咨询And老人
// @Tags 咨询
// @Accept application/json
// @Produce application/json
// @Param data query dto.GetConsultByConsultIdAndElderIdQuery true "GetConsultByConsultIdAndElderIdQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /consult/getConsultByConsultIdAndElderId [get]
func (c *consult) getConsultByConsultIdAndElderId(ctx ack.Context) {
	ack.Get(ctx, service.Consult.GetConsultByConsultIdAndElderId)
}

// 新增咨询
// @Summary 新增咨询
// @Tags 咨询
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddConsultQuery true "AddConsultQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /consult/addConsult [post]
func (c *consult) addConsult(ctx ack.Context) {
	ack.Post(ctx, service.Consult.AddConsult)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 咨询
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /consult/pageSearchElderByKey [get]
func (c *consult) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Consult.PageSearchElderByKey)
}

// 分页查询Intent
// @Summary 分页查询Intent
// @Tags 咨询
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageIntentionByKeyQuery true "PageIntentionByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /consult/pageIntentByKey [get]
func (c *consult) pageIntentByKey(ctx ack.Context) {
	ack.Get(ctx, service.Consult.PageIntentByKey)
}
