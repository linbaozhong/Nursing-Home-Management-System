package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/editNurseReserve", n.editNurseReserve)
	_g.Post("/deleteNurseReserve", n.deleteNurseReserve)
	_g.Get("/pageSearchElderByKey", n.pageSearchElderByKey)
	_g.Get("/listNurseStaff", n.listNurseStaff)
}

// 分页查询护理预约
// @Summary 分页查询护理预约
// @Tags 护理预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageNurseReserveByKeyQuery true "PageNurseReserveByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseReserve/pageNurseReserveByKey [get]
func (n *nursereserve) pageNurseReserveByKey(ctx ack.Context) {
	ack.Get(ctx, service.NurseReserve.PageNurseReserveByKey)
}

// 获取护理预约预约And老人
// @Summary 获取护理预约预约And老人
// @Tags 护理预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.GetNurseReserveByReserveIdAndElderIdQuery true "GetNurseReserveByReserveIdAndElderIdQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseReserve/getNurseReserveByReserveIdAndElderId [get]
func (n *nursereserve) getNurseReserveByReserveIdAndElderId(ctx ack.Context) {
	ack.Get(ctx, service.NurseReserve.GetNurseReserveByReserveIdAndElderId)
}

// 新增护理预约
// @Summary 新增护理预约
// @Tags 护理预约
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddNurseReserveQuery true "AddNurseReserveQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseReserve/addNurseReserve [post]
func (n *nursereserve) addNurseReserve(ctx ack.Context) {
	ack.Post(ctx, service.NurseReserve.AddNurseReserve)
}

// 编辑护理预约
// @Summary 编辑护理预约
// @Tags 护理预约
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditNurseReserveQuery true "EditNurseReserveQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseReserve/editNurseReserve [post]
func (n *nursereserve) editNurseReserve(ctx ack.Context) {
	ack.Post(ctx, service.NurseReserve.EditNurseReserve)
}

// 删除护理预约
// @Summary 删除护理预约
// @Tags 护理预约
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseReserve/deleteNurseReserve [post]
func (n *nursereserve) deleteNurseReserve(ctx ack.Context) {
	ack.Post(ctx, service.NurseReserve.DeleteNurseReserve)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 护理预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseReserve/pageSearchElderByKey [get]
func (n *nursereserve) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.NurseReserve.PageSearchElderByKey)
}

// 查询护理员工
// @Summary 查询护理员工
// @Tags 护理预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /nurseReserve/listNurseStaff [get]
func (n *nursereserve) listNurseStaff(ctx ack.Context) {
	ack.Get(ctx, service.NurseReserve.ListNurseStaff)
}
