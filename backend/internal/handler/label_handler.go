package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/editLabel", l.editLabel)
	_g.Post("/deleteLabel", l.deleteLabel)
	_g.Get("/pageLabelTypeByKey", l.pageLabelTypeByKey)
	_g.Get("/getLabelTypeById", l.getLabelTypeById)
	_g.Post("/addLabelType", l.addLabelType)
	_g.Post("/editLabelType", l.editLabelType)
	_g.Post("/deleteLabelType", l.deleteLabelType)
}

// 分页查询标签
// @Summary 分页查询标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageLabelByKeyQuery true "PageLabelByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /label/pageLabelByKey [get]
func (l *label) pageLabelByKey(ctx ack.Context) {
	ack.Get(ctx, service.Label.PageLabelByKey)
}

// 获取标签
// @Summary 获取标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /label/getLabelById [get]
func (l *label) getLabelById(ctx ack.Context) {
	ack.Get(ctx, service.Label.GetLabelById)
}

// 新增标签
// @Summary 新增标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddLabelQuery true "AddLabelQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /label/addLabel [post]
func (l *label) addLabel(ctx ack.Context) {
	ack.Post(ctx, service.Label.AddLabel)
}

// 编辑标签
// @Summary 编辑标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditLabelQuery true "EditLabelQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /label/editLabel [post]
func (l *label) editLabel(ctx ack.Context) {
	ack.Post(ctx, service.Label.EditLabel)
}

// 删除标签
// @Summary 删除标签
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /label/deleteLabel [post]
func (l *label) deleteLabel(ctx ack.Context) {
	ack.Post(ctx, service.Label.DeleteLabel)
}

// 分页查询标签分类
// @Summary 分页查询标签分类
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageLabelTypeByKeyQuery true "PageLabelTypeByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /label/pageLabelTypeByKey [get]
func (l *label) pageLabelTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.Label.PageLabelTypeByKey)
}

// 获取标签分类
// @Summary 获取标签分类
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /label/getLabelTypeById [get]
func (l *label) getLabelTypeById(ctx ack.Context) {
	ack.Get(ctx, service.Label.GetLabelTypeById)
}

// 新增标签分类
// @Summary 新增标签分类
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddLabelTypeQuery true "AddLabelTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /label/addLabelType [post]
func (l *label) addLabelType(ctx ack.Context) {
	ack.Post(ctx, service.Label.AddLabelType)
}

// 编辑标签分类
// @Summary 编辑标签分类
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditLabelTypeQuery true "EditLabelTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /label/editLabelType [post]
func (l *label) editLabelType(ctx ack.Context) {
	ack.Post(ctx, service.Label.EditLabelType)
}

// 删除标签分类
// @Summary 删除标签分类
// @Tags 标签
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /label/deleteLabelType [post]
func (l *label) deleteLabelType(ctx ack.Context) {
	ack.Post(ctx, service.Label.DeleteLabelType)
}
