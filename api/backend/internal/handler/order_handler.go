package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type order struct{}

func init() {
	ack.RegisterRoute(&order{})
}

func (o *order) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/order")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageOrderByKey", o.pageOrderByKey)
	_g.Post("/addOrder", o.addOrder)
	_g.Get("/getOrderById", o.getOrderById)
	_g.Put("/sendOrder", o.sendOrder)
}

func (o *order) pageOrderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Order.PageOrderByKey)
}

func (o *order) addOrder(ctx ack.Context) {
	ack.Post(ctx, service.Order.AddOrder)
}

func (o *order) getOrderById(ctx ack.Context) {
	ack.Get(ctx, service.Order.GetOrderById)
}

func (o *order) sendOrder(ctx ack.Context) {
	ack.Put(ctx, service.Order.SendOrder)
}
