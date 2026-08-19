package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/sendOrder", o.sendOrder)
}

// 分页查询订单
// @Summary 分页查询订单
// @Tags 订单
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageOrderByKeyQuery true "PageOrderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /order/pageOrderByKey [get]
func (o *order) pageOrderByKey(ctx ack.Context) {
	ack.Get(ctx, service.Order.PageOrderByKey)
}

// 新增订单
// @Summary 新增订单
// @Tags 订单
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddOrderQuery true "AddOrderQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /order/addOrder [post]
func (o *order) addOrder(ctx ack.Context) {
	ack.Post(ctx, service.Order.AddOrder)
}

// 获取订单
// @Summary 获取订单
// @Tags 订单
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /order/getOrderById [get]
func (o *order) getOrderById(ctx ack.Context) {
	ack.Get(ctx, service.Order.GetOrderById)
}

// 发送订单
// @Summary 发送订单
// @Tags 订单
// @Accept application/json
// @Produce application/json
// @Param data body dto.SendOrderQuery true "SendOrderQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /order/sendOrder [post]
func (o *order) sendOrder(ctx ack.Context) {
	ack.Post(ctx, service.Order.SendOrder)
}
