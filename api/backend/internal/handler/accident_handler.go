package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type accident struct{}

func init() {
	ack.RegisterRoute(&accident{})
}

func (a *accident) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/accident")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageAccidentByKey", a.pageAccidentByKey)
	_g.Get("/getAccidentById", a.getAccidentById)
	_g.Post("/addAccident", a.addAccident)
	_g.Put("/editAccident", a.editAccident)
	_g.Delete("/deleteAccident", a.deleteAccident)
}

func (a *accident) pageAccidentByKey(ctx ack.Context) {
	ack.Get(ctx, service.Accident.PageAccidentByKey)
}

func (a *accident) getAccidentById(ctx ack.Context) {
	ack.Get(ctx, service.Accident.GetAccidentById)
}

func (a *accident) addAccident(ctx ack.Context) {
	ack.Post(ctx, service.Accident.AddAccident)
}

func (a *accident) editAccident(ctx ack.Context) {
	ack.Put(ctx, service.Accident.EditAccident)
}

func (a *accident) deleteAccident(ctx ack.Context) {
	ack.Delete(ctx, service.Accident.DeleteAccident)
}
