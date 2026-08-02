package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type outward struct{}

func init() {
	ack.RegisterRoute(&outward{})
}

func (o *outward) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/outward")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageOutwardByKey", o.pageOutwardByKey)
	_g.Get("/getOutwardById", o.getOutwardById)
	_g.Post("/addOutward", o.addOutward)
	_g.Post("/editOutward", o.editOutward)
	_g.Post("/deleteOutward", o.deleteOutward)
	_g.Get("/pageSearchElderByKey", o.pageSearchElderByKey)
	_g.Post("/delayReturn", o.delayReturn)
	_g.Post("/recordReturn", o.recordReturn)
}

// 分页查询出库
// @Summary 分页查询出库
// @Tags 出库
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageOutwardByKeyQuery true "PageOutwardByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outward/pageOutwardByKey [get]
func (o *outward) pageOutwardByKey(ctx ack.Context) {
	ack.Get(ctx, service.Outward.PageOutwardByKey)
}

// 获取出库
// @Summary 获取出库
// @Tags 出库
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /outward/getOutwardById [get]
func (o *outward) getOutwardById(ctx ack.Context) {
	ack.Get(ctx, service.Outward.GetOutwardById)
}

// 新增出库
// @Summary 新增出库
// @Tags 出库
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddOutwardQuery true "AddOutwardQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outward/addOutward [post]
func (o *outward) addOutward(ctx ack.Context) {
	ack.Post(ctx, service.Outward.AddOutward)
}

// 编辑出库
// @Summary 编辑出库
// @Tags 出库
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditOutwardQuery true "EditOutwardQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outward/editOutward [post]
func (o *outward) editOutward(ctx ack.Context) {
	ack.Post(ctx, service.Outward.EditOutward)
}

// 删除出库
// @Summary 删除出库
// @Tags 出库
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /outward/deleteOutward [post]
func (o *outward) deleteOutward(ctx ack.Context) {
	ack.Post(ctx, service.Outward.DeleteOutward)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 出库
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outward/pageSearchElderByKey [get]
func (o *outward) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Outward.PageSearchElderByKey)
}

// 延迟归还
// @Summary 延迟归还
// @Tags 出库
// @Accept application/json
// @Produce application/json
// @Param data body dto.DelayReturnQuery true "DelayReturnQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outward/delayReturn [post]
func (o *outward) delayReturn(ctx ack.Context) {
	ack.Post(ctx, service.Outward.DelayReturn)
}

// 记录归还
// @Summary 记录归还
// @Tags 出库
// @Accept application/json
// @Produce application/json
// @Param data body dto.RecordReturnQuery true "RecordReturnQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outward/recordReturn [post]
func (o *outward) recordReturn(ctx ack.Context) {
	ack.Post(ctx, service.Outward.RecordReturn)
}
