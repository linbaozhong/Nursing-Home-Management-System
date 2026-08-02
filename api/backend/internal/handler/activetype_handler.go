package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type activetype struct{}

func init() {
	ack.RegisterRoute(&activetype{})
}

func (a *activetype) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/activeType")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageActiveTypeByKey", a.pageActiveTypeByKey)
	_g.Get("/getActiveTypeById", a.getActiveTypeById)
	_g.Post("/addActiveType", a.addActiveType)
	_g.Post("/editActiveType", a.editActiveType)
	_g.Post("/deleteActiveType", a.deleteActiveType)
}

// 分页查询活动分类
// @Summary 分页查询活动分类
// @Tags 活动分类
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageActiveTypeByKeyQuery true "PageActiveTypeByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /activeType/pageActiveTypeByKey [get]
func (a *activetype) pageActiveTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.ActiveType.PageActiveTypeByKey)
}

// 获取活动分类
// @Summary 获取活动分类
// @Tags 活动分类
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /activeType/getActiveTypeById [get]
func (a *activetype) getActiveTypeById(ctx ack.Context) {
	ack.Get(ctx, service.ActiveType.GetActiveTypeById)
}

// 新增活动分类
// @Summary 新增活动分类
// @Tags 活动分类
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddActiveTypeQuery true "AddActiveTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /activeType/addActiveType [post]
func (a *activetype) addActiveType(ctx ack.Context) {
	ack.Post(ctx, service.ActiveType.AddActiveType)
}

// 编辑活动分类
// @Summary 编辑活动分类
// @Tags 活动分类
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateActiveTypeQuery true "OperateActiveTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /activeType/editActiveType [post]
func (a *activetype) editActiveType(ctx ack.Context) {
	ack.Post(ctx, service.ActiveType.EditActiveType)
}

// 删除活动分类
// @Summary 删除活动分类
// @Tags 活动分类
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /activeType/deleteActiveType [post]
func (a *activetype) deleteActiveType(ctx ack.Context) {
	ack.Post(ctx, service.ActiveType.DeleteActiveType)
}
