package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type elderrecord struct{}

func init() {
	ack.RegisterRoute(&elderrecord{})
}

func (e *elderrecord) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/elderRecord")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageElderRecordByKey", e.pageElderRecordByKey)
	_g.Get("/getElderRecordById", e.getElderRecordById)
	_g.Post("/addElderRecord", e.addElderRecord)
	_g.Put("/editElderRecord", e.editElderRecord)
	_g.Delete("/deleteElderRecord", e.deleteElderRecord)
	_g.Get("/pageSearchElderByKey", e.pageSearchElderByKey)
	_g.Get("/pageSearchEmergencyContactByKey", e.pageSearchEmergencyContactByKey)
	_g.Get("/pageLabelByKey", e.pageLabelByKey)
	_g.Put("/editElderLabel", e.editElderLabel)
}

func (e *elderrecord) pageElderRecordByKey(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.PageElderRecordByKey)
}

func (e *elderrecord) getElderRecordById(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.GetElderRecordById)
}

func (e *elderrecord) addElderRecord(ctx ack.Context) {
	ack.Post(ctx, service.ElderRecord.AddElderRecord)
}

func (e *elderrecord) editElderRecord(ctx ack.Context) {
	ack.Put(ctx, service.ElderRecord.EditElderRecord)
}

func (e *elderrecord) deleteElderRecord(ctx ack.Context) {
	ack.Delete(ctx, service.ElderRecord.DeleteElderRecord)
}

func (e *elderrecord) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.PageSearchElderByKey)
}

func (e *elderrecord) pageSearchEmergencyContactByKey(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.PageSearchEmergencyContactByKey)
}

func (e *elderrecord) pageLabelByKey(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.PageLabelByKey)
}

func (e *elderrecord) editElderLabel(ctx ack.Context) {
	ack.Put(ctx, service.ElderRecord.EditElderLabel)
}
