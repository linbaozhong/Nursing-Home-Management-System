package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type nursereserve struct{}

func init() {
	ack.RegisterRoute(&nursereserve{})
}

func (n *nursereserve) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/nurseReserve")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageNurseReserveByKey", n.pageNurseReserveByKey)
	_g.Get("/getNurseReserveByReserveIdAndElderId", n.getNurseReserveByReserveIdAndElderId)
	_g.Post("/addNurseReserve", n.addNurseReserve)
	_g.Put("/editNurseReserve", n.editNurseReserve)
	_g.Delete("/deleteNurseReserve", n.deleteNurseReserve)
	_g.Get("/pageSearchElderByKey", n.pageSearchElderByKey)
	_g.Get("/listNurseStaff", n.listNurseStaff)
}

func (n *nursereserve) pageNurseReserveByKey(ctx ack.Context) {
	ack.Get(ctx, service.NurseReserve.PageNurseReserveByKey)
}

func (n *nursereserve) getNurseReserveByReserveIdAndElderId(ctx ack.Context) {
	ack.Get(ctx, service.NurseReserve.GetNurseReserveByReserveIdAndElderId)
}

func (n *nursereserve) addNurseReserve(ctx ack.Context) {
	ack.Post(ctx, service.NurseReserve.AddNurseReserve)
}

func (n *nursereserve) editNurseReserve(ctx ack.Context) {
	ack.Put(ctx, service.NurseReserve.EditNurseReserve)
}

func (n *nursereserve) deleteNurseReserve(ctx ack.Context) {
	ack.Delete(ctx, service.NurseReserve.DeleteNurseReserve)
}

func (n *nursereserve) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.NurseReserve.PageSearchElderByKey)
}

func (n *nursereserve) listNurseStaff(ctx ack.Context) {
	ack.Get(ctx, service.NurseReserve.ListNurseStaff)
}
