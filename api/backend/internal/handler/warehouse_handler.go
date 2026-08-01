package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type warehouse struct{}

func init() {
	ack.RegisterRoute(&warehouse{})
}

func (w *warehouse) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/warehouse")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageWarehouseByKey", w.pageWarehouseByKey)
	_g.Get("/listWarehouseStaff", w.listWarehouseStaff)
	_g.Post("/addWarehouse", w.addWarehouse)
	_g.Get("/getWarehouseById", w.getWarehouseById)
	_g.Put("/editWarehouse", w.editWarehouse)
	_g.Delete("/deleteWarehouse", w.deleteWarehouse)
}

func (w *warehouse) pageWarehouseByKey(ctx ack.Context) {
	ack.Get(ctx, service.Warehouse.PageWarehouseByKey)
}

func (w *warehouse) listWarehouseStaff(ctx ack.Context) {
	ack.Get(ctx, service.Warehouse.ListWarehouseStaff)
}

func (w *warehouse) addWarehouse(ctx ack.Context) {
	ack.Post(ctx, service.Warehouse.AddWarehouse)
}

func (w *warehouse) getWarehouseById(ctx ack.Context) {
	ack.Get(ctx, service.Warehouse.GetWarehouseById)
}

func (w *warehouse) editWarehouse(ctx ack.Context) {
	ack.Put(ctx, service.Warehouse.EditWarehouse)
}

func (w *warehouse) deleteWarehouse(ctx ack.Context) {
	ack.Delete(ctx, service.Warehouse.DeleteWarehouse)
}
