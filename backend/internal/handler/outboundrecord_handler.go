package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type outboundrecord struct{}

func init() {
	ack.RegisterRoute(&outboundrecord{})
}

func (o *outboundrecord) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/outboundRecord")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageOutboundRecordByKey", o.pageOutboundRecordByKey)
	_g.Get("/getOutboundRecordById", o.getOutboundRecordById)
	_g.Get("/pageSearchElderByKey", o.pageSearchElderByKey)
	_g.Get("/pageWarehouseMaterialByKey", o.pageWarehouseMaterialByKey)
	_g.Post("/addOutboundRecord", o.addOutboundRecord)
	_g.Post("/auditOutboundRecord", o.auditOutboundRecord)
	_g.Post("/deleteOutboundRecord", o.deleteOutboundRecord)
}

// 分页查询出库记录
// @Summary 分页查询出库记录
// @Tags 出库记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageOutboundRecordByKeyQuery true "PageOutboundRecordByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outboundRecord/pageOutboundRecordByKey [get]
func (o *outboundrecord) pageOutboundRecordByKey(ctx ack.Context) {
	ack.Get(ctx, service.OutboundRecord.PageOutboundRecordByKey)
}

// 获取出库记录
// @Summary 获取出库记录
// @Tags 出库记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /outboundRecord/getOutboundRecordById [get]
func (o *outboundrecord) getOutboundRecordById(ctx ack.Context) {
	ack.Get(ctx, service.OutboundRecord.GetOutboundRecordById)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 出库记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outboundRecord/pageSearchElderByKey [get]
func (o *outboundrecord) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.OutboundRecord.PageSearchElderByKey)
}

// 分页查询仓库物料
// @Summary 分页查询仓库物料
// @Tags 出库记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageWarehouseMaterialByKeyQuery true "PageWarehouseMaterialByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outboundRecord/pageWarehouseMaterialByKey [get]
func (o *outboundrecord) pageWarehouseMaterialByKey(ctx ack.Context) {
	ack.Get(ctx, service.OutboundRecord.PageWarehouseMaterialByKey)
}

// 新增出库记录
// @Summary 新增出库记录
// @Tags 出库记录
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddOutboundRecordQuery true "AddOutboundRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outboundRecord/addOutboundRecord [post]
func (o *outboundrecord) addOutboundRecord(ctx ack.Context) {
	ack.Post(ctx, service.OutboundRecord.AddOutboundRecord)
}

// 审核出库记录
// @Summary 审核出库记录
// @Tags 出库记录
// @Accept application/json
// @Produce application/json
// @Param data body dto.AuditOutboundRecordQuery true "AuditOutboundRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /outboundRecord/auditOutboundRecord [post]
func (o *outboundrecord) auditOutboundRecord(ctx ack.Context) {
	ack.Post(ctx, service.OutboundRecord.AuditOutboundRecord)
}

// 删除出库记录
// @Summary 删除出库记录
// @Tags 出库记录
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /outboundRecord/deleteOutboundRecord [post]
func (o *outboundrecord) deleteOutboundRecord(ctx ack.Context) {
	ack.Post(ctx, service.OutboundRecord.DeleteOutboundRecord)
}
