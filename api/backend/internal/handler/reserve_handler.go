package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type reserve struct{}

func init() {
	ack.RegisterRoute(&reserve{})
}

func (r *reserve) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/reserve")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageReserveByKey", r.pageReserveByKey)
	_g.Get("/getReserveById", r.getReserveById)
	_g.Post("/addReserve", r.addReserve)
	_g.Put("/editReserve", r.editReserve)
	_g.Delete("/deleteReserve", r.deleteReserve)
	_g.Get("/pageSearchElderByKey", r.pageSearchElderByKey)
	_g.Get("/pageSearchStaffByKey", r.pageSearchStaffByKey)
	_g.Get("/getBuildTree", r.getBuildTree)
	_g.Get("/getReserveByReserveIdAndElderId", r.getReserveByReserveIdAndElderId)
	_g.Put("/refund", r.refund)
	_g.Get("/reserveExpireJob", r.reserveExpireJob)
}

func (r *reserve) pageReserveByKey(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.PageReserveByKey)
}

func (r *reserve) getReserveById(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.GetReserveById)
}

func (r *reserve) addReserve(ctx ack.Context) {
	ack.Post(ctx, service.Reserve.AddReserve)
}

func (r *reserve) editReserve(ctx ack.Context) {
	ack.Put(ctx, service.Reserve.EditReserve)
}

func (r *reserve) deleteReserve(ctx ack.Context) {
	ack.Delete(ctx, service.Reserve.DeleteReserve)
}

func (r *reserve) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.PageSearchElderByKey)
}

func (r *reserve) pageSearchStaffByKey(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.PageSearchStaffByKey)
}

func (r *reserve) getBuildTree(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.GetBuildTree)
}

func (r *reserve) getReserveByReserveIdAndElderId(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.GetReserveByReserveIdAndElderId)
}

func (r *reserve) refund(ctx ack.Context) {
	ack.Put(ctx, service.Reserve.Refund)
}

func (r *reserve) reserveExpireJob(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.ReserveExpireJob)
}
