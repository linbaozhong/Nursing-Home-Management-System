package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type warehouserecord struct{}

func init() {
	ack.RegisterRoute(&warehouserecord{})
}

func (w *warehouserecord) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/warehouseRecord")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageWarehouseRecordByKey", w.pageWarehouseRecordByKey)
	_g.Get("/listWarehouse", w.listWarehouse)
	_g.Get("/listWarehouseStaff", w.listWarehouseStaff)
	_g.Get("/pageMaterialByKey", w.pageMaterialByKey)
	_g.Post("/addWarehouseRecord", w.addWarehouseRecord)
	_g.Get("/getWarehouseRecordById", w.getWarehouseRecordById)
	_g.Put("/auditWarehouseRecord", w.auditWarehouseRecord)
	_g.Delete("/deleteWarehouseRecord", w.deleteWarehouseRecord)
}

func (w *warehouserecord) pageWarehouseRecordByKey(ctx ack.Context) {
	ack.Get(ctx, service.WarehouseRecord.PageWarehouseRecordByKey)
}

func (w *warehouserecord) listWarehouse(ctx ack.Context) {
	ack.Get(ctx, service.WarehouseRecord.ListWarehouse)
}

func (w *warehouserecord) listWarehouseStaff(ctx ack.Context) {
	ack.Get(ctx, service.Warehouse.ListWarehouseStaff)
}

func (w *warehouserecord) pageMaterialByKey(ctx ack.Context) {
	ack.Get(ctx, service.Material.PageMaterialByKey)
}

func (w *warehouserecord) addWarehouseRecord(ctx ack.Context) {
	ack.Post(ctx, service.WarehouseRecord.AddWarehouseRecord)
}

func (w *warehouserecord) getWarehouseRecordById(ctx ack.Context) {
	ack.Get(ctx, service.WarehouseRecord.GetWarehouseRecordById)
}

func (w *warehouserecord) auditWarehouseRecord(ctx ack.Context) {
	ack.Put(ctx, service.WarehouseRecord.AuditWarehouseRecord)
}

func (w *warehouserecord) deleteWarehouseRecord(ctx ack.Context) {
	ack.Delete(ctx, service.WarehouseRecord.DeleteWarehouseRecord)
}
