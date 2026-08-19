package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/editIntent", i.editIntent)
	_g.Post("/deleteIntent", i.deleteIntent)

	// 标签
	_g.Get("/listLabel", i.listLabel)
	_g.Get("/pageSearchElderByKey", i.pageSearchElderByKey)
	_g.Get("/getElderLabelById", i.getElderLabelById)
	_g.Get("/getEditElderLabelById", i.getEditElderLabelById)
	_g.Post("/editElderLabel", i.editElderLabel)

	// 回访计划
	_g.Get("/pageVisitPlan", i.visitPlanByKey)
	_g.Post("/addVisitPlan", i.addVisitPlan)
	_g.Post("/executeVisitPlan", i.executeVisitPlan)
	_g.Post("/deleteVisitPlan", i.deleteVisitPlan)

	// 沟通记录
	_g.Get("/pageCommunicationRecord", i.pageCommunicationRecord)
	_g.Post("/addCommunicationRecord", i.addCommunicationRecord)
	_g.Post("/editCommunicationRecord", i.editCommunicationRecord)
	_g.Post("/deleteCommunicationRecord", i.deleteCommunicationRecord)
}

// 分页查询Intent
// @Summary 分页查询Intent
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageIntentionByKeyQuery true "PageIntentionByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/pageIntentByKey [get]
func (i *intention) pageIntentByKey(ctx ack.Context) {
	ack.Get(ctx, service.Intention.PageIntentByKey)
}

// 获取Intent
// @Summary 获取Intent
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/getIntentById [get]
func (i *intention) getIntentById(ctx ack.Context) {
	ack.Get(ctx, service.Intention.GetIntentById)
}

// 新增Intent
// @Summary 新增Intent
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddIntentQuery true "AddIntentQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/addIntent [post]
func (i *intention) addIntent(ctx ack.Context) {
	ack.Post(ctx, service.Intention.AddIntent)
}

// 编辑Intent
// @Summary 编辑Intent
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditIntentQuery true "EditIntentQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/editIntent [post]
func (i *intention) editIntent(ctx ack.Context) {
	ack.Post(ctx, service.Intention.EditIntent)
}

// 删除Intent
// @Summary 删除Intent
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/deleteIntent [post]
func (i *intention) deleteIntent(ctx ack.Context) {
	ack.Post(ctx, service.Intention.DeleteIntent)
}

// 查询标签
// @Summary 查询标签
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data query dto.ListLabelQuery true "ListLabelQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/listLabel [get]
func (i *intention) listLabel(ctx ack.Context) {
	ack.Get(ctx, service.Intention.ListLabel)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/pageSearchElderByKey [get]
func (i *intention) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Intention.PageSearchElderByKey)
}

// 获取老人标签
// @Summary 获取老人标签
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data query dto.GetElderLabelByIdQuery true "GetElderLabelByIdQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/getElderLabelById [get]
func (i *intention) getElderLabelById(ctx ack.Context) {
	ack.Get(ctx, service.Intention.GetElderLabelById)
}

// 获取Edit老人标签
// @Summary 获取Edit老人标签
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/getEditElderLabelById [get]
func (i *intention) getEditElderLabelById(ctx ack.Context) {
	ack.Get(ctx, service.Intention.GetEditElderLabelById)
}

// 编辑老人标签
// @Summary 编辑老人标签
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditElderLabelQuery true "EditElderLabelQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/editElderLabel [post]
func (i *intention) editElderLabel(ctx ack.Context) {
	ack.Post(ctx, service.Intention.EditElderLabel)
}

// 计划
// @Summary 计划
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageVisitPlanQuery true "PageVisitPlanQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/pageVisitPlan [get]
func (i *intention) visitPlanByKey(ctx ack.Context) {
	ack.Get(ctx, service.Intention.PageVisitPlan)
}

// 新增回访计划
// @Summary 新增回访计划
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddVisitPlanQuery true "AddVisitPlanQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/addVisitPlan [post]
func (i *intention) addVisitPlan(ctx ack.Context) {
	ack.Post(ctx, service.Intention.AddVisitPlan)
}

// 执行回访计划
// @Summary 执行回访计划
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/executeVisitPlan [post]
func (i *intention) executeVisitPlan(ctx ack.Context) {
	ack.Post(ctx, service.Intention.ExecuteVisitPlan)
}

// 删除回访计划
// @Summary 删除回访计划
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/deleteVisitPlan [post]
func (i *intention) deleteVisitPlan(ctx ack.Context) {
	ack.Post(ctx, service.Intention.DeleteVisitPlan)
}

// 分页查询沟通记录
// @Summary 分页查询沟通记录
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageCommunicationRecordQuery true "PageCommunicationRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/pageCommunicationRecord [get]
func (i *intention) pageCommunicationRecord(ctx ack.Context) {
	ack.Get(ctx, service.Intention.PageCommunicationRecord)
}

// 新增沟通记录
// @Summary 新增沟通记录
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddCommunicationRecordQuery true "AddCommunicationRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/addCommunicationRecord [post]
func (i *intention) addCommunicationRecord(ctx ack.Context) {
	ack.Post(ctx, service.Intention.AddCommunicationRecord)
}

// 编辑沟通记录
// @Summary 编辑沟通记录
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditCommunicationRecordQuery true "EditCommunicationRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/editCommunicationRecord [post]
func (i *intention) editCommunicationRecord(ctx ack.Context) {
	ack.Post(ctx, service.Intention.EditCommunicationRecord)
}

// 删除沟通记录
// @Summary 删除沟通记录
// @Tags 意向客户
// @Accept application/json
// @Produce application/json
// @Param data body dto.DeleteCommunicationRecordQuery true "DeleteCommunicationRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /intention/deleteCommunicationRecord [post]
func (i *intention) deleteCommunicationRecord(ctx ack.Context) {
	ack.Post(ctx, service.Intention.DeleteCommunicationRecord)
}
