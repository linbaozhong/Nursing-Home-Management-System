package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type outward struct{}

func init() {
	ack.RegisterRoute(&outward{})
}

func (o *outward) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/outward")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageOutwardByKey", o.pageOutwardByKey)
	_g.Get("/getOutwardById", o.getOutwardById)
	_g.Post("/addOutward", o.addOutward)
	_g.Put("/editOutward", o.editOutward)
	_g.Delete("/deleteOutward", o.deleteOutward)
	_g.Get("/pageSearchElderByKey", o.pageSearchElderByKey)
	_g.Put("/delayReturn", o.delayReturn)
	_g.Put("/recordReturn", o.recordReturn)
}

func (o *outward) pageOutwardByKey(ctx ack.Context) {
	ack.Get(ctx, service.Outward.PageOutwardByKey)
}

func (o *outward) getOutwardById(ctx ack.Context) {
	ack.Get(ctx, service.Outward.GetOutwardById)
}

func (o *outward) addOutward(ctx ack.Context) {
	ack.Post(ctx, service.Outward.AddOutward)
}

func (o *outward) editOutward(ctx ack.Context) {
	ack.Put(ctx, service.Outward.EditOutward)
}

func (o *outward) deleteOutward(ctx ack.Context) {
	ack.Delete(ctx, service.Outward.DeleteOutward)
}

func (o *outward) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Outward.PageSearchElderByKey)
}

func (o *outward) delayReturn(ctx ack.Context) {
	ack.Put(ctx, service.Outward.DelayReturn)
}

func (o *outward) recordReturn(ctx ack.Context) {
	ack.Put(ctx, service.Outward.RecordReturn)
}
