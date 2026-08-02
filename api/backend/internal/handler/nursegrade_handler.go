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
	_g.Post("/editNurseGrade", n.editNurseGrade)
	_g.Post("/deleteNurseGrade", n.deleteNurseGrade)
	_g.Get("/pageNurseByKey", n.pageNurseByKey)
	_g.Get("/getNurseById", n.getNurseById)
	_g.Post("/addNurse", n.addNurse)
	_g.Post("/editNurse", n.editNurse)
	_g.Post("/deleteNurse", n.deleteNurse)
}

// 分页查询护理等级
// @Summary 分页查询护理等级
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageNurseGradeByKeyQuery true "PageNurseGradeByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/pageNurseGradeByKey [get]
func (n *nursegrade) pageNurseGradeByKey(ctx ack.Context) {
	ack.Get(ctx, service.NurseGrade.PageNurseGradeByKey)
}

// 获取护理等级
// @Summary 获取护理等级
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/getNurseGradeById [get]
func (n *nursegrade) getNurseGradeById(ctx ack.Context) {
	ack.Get(ctx, service.NurseGrade.GetNurseGradeById)
}

// 新增护理等级
// @Summary 新增护理等级
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddNurseGradeQuery true "AddNurseGradeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/addNurseGrade [post]
func (n *nursegrade) addNurseGrade(ctx ack.Context) {
	ack.Post(ctx, service.NurseGrade.AddNurseGrade)
}

// 编辑护理等级
// @Summary 编辑护理等级
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditNurseGradeQuery true "EditNurseGradeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/editNurseGrade [post]
func (n *nursegrade) editNurseGrade(ctx ack.Context) {
	ack.Post(ctx, service.NurseGrade.EditNurseGrade)
}

// 删除护理等级
// @Summary 删除护理等级
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/deleteNurseGrade [post]
func (n *nursegrade) deleteNurseGrade(ctx ack.Context) {
	ack.Post(ctx, service.NurseGrade.DeleteNurseGrade)
}

// 分页查询护理
// @Summary 分页查询护理
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageNurseByKeyQuery true "PageNurseByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/pageNurseByKey [get]
func (n *nursegrade) pageNurseByKey(ctx ack.Context) {
	ack.Get(ctx, service.NurseGrade.PageNurseByKey)
}

// 获取护理
// @Summary 获取护理
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/getNurseById [get]
func (n *nursegrade) getNurseById(ctx ack.Context) {
	ack.Get(ctx, service.NurseGrade.GetNurseById)
}

// 新增护理
// @Summary 新增护理
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddNurseQuery true "AddNurseQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/addNurse [post]
func (n *nursegrade) addNurse(ctx ack.Context) {
	ack.Post(ctx, service.NurseGrade.AddNurse)
}

// 编辑护理
// @Summary 编辑护理
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditNurseQuery true "EditNurseQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/editNurse [post]
func (n *nursegrade) editNurse(ctx ack.Context) {
	ack.Post(ctx, service.NurseGrade.EditNurse)
}

// 删除护理
// @Summary 删除护理
// @Tags 护理等级
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseGrade/deleteNurse [post]
func (n *nursegrade) deleteNurse(ctx ack.Context) {
	ack.Post(ctx, service.NurseGrade.DeleteNurse)
}
