package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
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
	_g.Put("/auditOutboundRecord", o.auditOutboundRecord)
	_g.Delete("/deleteOutboundRecord", o.deleteOutboundRecord)
}

func (o *outboundrecord) pageOutboundRecordByKey(ctx ack.Context) {
	ack.Get(ctx, service.OutboundRecord.PageOutboundRecordByKey)
}

func (o *outboundrecord) getOutboundRecordById(ctx ack.Context) {
	ack.Get(ctx, service.OutboundRecord.GetOutboundRecordById)
}

func (o *outboundrecord) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.OutboundRecord.PageSearchElderByKey)
}

func (o *outboundrecord) pageWarehouseMaterialByKey(ctx ack.Context) {
	ack.Get(ctx, service.OutboundRecord.PageWarehouseMaterialByKey)
}

func (o *outboundrecord) addOutboundRecord(ctx ack.Context) {
	ack.Post(ctx, service.OutboundRecord.AddOutboundRecord)
}

func (o *outboundrecord) auditOutboundRecord(ctx ack.Context) {
	ack.Put(ctx, service.OutboundRecord.AuditOutboundRecord)
}

func (o *outboundrecord) deleteOutboundRecord(ctx ack.Context) {
	ack.Delete(ctx, service.OutboundRecord.DeleteOutboundRecord)
}
