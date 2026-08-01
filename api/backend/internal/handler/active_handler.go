package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type active struct{}

func init() {
	ack.RegisterRoute(&active{})
}

func (a *active) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/active")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageActiveByKey", a.pageActiveByKey)
	_g.Get("/getActiveById", a.getActiveById)
	_g.Post("/addActive", a.addActive)
	_g.Put("/editActive", a.editActive)
	_g.Delete("/deleteActive", a.deleteActive)
	_g.Get("/pageSearchElderByKey", a.pageSearchElderByKey)
}

func (a *active) pageActiveByKey(ctx ack.Context) {
	ack.Get(ctx, service.Active.PageActiveByKey)
}

func (a *active) getActiveById(ctx ack.Context) {
	ack.Get(ctx, service.Active.GetActiveById)
}

func (a *active) addActive(ctx ack.Context) {
	ack.Post(ctx, service.Active.AddActive)
}

func (a *active) editActive(ctx ack.Context) {
	ack.Put(ctx, service.Active.EditActive)
}

func (a *active) deleteActive(ctx ack.Context) {
	ack.Delete(ctx, service.Active.DeleteActive)
}

func (a *active) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Active.PageSearchElderByKey)
}
