package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type retreatapply struct{}

func init() {
	ack.RegisterRoute(&retreatapply{})
}

func (r *retreatapply) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/retreatApply")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageRetreatApplyByKey", r.pageRetreatApplyByKey)
	_g.Get("/getRetreatApplyById", r.getRetreatApplyById)
	_g.Post("/addRetreatApply", r.addRetreatApply)
	_g.Post("/editRetreatApply", r.editRetreatApply)
	_g.Post("/deleteRetreatApply", r.deleteRetreatApply)
	_g.Get("/pageSearchElderByKey", r.pageSearchElderByKey)
}

// 分页查询退住Apply
// @Summary 分页查询退住Apply
// @Tags 退住申请
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageRetreatApplyByKeyReq true "PageRetreatApplyByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatApply/pageRetreatApplyByKey [get]
func (r *retreatapply) pageRetreatApplyByKey(ctx ack.Context) {
	ack.Get(ctx, service.RetreatApply.PageRetreatApplyByKey)
}

// 获取退住Apply
// @Summary 获取退住Apply
// @Tags 退住申请
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatApply/getRetreatApplyById [get]
func (r *retreatapply) getRetreatApplyById(ctx ack.Context) {
	ack.Get(ctx, service.RetreatApply.GetRetreatApplyById)
}

// 新增退住Apply
// @Summary 新增退住Apply
// @Tags 退住申请
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddRetreatApplyReq true "AddRetreatApplyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatApply/addRetreatApply [post]
func (r *retreatapply) addRetreatApply(ctx ack.Context) {
	ack.Post(ctx, service.RetreatApply.AddRetreatApply)
}

// 编辑退住Apply
// @Summary 编辑退住Apply
// @Tags 退住申请
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditRetreatApplyReq true "EditRetreatApplyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatApply/editRetreatApply [post]
func (r *retreatapply) editRetreatApply(ctx ack.Context) {
	ack.Post(ctx, service.RetreatApply.EditRetreatApply)
}

// 删除退住Apply
// @Summary 删除退住Apply
// @Tags 退住申请
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatApply/deleteRetreatApply [post]
func (r *retreatapply) deleteRetreatApply(ctx ack.Context) {
	ack.Post(ctx, service.RetreatApply.DeleteRetreatApply)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 退住申请
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyReq true "PageSearchElderByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatApply/pageSearchElderByKey [get]
func (r *retreatapply) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.RetreatApply.PageSearchElderByKey)
}
