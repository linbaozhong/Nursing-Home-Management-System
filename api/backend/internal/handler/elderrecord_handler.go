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
	_g.Post("/editElderRecord", e.editElderRecord)
	_g.Post("/deleteElderRecord", e.deleteElderRecord)
	_g.Get("/pageSearchElderByKey", e.pageSearchElderByKey)
	_g.Get("/pageSearchEmergencyContactByKey", e.pageSearchEmergencyContactByKey)
	_g.Get("/pageLabelByKey", e.pageLabelByKey)
	_g.Post("/editElderLabel", e.editElderLabel)
}

// 分页查询老人记录
// @Summary 分页查询老人记录
// @Tags 老人档案
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageElderRecordByKeyQuery true "PageElderRecordByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /elderRecord/pageElderRecordByKey [get]
func (e *elderrecord) pageElderRecordByKey(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.PageElderRecordByKey)
}

// 获取老人记录
// @Summary 获取老人记录
// @Tags 老人档案
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /elderRecord/getElderRecordById [get]
func (e *elderrecord) getElderRecordById(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.GetElderRecordById)
}

// 新增老人档案
// @Summary 新增老人档案
// @Tags 老人档案
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddElderRecordQuery true "AddElderRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /elderRecord/addElderRecord [post]
func (e *elderrecord) addElderRecord(ctx ack.Context) {
	ack.Post(ctx, service.ElderRecord.AddElderRecord)
}

// 编辑老人档案
// @Summary 编辑老人档案
// @Tags 老人档案
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditElderRecordQuery true "EditElderRecordQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /elderRecord/editElderRecord [post]
func (e *elderrecord) editElderRecord(ctx ack.Context) {
	ack.Post(ctx, service.ElderRecord.EditElderRecord)
}

// 删除老人档案
// @Summary 删除老人档案
// @Tags 老人档案
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /elderRecord/deleteElderRecord [post]
func (e *elderrecord) deleteElderRecord(ctx ack.Context) {
	ack.Post(ctx, service.ElderRecord.DeleteElderRecord)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 老人档案
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /elderRecord/pageSearchElderByKey [get]
func (e *elderrecord) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.PageSearchElderByKey)
}

// 分页查询EmergencyContact
// @Summary 分页查询EmergencyContact
// @Tags 老人档案
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchEmergencyContactByKeyQuery true "PageSearchEmergencyContactByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /elderRecord/pageSearchEmergencyContactByKey [get]
func (e *elderrecord) pageSearchEmergencyContactByKey(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.PageSearchEmergencyContactByKey)
}

// 分页查询标签
// @Summary 分页查询标签
// @Tags 老人档案
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageLabelByKeyQuery true "PageLabelByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /elderRecord/pageLabelByKey [get]
func (e *elderrecord) pageLabelByKey(ctx ack.Context) {
	ack.Get(ctx, service.ElderRecord.PageLabelByKey)
}

// 编辑老人标签
// @Summary 编辑老人标签
// @Tags 老人档案
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditElderLabelQuery true "EditElderLabelQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /elderRecord/editElderLabel [post]
func (e *elderrecord) editElderLabel(ctx ack.Context) {
	ack.Post(ctx, service.ElderRecord.EditElderLabel)
}
