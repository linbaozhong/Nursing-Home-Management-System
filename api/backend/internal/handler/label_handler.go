package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type label struct{}

func init() {
	ack.RegisterRoute(&label{})
}

func (l *label) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/label")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageLabelByKey", l.pageLabelByKey)
	_g.Get("/getLabelById", l.getLabelById)
	_g.Post("/addLabel", l.addLabel)
	_g.Put("/editLabel", l.editLabel)
	_g.Delete("/deleteLabel", l.deleteLabel)
	_g.Get("/pageLabelTypeByKey", l.pageLabelTypeByKey)
	_g.Get("/getLabelTypeById", l.getLabelTypeById)
	_g.Post("/addLabelType", l.addLabelType)
	_g.Put("/editLabelType", l.editLabelType)
	_g.Delete("/deleteLabelType", l.deleteLabelType)
}

func (l *label) pageLabelByKey(ctx ack.Context) {
	ack.Get(ctx, service.Label.PageLabelByKey)
}

func (l *label) getLabelById(ctx ack.Context) {
	ack.Get(ctx, service.Label.GetLabelById)
}

func (l *label) addLabel(ctx ack.Context) {
	ack.Post(ctx, service.Label.AddLabel)
}

func (l *label) editLabel(ctx ack.Context) {
	ack.Put(ctx, service.Label.EditLabel)
}

func (l *label) deleteLabel(ctx ack.Context) {
	ack.Delete(ctx, service.Label.DeleteLabel)
}

func (l *label) pageLabelTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.Label.PageLabelTypeByKey)
}

func (l *label) getLabelTypeById(ctx ack.Context) {
	ack.Get(ctx, service.Label.GetLabelTypeById)
}

func (l *label) addLabelType(ctx ack.Context) {
	ack.Post(ctx, service.Label.AddLabelType)
}

func (l *label) editLabelType(ctx ack.Context) {
	ack.Put(ctx, service.Label.EditLabelType)
}

func (l *label) deleteLabelType(ctx ack.Context) {
	ack.Delete(ctx, service.Label.DeleteLabelType)
}
