package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type intention struct{}

func init() {
	ack.RegisterRoute(&intention{})
}

func (i *intention) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/intention")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageIntentByKey", i.pageIntentByKey)
	_g.Get("/getIntentById", i.getIntentById)
	_g.Post("/addIntent", i.addIntent)
	_g.Put("/editIntent", i.editIntent)
	_g.Delete("/deleteIntent", i.deleteIntent)

	// 标签
	_g.Get("/listLabel", i.listLabel)
	_g.Get("/pageSearchElderByKey", i.pageSearchElderByKey)
	_g.Get("/getElderLabelById", i.getElderLabelById)
	_g.Get("/getEditElderLabelById", i.getEditElderLabelById)
	_g.Put("/editElderLabel", i.editElderLabel)

	// 回访计划
	_g.Get("/pageVisitPlan", i.visitPlanByKey)
	_g.Post("/addVisitPlan", i.addVisitPlan)
	_g.Put("/executeVisitPlan", i.executeVisitPlan)
	_g.Delete("/deleteVisitPlan", i.deleteVisitPlan)

	// 沟通记录
	_g.Get("/pageCommunicationRecord", i.pageCommunicationRecord)
	_g.Post("/addCommunicationRecord", i.addCommunicationRecord)
	_g.Put("/editCommunicationRecord", i.editCommunicationRecord)
	_g.Delete("/deleteCommunicationRecord", i.deleteCommunicationRecord)
}

func (i *intention) pageIntentByKey(ctx ack.Context) {
	ack.Get(ctx, service.Intention.PageIntentByKey)
}

func (i *intention) getIntentById(ctx ack.Context) {
	ack.Get(ctx, service.Intention.GetIntentById)
}

func (i *intention) addIntent(ctx ack.Context) {
	ack.Post(ctx, service.Intention.AddIntent)
}

func (i *intention) editIntent(ctx ack.Context) {
	ack.Put(ctx, service.Intention.EditIntent)
}

func (i *intention) deleteIntent(ctx ack.Context) {
	ack.Delete(ctx, service.Intention.DeleteIntent)
}

func (i *intention) listLabel(ctx ack.Context) {
	ack.Get(ctx, service.Intention.ListLabel)
}

func (i *intention) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Intention.PageSearchElderByKey)
}

func (i *intention) getElderLabelById(ctx ack.Context) {
	ack.Get(ctx, service.Intention.GetElderLabelById)
}

func (i *intention) getEditElderLabelById(ctx ack.Context) {
	ack.Get(ctx, service.Intention.GetEditElderLabelById)
}

func (i *intention) editElderLabel(ctx ack.Context) {
	ack.Put(ctx, service.Intention.EditElderLabel)
}

func (i *intention) visitPlanByKey(ctx ack.Context) {
	ack.Get(ctx, service.Intention.PageVisitPlan)
}

func (i *intention) addVisitPlan(ctx ack.Context) {
	ack.Post(ctx, service.Intention.AddVisitPlan)
}

func (i *intention) executeVisitPlan(ctx ack.Context) {
	ack.Put(ctx, service.Intention.ExecuteVisitPlan)
}

func (i *intention) deleteVisitPlan(ctx ack.Context) {
	ack.Delete(ctx, service.Intention.DeleteVisitPlan)
}

func (i *intention) pageCommunicationRecord(ctx ack.Context) {
	ack.Get(ctx, service.Intention.PageCommunicationRecord)
}

func (i *intention) addCommunicationRecord(ctx ack.Context) {
	ack.Post(ctx, service.Intention.AddCommunicationRecord)
}

func (i *intention) editCommunicationRecord(ctx ack.Context) {
	ack.Put(ctx, service.Intention.EditCommunicationRecord)
}

func (i *intention) deleteCommunicationRecord(ctx ack.Context) {
	ack.Delete(ctx, service.Intention.DeleteCommunicationRecord)
}
