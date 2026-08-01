package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type retreatapply struct{}

func init() {
	ack.RegisterRoute(&retreatapply{})
}

func (r *retreatapply) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/retreatApply")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageRetreatApplyByKey", r.pageRetreatApplyByKey)
	_g.Get("/getRetreatApplyById", r.getRetreatApplyById)
	_g.Post("/addRetreatApply", r.addRetreatApply)
	_g.Put("/editRetreatApply", r.editRetreatApply)
	_g.Delete("/deleteRetreatApply", r.deleteRetreatApply)
	_g.Get("/pageSearchElderByKey", r.pageSearchElderByKey)
}

func (r *retreatapply) pageRetreatApplyByKey(ctx ack.Context) {
	ack.Get(ctx, service.RetreatApply.PageRetreatApplyByKey)
}

func (r *retreatapply) getRetreatApplyById(ctx ack.Context) {
	ack.Get(ctx, service.RetreatApply.GetRetreatApplyById)
}

func (r *retreatapply) addRetreatApply(ctx ack.Context) {
	ack.Post(ctx, service.RetreatApply.AddRetreatApply)
}

func (r *retreatapply) editRetreatApply(ctx ack.Context) {
	ack.Put(ctx, service.RetreatApply.EditRetreatApply)
}

func (r *retreatapply) deleteRetreatApply(ctx ack.Context) {
	ack.Delete(ctx, service.RetreatApply.DeleteRetreatApply)
}

func (r *retreatapply) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.RetreatApply.PageSearchElderByKey)
}
