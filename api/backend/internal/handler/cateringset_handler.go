package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
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
	_g.Put("/editCateringSet", c.editCateringSet)
	_g.Delete("/deleteCateringSet", c.deleteCateringSet)
}

func (c *cateringset) pageCateringSetByKey(ctx ack.Context) {
	ack.Get(ctx, service.CateringSet.PageCateringSetByKey)
}

func (c *cateringset) getCateringSetById(ctx ack.Context) {
	ack.Get(ctx, service.CateringSet.GetCateringSetById)
}

func (c *cateringset) addCateringSet(ctx ack.Context) {
	ack.Post(ctx, service.CateringSet.AddCateringSet)
}

func (c *cateringset) editCateringSet(ctx ack.Context) {
	ack.Put(ctx, service.CateringSet.EditCateringSet)
}

func (c *cateringset) deleteCateringSet(ctx ack.Context) {
	ack.Delete(ctx, service.CateringSet.DeleteCateringSet)
}
