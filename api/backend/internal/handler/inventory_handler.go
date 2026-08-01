package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type inventory struct{}

func init() {
	ack.RegisterRoute(&inventory{})
}

func (i *inventory) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/inventory")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageInventoryByKey", i.pageInventoryByKey)
	_g.Get("/getInventoryById", i.getInventoryById)
	_g.Get("/pageInventoryRecordByKey", i.pageInventoryRecordByKey)
	_g.Get("/getInventoryRecordById", i.getInventoryRecordById)
	_g.Post("/addInventory", i.addInventory)
	_g.Put("/editInventory", i.editInventory)
	_g.Put("/auditInventory", i.auditInventory)
}

func (i *inventory) pageInventoryByKey(ctx ack.Context) {
	ack.Get(ctx, service.Inventory.PageInventoryByKey)
}

func (i *inventory) getInventoryById(ctx ack.Context) {
	ack.Get(ctx, service.Inventory.GetInventoryById)
}

func (i *inventory) pageInventoryRecordByKey(ctx ack.Context) {
	ack.Get(ctx, service.Inventory.PageInventoryRecordByKey)
}

func (i *inventory) getInventoryRecordById(ctx ack.Context) {
	ack.Get(ctx, service.Inventory.GetInventoryRecordById)
}

func (i *inventory) addInventory(ctx ack.Context) {
	ack.Post(ctx, service.Inventory.AddInventory)
}

func (i *inventory) editInventory(ctx ack.Context) {
	ack.Put(ctx, service.Inventory.EditInventory)
}

func (i *inventory) auditInventory(ctx ack.Context) {
	ack.Put(ctx, service.Inventory.AuditInventory)
}
