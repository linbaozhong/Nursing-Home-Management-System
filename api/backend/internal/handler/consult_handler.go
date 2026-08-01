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

func (c *consult) pageConsultByKey(ctx ack.Context) {
	ack.Get(ctx, service.Consult.PageConsultByKey)
}

func (c *consult) getConsultByConsultIdAndElderId(ctx ack.Context) {
	ack.Get(ctx, service.Consult.GetConsultByConsultIdAndElderId)
}

func (c *consult) addConsult(ctx ack.Context) {
	ack.Post(ctx, service.Consult.AddConsult)
}

func (c *consult) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Consult.PageSearchElderByKey)
}

func (c *consult) pageIntentByKey(ctx ack.Context) {
	ack.Get(ctx, service.Consult.PageIntentByKey)
}
