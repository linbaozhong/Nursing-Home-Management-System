package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type accident struct{}

func init() {
	ack.RegisterRoute(&accident{})
}

func (a *accident) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/accident")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageAccidentByKey", a.pageAccidentByKey)
	_g.Get("/getAccidentById", a.getAccidentById)
	_g.Post("/addAccident", a.addAccident)
	_g.Post("/editAccident", a.editAccident)
	_g.Post("/deleteAccident", a.deleteAccident)
}

// 分页查询事故
// @Summary 分页查询事故
// @Tags 事故记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageAccidentByKeyReq true "PageAccidentByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /accident/pageAccidentByKey [get]
func (a *accident) pageAccidentByKey(ctx ack.Context) {
	ack.Get(ctx, service.Accident.PageAccidentByKey)
}

// 获取事故
// @Summary 获取事故
// @Tags 事故记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /accident/getAccidentById [get]
func (a *accident) getAccidentById(ctx ack.Context) {
	ack.Get(ctx, service.Accident.GetAccidentById)
}

// 新增事故
// @Summary 新增事故
// @Tags 事故记录
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddAccidentReq true "AddAccidentReq"
// @Success 200 {object} dto.EmptyResp
// @Router /accident/addAccident [post]
func (a *accident) addAccident(ctx ack.Context) {
	ack.Post(ctx, service.Accident.AddAccident)
}

// 编辑事故
// @Summary 编辑事故
// @Tags 事故记录
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditAccidentReq true "EditAccidentReq"
// @Success 200 {object} dto.EmptyResp
// @Router /accident/editAccident [post]
func (a *accident) editAccident(ctx ack.Context) {
	ack.Post(ctx, service.Accident.EditAccident)
}

// 删除事故
// @Summary 删除事故
// @Tags 事故记录
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /accident/deleteAccident [post]
func (a *accident) deleteAccident(ctx ack.Context) {
	ack.Post(ctx, service.Accident.DeleteAccident)
}
