package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/editVisit", v.editVisit)
	_g.Post("/recordLeave", v.recordLeave)
	_g.Post("/deleteVisit", v.deleteVisit)
}

// 分页查询回访
// @Summary 分页查询回访
// @Tags 回访
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageVisitByKeyQuery true "PageVisitByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /visit/pageVisitByKey [get]
func (v *visit) pageVisitByKey(ctx ack.Context) {
	ack.Get(ctx, service.Visit.PageVisitByKey)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 回访
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /visit/pageSearchElderByKey [get]
func (v *visit) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.PageSearchElderByKey)
}

// 新增回访
// @Summary 新增回访
// @Tags 回访
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddVisitQuery true "AddVisitQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /visit/addVisit [post]
func (v *visit) addVisit(ctx ack.Context) {
	ack.Post(ctx, service.Visit.AddVisit)
}

// 获取回访
// @Summary 获取回访
// @Tags 回访
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /visit/getVisitById [get]
func (v *visit) getVisitById(ctx ack.Context) {
	ack.Get(ctx, service.Visit.GetVisitById)
}

// 编辑回访
// @Summary 编辑回访
// @Tags 回访
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditVisitQuery true "EditVisitQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /visit/editVisit [post]
func (v *visit) editVisit(ctx ack.Context) {
	ack.Post(ctx, service.Visit.EditVisit)
}

// 记录离院
// @Summary 记录离院
// @Tags 回访
// @Accept application/json
// @Produce application/json
// @Param data body dto.RecordLeaveQuery true "RecordLeaveQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /visit/recordLeave [post]
func (v *visit) recordLeave(ctx ack.Context) {
	ack.Post(ctx, service.Visit.RecordLeave)
}

// 删除回访
// @Summary 删除回访
// @Tags 回访
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /visit/deleteVisit [post]
func (v *visit) deleteVisit(ctx ack.Context) {
	ack.Post(ctx, service.Visit.DeleteVisit)
}
