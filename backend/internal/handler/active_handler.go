package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type active struct{}

func init() {
	ack.RegisterRoute(&active{})
}

func (a *active) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/active")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageActiveByKey", a.pageActiveByKey)
	_g.Get("/getActiveById", a.getActiveById)
	_g.Post("/addActive", a.addActive)
	_g.Post("/editActive", a.editActive)
	_g.Post("/deleteActive", a.deleteActive)
	_g.Get("/pageSearchElderByKey", a.pageSearchElderByKey)
}

// 分页查询活动
// @Summary 分页查询活动
// @Tags 活动
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageActiveByKeyReq true "PageActiveByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /active/pageActiveByKey [get]
func (a *active) pageActiveByKey(ctx ack.Context) {
	ack.Get(ctx, service.Active.PageActiveByKey)
}

// 获取活动
// @Summary 获取活动
// @Tags 活动
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /active/getActiveById [get]
func (a *active) getActiveById(ctx ack.Context) {
	ack.Get(ctx, service.Active.GetActiveById)
}

// 新增活动
// @Summary 新增活动
// @Tags 活动
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateActiveReq true "OperateActiveReq"
// @Success 200 {object} dto.EmptyResp
// @Router /active/addActive [post]
func (a *active) addActive(ctx ack.Context) {
	ack.Post(ctx, service.Active.AddActive)
}

// 编辑活动
// @Summary 编辑活动
// @Tags 活动
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateActiveReq true "OperateActiveReq"
// @Success 200 {object} dto.EmptyResp
// @Router /active/editActive [post]
func (a *active) editActive(ctx ack.Context) {
	ack.Post(ctx, service.Active.EditActive)
}

// 删除活动
// @Summary 删除活动
// @Tags 活动
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /active/deleteActive [post]
func (a *active) deleteActive(ctx ack.Context) {
	ack.Post(ctx, service.Active.DeleteActive)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 活动
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyReq true "PageSearchElderByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /active/pageSearchElderByKey [get]
func (a *active) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Active.PageSearchElderByKey)
}
