package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/editCheckContract", c.editCheckContract)
	_g.Post("/deleteCheckContract", c.deleteCheckContract)
	// 回访计划（Java 中归属 Intention，Go 端合并保留）
	_g.Get("/pageVisitPlan", c.pageVisitPlan)
	_g.Post("/addVisitPlan", c.addVisitPlan)
	_g.Post("/executeVisitPlan", c.executeVisitPlan)
	_g.Post("/deleteVisitPlan", c.deleteVisitPlan)
	// 沟通记录（Java 中归属 Intention，Go 端合并保留）
	_g.Get("/pageCommunicationRecord", c.pageCommunicationRecord)
	_g.Post("/addCommunicationRecord", c.addCommunicationRecord)
	_g.Post("/editCommunicationRecord", c.editCommunicationRecord)
	_g.Post("/deleteCommunicationRecord", c.deleteCommunicationRecord)
}

// 分页查询入住合同
// @Summary 分页查询入住合同
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/pageCheckContractByKey [get]
func (c *checkContract) pageCheckContractByKey(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.PageCheckContractByKey)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/pageSearchElderByKey [get]
func (c *checkContract) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.PageSearchElderByKey)
}

// 查询护理等级
// @Summary 查询护理等级
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/listNurseGrade [get]
func (c *checkContract) listNurseGrade(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.ListNurseGrade)
}

// 查询餐饮套餐
// @Summary 查询餐饮套餐
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/listCateringSet [get]
func (c *checkContract) listCateringSet(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.ListCateringSet)
}

// 获取楼栋树
// @Summary 获取楼栋树
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/getBuildTree [get]
func (c *checkContract) getBuildTree(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.GetBuildTree)
}

// 获取床位
// @Summary 获取床位
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/getBedById [get]
func (c *checkContract) getBedById(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.GetBedById)
}

// 新增入住办理
// @Summary 新增入住办理
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data body dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/addCheckContract [post]
func (c *checkContract) addCheckContract(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.AddCheckContract)
}

// 获取入住合同
// @Summary 获取入住合同
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/getCheckContractById [get]
func (c *checkContract) getCheckContractById(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.GetCheckContractById)
}

// 编辑入住办理
// @Summary 编辑入住办理
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data body dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/editCheckContract [post]
func (c *checkContract) editCheckContract(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.EditCheckContract)
}

// 删除入住办理
// @Summary 删除入住办理
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data body dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/deleteCheckContract [post]
func (c *checkContract) deleteCheckContract(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.DeleteCheckContract)
}

// 分页查询回访计划
// @Summary 分页查询回访计划
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageVisitPlanQuery true "PageVisitPlanQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/pageVisitPlan [get]
func (c *checkContract) pageVisitPlan(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.PageVisitPlan)
}

// 新增回访计划
// @Summary 新增回访计划
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddVisitPlanQuery true "AddVisitPlanQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/addVisitPlan [post]
func (c *checkContract) addVisitPlan(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.AddVisitPlan)
}

// 执行回访计划
// @Summary 执行回访计划
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data body dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/executeVisitPlan [post]
func (c *checkContract) executeVisitPlan(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.ExecuteVisitPlan)
}

// 删除回访计划
// @Summary 删除回访计划
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data body dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/deleteVisitPlan [post]
func (c *checkContract) deleteVisitPlan(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.DeleteVisitPlan)
}

// 分页查询沟通记录
// @Summary 分页查询沟通记录
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageCommunicationRecordQuery true "PageCommunicationRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/pageCommunicationRecord [get]
func (c *checkContract) pageCommunicationRecord(ctx ack.Context) {
	ack.Get(ctx, service.CheckContract.PageCommunicationRecord)
}

// 新增沟通记录
// @Summary 新增沟通记录
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddCommunicationRecordQuery true "AddCommunicationRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/addCommunicationRecord [post]
func (c *checkContract) addCommunicationRecord(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.AddCommunicationRecord)
}

// 编辑沟通记录
// @Summary 编辑沟通记录
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditCommunicationRecordQuery true "EditCommunicationRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/editCommunicationRecord [post]
func (c *checkContract) editCommunicationRecord(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.EditCommunicationRecord)
}

// 删除沟通记录
// @Summary 删除沟通记录
// @Tags 入住办理
// @Accept application/json
// @Produce application/json
// @Param data body dto.DeleteCommunicationRecordQuery true "DeleteCommunicationRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /checkContract/deleteCommunicationRecord [post]
func (c *checkContract) deleteCommunicationRecord(ctx ack.Context) {
	ack.Post(ctx, service.CheckContract.DeleteCommunicationRecord)
}
