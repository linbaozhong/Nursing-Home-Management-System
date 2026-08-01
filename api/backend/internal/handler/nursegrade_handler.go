package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type nursegrade struct{}

func init() {
	ack.RegisterRoute(&nursegrade{})
}

func (n *nursegrade) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/nurseGrade")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageNurseGradeByKey", n.pageNurseGradeByKey)
	_g.Get("/getNurseGradeById", n.getNurseGradeById)
	_g.Post("/addNurseGrade", n.addNurseGrade)
	_g.Put("/editNurseGrade", n.editNurseGrade)
	_g.Delete("/deleteNurseGrade", n.deleteNurseGrade)
	_g.Get("/pageNurseByKey", n.pageNurseByKey)
	_g.Get("/getNurseById", n.getNurseById)
	_g.Post("/addNurse", n.addNurse)
	_g.Put("/editNurse", n.editNurse)
	_g.Delete("/deleteNurse", n.deleteNurse)
}

func (n *nursegrade) pageNurseGradeByKey(ctx ack.Context) {
	ack.Get(ctx, service.NurseGrade.PageNurseGradeByKey)
}

func (n *nursegrade) getNurseGradeById(ctx ack.Context) {
	ack.Get(ctx, service.NurseGrade.GetNurseGradeById)
}

func (n *nursegrade) addNurseGrade(ctx ack.Context) {
	ack.Post(ctx, service.NurseGrade.AddNurseGrade)
}

func (n *nursegrade) editNurseGrade(ctx ack.Context) {
	ack.Put(ctx, service.NurseGrade.EditNurseGrade)
}

func (n *nursegrade) deleteNurseGrade(ctx ack.Context) {
	ack.Delete(ctx, service.NurseGrade.DeleteNurseGrade)
}

func (n *nursegrade) pageNurseByKey(ctx ack.Context) {
	ack.Get(ctx, service.NurseGrade.PageNurseByKey)
}

func (n *nursegrade) getNurseById(ctx ack.Context) {
	ack.Get(ctx, service.NurseGrade.GetNurseById)
}

func (n *nursegrade) addNurse(ctx ack.Context) {
	ack.Post(ctx, service.NurseGrade.AddNurse)
}

func (n *nursegrade) editNurse(ctx ack.Context) {
	ack.Put(ctx, service.NurseGrade.EditNurse)
}

func (n *nursegrade) deleteNurse(ctx ack.Context) {
	ack.Delete(ctx, service.NurseGrade.DeleteNurse)
}
