package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type source struct{}

func init() {
	ack.RegisterRoute(&source{})
}

func (s *source) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/source")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageSourceByKey", s.pageSourceByKey)
	_g.Post("/addSource", s.addSource)
	_g.Get("/getSourceById", s.getSourceById)
	_g.Post("/editSource", s.editSource)
	_g.Post("/deleteSource", s.deleteSource)
}

// 分页查询来源
// @Summary 分页查询来源
// @Tags 来源
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSourceByKeyQuery true "PageSourceByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /source/pageSourceByKey [get]
func (s *source) pageSourceByKey(ctx ack.Context) {
	ack.Get(ctx, service.Source.PageSourceByKey)
}

// 新增来源
// @Summary 新增来源
// @Tags 来源
// @Accept application/json
// @Produce application/json
// @Param data body dto.StringReq true "StringReq"
// @Success 200 {object} dto.EmptyResp
// @Router /source/addSource [post]
func (s *source) addSource(ctx ack.Context) {
	ack.Post(ctx, service.Source.AddSource)
}

// 获取来源
// @Summary 获取来源
// @Tags 来源
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /source/getSourceById [get]
func (s *source) getSourceById(ctx ack.Context) {
	ack.Get(ctx, service.Source.GetSourceById)
}

// 编辑来源
// @Summary 编辑来源
// @Tags 来源
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateSourceQuery true "OperateSourceQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /source/editSource [post]
func (s *source) editSource(ctx ack.Context) {
	ack.Post(ctx, service.Source.EditSource)
}

// 删除来源
// @Summary 删除来源
// @Tags 来源
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /source/deleteSource [post]
func (s *source) deleteSource(ctx ack.Context) {
	ack.Post(ctx, service.Source.DeleteSource)
}
