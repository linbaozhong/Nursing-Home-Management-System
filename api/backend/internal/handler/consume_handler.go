package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type consume struct{}

func init() {
	ack.RegisterRoute(&consume{})
}

func (c *consume) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/consume")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageConsumeByKey", c.pageConsumeByKey)
	_g.Get("/getConsumeById", c.getConsumeById)
	_g.Post("/addConsume", c.addConsume)
	_g.Put("/editConsume", c.editConsume)
	_g.Delete("/deleteConsume", c.deleteConsume)
}

func (c *consume) pageConsumeByKey(ctx ack.Context) {
	ack.Get(ctx, service.Consume.PageConsumeByKey)
}

func (c *consume) getConsumeById(ctx ack.Context) {
	ack.Get(ctx, service.Consume.GetConsumeById)
}

func (c *consume) addConsume(ctx ack.Context) {
	ack.Post(ctx, service.Consume.AddConsume)
}

func (c *consume) editConsume(ctx ack.Context) {
	ack.Put(ctx, service.Consume.EditConsume)
}

func (c *consume) deleteConsume(ctx ack.Context) {
	ack.Delete(ctx, service.Consume.DeleteConsume)
}
