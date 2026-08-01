package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type activetype struct{}

func init() {
	ack.RegisterRoute(&activetype{})
}

func (a *activetype) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/activeType")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageActiveTypeByKey", a.pageActiveTypeByKey)
	_g.Get("/getActiveTypeById", a.getActiveTypeById)
	_g.Post("/addActiveType", a.addActiveType)
	_g.Put("/editActiveType", a.editActiveType)
	_g.Delete("/deleteActiveType", a.deleteActiveType)
}

func (a *activetype) pageActiveTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.ActiveType.PageActiveTypeByKey)
}

func (a *activetype) getActiveTypeById(ctx ack.Context) {
	ack.Get(ctx, service.ActiveType.GetActiveTypeById)
}

func (a *activetype) addActiveType(ctx ack.Context) {
	ack.Post(ctx, service.ActiveType.AddActiveType)
}

func (a *activetype) editActiveType(ctx ack.Context) {
	ack.Put(ctx, service.ActiveType.EditActiveType)
}

func (a *activetype) deleteActiveType(ctx ack.Context) {
	ack.Delete(ctx, service.ActiveType.DeleteActiveType)
}
