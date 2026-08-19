package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type dishes struct{}

func init() {
	ack.RegisterRoute(&dishes{})
}

func (d *dishes) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/dishes")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageDishesByKey", d.pageDishesByKey)
	_g.Get("/getDishesById", d.getDishesById)
	_g.Post("/addDishes", d.addDishes)
	_g.Post("/editDishes", d.editDishes)
	_g.Post("/deleteDishes", d.deleteDishes)
	_g.Get("/pageDishesTypeByKey", d.pageDishesTypeByKey)
	_g.Get("/getDishesTypeById", d.getDishesTypeById)
	_g.Post("/addDishesType", d.addDishesType)
	_g.Post("/editDishesType", d.editDishesType)
	_g.Post("/deleteDishesType", d.deleteDishesType)
}

// 分页查询Dishes
// @Summary 分页查询Dishes
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageDishesByKeyQuery true "PageDishesByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/pageDishesByKey [get]
func (d *dishes) pageDishesByKey(ctx ack.Context) {
	ack.Get(ctx, service.Dishes.PageDishesByKey)
}

// 获取Dishes
// @Summary 获取Dishes
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/getDishesById [get]
func (d *dishes) getDishesById(ctx ack.Context) {
	ack.Get(ctx, service.Dishes.GetDishesById)
}

// 新增Dishes
// @Summary 新增Dishes
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddDishesQuery true "AddDishesQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/addDishes [post]
func (d *dishes) addDishes(ctx ack.Context) {
	ack.Post(ctx, service.Dishes.AddDishes)
}

// 编辑Dishes
// @Summary 编辑Dishes
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditDishesQuery true "EditDishesQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/editDishes [post]
func (d *dishes) editDishes(ctx ack.Context) {
	ack.Post(ctx, service.Dishes.EditDishes)
}

// 删除Dishes
// @Summary 删除Dishes
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/deleteDishes [post]
func (d *dishes) deleteDishes(ctx ack.Context) {
	ack.Post(ctx, service.Dishes.DeleteDishes)
}

// 分页查询Dishes分类
// @Summary 分页查询Dishes分类
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageDishesTypeByKeyQuery true "PageDishesTypeByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/pageDishesTypeByKey [get]
func (d *dishes) pageDishesTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.Dishes.PageDishesTypeByKey)
}

// 获取Dishes分类
// @Summary 获取Dishes分类
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/getDishesTypeById [get]
func (d *dishes) getDishesTypeById(ctx ack.Context) {
	ack.Get(ctx, service.Dishes.GetDishesTypeById)
}

// 新增Dishes分类
// @Summary 新增Dishes分类
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddDishesTypeQuery true "AddDishesTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/addDishesType [post]
func (d *dishes) addDishesType(ctx ack.Context) {
	ack.Post(ctx, service.Dishes.AddDishesType)
}

// 编辑Dishes分类
// @Summary 编辑Dishes分类
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditDishesTypeQuery true "EditDishesTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/editDishesType [post]
func (d *dishes) editDishesType(ctx ack.Context) {
	ack.Post(ctx, service.Dishes.EditDishesType)
}

// 删除Dishes分类
// @Summary 删除Dishes分类
// @Tags 菜品
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /dishes/deleteDishesType [post]
func (d *dishes) deleteDishesType(ctx ack.Context) {
	ack.Post(ctx, service.Dishes.DeleteDishesType)
}
