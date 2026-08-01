package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type checkContract struct{}

func init() {
	ack.RegisterRoute(&checkContract{})
}

func (c *checkContract) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/checkContract")
	_g.Use(lib.AuthMiddleware())

	// 入住办理
	_g.Get("/pageCheckContractByKey", c.pageCheckContractByKey)
	_g.Get("/pageSearchElderByKey", c.pageSearchElderByKey)
	_g.Get("/listNurseGrade", c.listNurseGrade)
	_g.Get("/listCateringSet", c.listCateringSet)
	_g.Get("/getBuildTree", c.getBuildTree)
	_g.Get("/getBedById", c.getBedById)
	_g.Post("/addCheckContract", c.addCheckContract)
	_g.Get("/getCheckContractById", c.getCheckContractById)
	_g.Put("/editCheckContract", c.editCheckContract)
	_g.Delete("/deleteCheckContract", c.deleteCheckContract)
	// 回访计划（Java 中归属 Intention，Go 端合并保留）
	_g.Get("/pageVisitPlan", c.pageVisitPlan)
	_g.Post("/addVisitPlan", c.addVisitPlan)
	_g.Put("/executeVisitPlan", c.executeVisitPlan)
	_g.Delete("/deleteVisitPlan", c.deleteVisitPlan)
	// 沟通记录（Java 中归属 Intention，Go 端合并保留）
	_g.Get("/pageCommunicationRecord", c.pageCommunicationRecord)
	_g.Post("/addCommunicationRecord", c.addCommunicationRecord)
	_g.Put("/editCommunicationRecord", c.editCommunicationRecord)
	_g.Delete("/deleteCommunicationRecord", c.deleteCommunicationRecord)
}

// 入住办理
func (c *checkContract) pageCheckContractByKey(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.PageCheckContractByKey)
}

func (c *checkContract) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.PageSearchElderByKey)
}

func (c *checkContract) listNurseGrade(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.ListNurseGrade)
}

func (c *checkContract) listCateringSet(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.ListCateringSet)
}

func (c *checkContract) getBuildTree(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.GetBuildTree)
}

func (c *checkContract) getBedById(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.GetBedById)
}

func (c *checkContract) addCheckContract(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.AddCheckContract)
}

func (c *checkContract) getCheckContractById(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.GetCheckContractById)
}

func (c *checkContract) editCheckContract(ctx ack.Context) {
	ack.Put(ctx, service.CheckContract.EditCheckContract)
}

func (c *checkContract) deleteCheckContract(ctx ack.Context) {
	ack.Delete(ctx, service.CheckContract.DeleteCheckContract)
}

// 回访计划
func (c *checkContract) pageVisitPlan(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.PageVisitPlan)
}

func (c *checkContract) addVisitPlan(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.AddVisitPlan)
}

func (c *checkContract) executeVisitPlan(ctx ack.Context) {
	ack.Put(ctx, service.CheckContract.ExecuteVisitPlan)
}

func (c *checkContract) deleteVisitPlan(ctx ack.Context) {
	ack.Delete(ctx, service.CheckContract.DeleteVisitPlan)
}

// 沟通记录
func (c *checkContract) pageCommunicationRecord(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.PageCommunicationRecord)
}

func (c *checkContract) addCommunicationRecord(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.AddCommunicationRecord)
}

func (c *checkContract) editCommunicationRecord(ctx ack.Context) {
	ack.Put(ctx, service.CheckContract.EditCommunicationRecord)
}

func (c *checkContract) deleteCommunicationRecord(ctx ack.Context) {
	ack.Delete(ctx, service.CheckContract.DeleteCommunicationRecord)
}
