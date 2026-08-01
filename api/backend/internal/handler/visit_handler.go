package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type visit struct{}

func init() {
	ack.RegisterRoute(&visit{})
}

func (v *visit) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/visit")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageVisitByKey", v.pageVisitByKey)
	_g.Get("/pageSearchElderByKey", v.pageSearchElderByKey)
	_g.Post("/addVisit", v.addVisit)
	_g.Get("/getVisitById", v.getVisitById)
	_g.Put("/editVisit", v.editVisit)
	_g.Put("/recordLeave", v.recordLeave)
	_g.Delete("/deleteVisit", v.deleteVisit)
}

func (v *visit) pageVisitByKey(ctx ack.Context) {
	ack.Get(ctx, service.Visit.PageVisitByKey)
}

func (v *visit) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.PageSearchElderByKey)
}

func (v *visit) addVisit(ctx ack.Context) {
	ack.Post(ctx, service.Visit.AddVisit)
}

func (v *visit) getVisitById(ctx ack.Context) {
	ack.Get(ctx, service.Visit.GetVisitById)
}

func (v *visit) editVisit(ctx ack.Context) {
	ack.Put(ctx, service.Visit.EditVisit)
}

func (v *visit) recordLeave(ctx ack.Context) {
	ack.Put(ctx, service.Visit.RecordLeave)
}

func (v *visit) deleteVisit(ctx ack.Context) {
	ack.Delete(ctx, service.Visit.DeleteVisit)
}
