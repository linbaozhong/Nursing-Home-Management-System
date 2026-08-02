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
	_g.Post("/editStaff", s.editStaff)
	_g.Post("/leaveStaff", s.leaveStaff)
}

// 获取角色
// @Summary 获取角色
// @Tags 员工
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /staff/getRole [get]
func (s *staff) getRole(ctx ack.Context) {
	ack.Get(ctx, service.Staff.GetRole)
}

// 分页查询员工
// @Summary 分页查询员工
// @Tags 员工
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageStaffByKeyQuery true "PageStaffByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /staff/pageStaffByKey [get]
func (s *staff) pageStaffByKey(ctx ack.Context) {
	ack.Get(ctx, service.Staff.PageStaffByKey)
}

// 新增员工
// @Summary 新增员工
// @Tags 员工
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateStaffQuery true "OperateStaffQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /staff/addStaff [post]
func (s *staff) addStaff(ctx ack.Context) {
	ack.Post(ctx, service.Staff.AddStaff)
}

// 获取员工
// @Summary 获取员工
// @Tags 员工
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /staff/getStaffById [get]
func (s *staff) getStaffById(ctx ack.Context) {
	ack.Get(ctx, service.Staff.GetStaffById)
}

// 编辑员工
// @Summary 编辑员工
// @Tags 员工
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateStaffQuery true "OperateStaffQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /staff/editStaff [post]
func (s *staff) editStaff(ctx ack.Context) {
	ack.Post(ctx, service.Staff.EditStaff)
}

// 员工
// @Summary 员工
// @Tags 员工
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /staff/leaveStaff [post]
func (s *staff) leaveStaff(ctx ack.Context) {
	ack.Post(ctx, service.Staff.LeaveStaff)
}
