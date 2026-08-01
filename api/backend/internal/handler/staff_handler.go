package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type staff struct{}

func init() {
	ack.RegisterRoute(&staff{})
}

func (s *staff) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/staff")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/getRole", s.getRole)
	_g.Get("/pageStaffByKey", s.pageStaffByKey)
	_g.Post("/addStaff", s.addStaff)
	_g.Get("/getStaffById", s.getStaffById)
	_g.Put("/editStaff", s.editStaff)
	_g.Delete("/leaveStaff", s.leaveStaff)
}

func (s *staff) getRole(ctx ack.Context) {
	ack.Get(ctx, service.Staff.GetRole)
}

func (s *staff) pageStaffByKey(ctx ack.Context) {
	ack.Get(ctx, service.Staff.PageStaffByKey)
}

func (s *staff) addStaff(ctx ack.Context) {
	ack.Post(ctx, service.Staff.AddStaff)
}

func (s *staff) getStaffById(ctx ack.Context) {
	ack.Get(ctx, service.Staff.GetStaffById)
}

func (s *staff) editStaff(ctx ack.Context) {
	ack.Put(ctx, service.Staff.EditStaff)
}

func (s *staff) leaveStaff(ctx ack.Context) {
	ack.Delete(ctx, service.Staff.LeaveStaff)
}
