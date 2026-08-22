package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type cateringset struct{}

func init() {
	ack.RegisterRoute(&cateringset{})
}

func (c *cateringset) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/cateringSet")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageCateringSetByKey", c.pageCateringSetByKey)
	_g.Get("/getCateringSetById", c.getCateringSetById)
	_g.Post("/addCateringSet", c.addCateringSet)
	_g.Post("/editCateringSet", c.editCateringSet)
	_g.Post("/deleteCateringSet", c.deleteCateringSet)
}

// 分页查询餐饮套餐
// @Summary 分页查询餐饮套餐
// @Tags 餐饮套餐
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageCateringSetByKeyReq true "PageCateringSetByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /cateringSet/pageCateringSetByKey [get]
func (c *cateringset) pageCateringSetByKey(ctx ack.Context) {
	ack.Get(ctx, service.CateringSet.PageCateringSetByKey)
}

// 获取餐饮套餐
// @Summary 获取餐饮套餐
// @Tags 餐饮套餐
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /cateringSet/getCateringSetById [get]
func (c *cateringset) getCateringSetById(ctx ack.Context) {
	ack.Get(ctx, service.CateringSet.GetCateringSetById)
}

// 新增餐饮套餐
// @Summary 新增餐饮套餐
// @Tags 餐饮套餐
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateCateringSetReq true "OperateCateringSetReq"
// @Success 200 {object} dto.EmptyResp
// @Router /cateringSet/addCateringSet [post]
func (c *cateringset) addCateringSet(ctx ack.Context) {
	ack.Post(ctx, service.CateringSet.AddCateringSet)
}

// 编辑餐饮套餐
// @Summary 编辑餐饮套餐
// @Tags 餐饮套餐
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateCateringSetReq true "OperateCateringSetReq"
// @Success 200 {object} dto.EmptyResp
// @Router /cateringSet/editCateringSet [post]
func (c *cateringset) editCateringSet(ctx ack.Context) {
	ack.Post(ctx, service.CateringSet.EditCateringSet)
}

// 删除餐饮套餐
// @Summary 删除餐饮套餐
// @Tags 餐饮套餐
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /cateringSet/deleteCateringSet [post]
func (c *cateringset) deleteCateringSet(ctx ack.Context) {
	ack.Post(ctx, service.CateringSet.DeleteCateringSet)
}
