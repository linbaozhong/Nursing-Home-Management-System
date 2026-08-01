package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
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
	_g.Put("/editDishes", d.editDishes)
	_g.Delete("/deleteDishes", d.deleteDishes)
	_g.Get("/pageDishesTypeByKey", d.pageDishesTypeByKey)
	_g.Get("/getDishesTypeById", d.getDishesTypeById)
	_g.Post("/addDishesType", d.addDishesType)
	_g.Put("/editDishesType", d.editDishesType)
	_g.Delete("/deleteDishesType", d.deleteDishesType)
}

func (d *dishes) pageDishesByKey(ctx ack.Context) {
	ack.Get(ctx, service.Dishes.PageDishesByKey)
}

func (d *dishes) getDishesById(ctx ack.Context) {
	ack.Get(ctx, service.Dishes.GetDishesById)
}

func (d *dishes) addDishes(ctx ack.Context) {
	ack.Post(ctx, service.Dishes.AddDishes)
}

func (d *dishes) editDishes(ctx ack.Context) {
	ack.Put(ctx, service.Dishes.EditDishes)
}

func (d *dishes) deleteDishes(ctx ack.Context) {
	ack.Delete(ctx, service.Dishes.DeleteDishes)
}

func (d *dishes) pageDishesTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.Dishes.PageDishesTypeByKey)
}

func (d *dishes) getDishesTypeById(ctx ack.Context) {
	ack.Get(ctx, service.Dishes.GetDishesTypeById)
}

func (d *dishes) addDishesType(ctx ack.Context) {
	ack.Post(ctx, service.Dishes.AddDishesType)
}

func (d *dishes) editDishesType(ctx ack.Context) {
	ack.Put(ctx, service.Dishes.EditDishesType)
}

func (d *dishes) deleteDishesType(ctx ack.Context) {
	ack.Delete(ctx, service.Dishes.DeleteDishesType)
}
