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
	_g.Post("/editReserve", r.editReserve)
	_g.Post("/deleteReserve", r.deleteReserve)
	_g.Get("/pageSearchElderByKey", r.pageSearchElderByKey)
	_g.Get("/pageSearchStaffByKey", r.pageSearchStaffByKey)
	_g.Get("/getBuildTree", r.getBuildTree)
	_g.Get("/getReserveByReserveIdAndElderId", r.getReserveByReserveIdAndElderId)
	_g.Post("/refund", r.refund)
	_g.Get("/reserveExpireJob", r.reserveExpireJob)
}

// 分页查询预约
// @Summary 分页查询预约
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageReserveByKeyQuery true "PageReserveByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/pageReserveByKey [get]
func (r *reserve) pageReserveByKey(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.PageReserveByKey)
}

// 获取预约
// @Summary 获取预约
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/getReserveById [get]
func (r *reserve) getReserveById(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.GetReserveById)
}

// 新增预约
// @Summary 新增预约
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddReserveQuery true "AddReserveQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/addReserve [post]
func (r *reserve) addReserve(ctx ack.Context) {
	ack.Post(ctx, service.Reserve.AddReserve)
}

// 编辑预约
// @Summary 编辑预约
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditReserveQuery true "EditReserveQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/editReserve [post]
func (r *reserve) editReserve(ctx ack.Context) {
	ack.Post(ctx, service.Reserve.EditReserve)
}

// 删除预约
// @Summary 删除预约
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/deleteReserve [post]
func (r *reserve) deleteReserve(ctx ack.Context) {
	ack.Post(ctx, service.Reserve.DeleteReserve)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/pageSearchElderByKey [get]
func (r *reserve) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.PageSearchElderByKey)
}

// 分页查询员工
// @Summary 分页查询员工
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchStaffByKeyQuery true "PageSearchStaffByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/pageSearchStaffByKey [get]
func (r *reserve) pageSearchStaffByKey(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.PageSearchStaffByKey)
}

// 获取楼栋树
// @Summary 获取楼栋树
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/getBuildTree [get]
func (r *reserve) getBuildTree(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.GetBuildTree)
}

// 获取预约预约And老人
// @Summary 获取预约预约And老人
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/getReserveByReserveIdAndElderId [get]
func (r *reserve) getReserveByReserveIdAndElderId(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.GetReserveByReserveIdAndElderId)
}

// 预约
// @Summary 预约
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/refund [post]
func (r *reserve) refund(ctx ack.Context) {
	ack.Post(ctx, service.Reserve.Refund)
}

// ExpireJob
// @Summary ExpireJob
// @Tags 预约
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /reserve/reserveExpireJob [get]
func (r *reserve) reserveExpireJob(ctx ack.Context) {
	ack.Get(ctx, service.Reserve.ReserveExpireJob)
}
