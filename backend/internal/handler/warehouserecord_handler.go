package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/auditWarehouseRecord", w.auditWarehouseRecord)
	_g.Post("/deleteWarehouseRecord", w.deleteWarehouseRecord)
}

// 分页查询仓库记录
// @Summary 分页查询仓库记录
// @Tags 入库记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageWarehouseRecordByKeyReq true "PageWarehouseRecordByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouseRecord/pageWarehouseRecordByKey [get]
func (w *warehouserecord) pageWarehouseRecordByKey(ctx ack.Context) {
	ack.Get(ctx, service.WarehouseRecord.PageWarehouseRecordByKey)
}

// 查询仓库
// @Summary 查询仓库
// @Tags 入库记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouseRecord/listWarehouse [get]
func (w *warehouserecord) listWarehouse(ctx ack.Context) {
	ack.Get(ctx, service.WarehouseRecord.ListWarehouse)
}

// 查询仓库员工
// @Summary 查询仓库员工
// @Tags 入库记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouseRecord/listWarehouseStaff [get]
func (w *warehouserecord) listWarehouseStaff(ctx ack.Context) {
	ack.Get(ctx, service.Warehouse.ListWarehouseStaff)
}

// 分页查询物料
// @Summary 分页查询物料
// @Tags 入库记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageMaterialByKeyReq true "PageMaterialByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouseRecord/pageMaterialByKey [get]
func (w *warehouserecord) pageMaterialByKey(ctx ack.Context) {
	ack.Get(ctx, service.Material.PageMaterialByKey)
}

// 新增入库记录
// @Summary 新增入库记录
// @Tags 入库记录
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddWarehouseRecordReq true "AddWarehouseRecordReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouseRecord/addWarehouseRecord [post]
func (w *warehouserecord) addWarehouseRecord(ctx ack.Context) {
	ack.Post(ctx, service.WarehouseRecord.AddWarehouseRecord)
}

// 获取仓库记录
// @Summary 获取仓库记录
// @Tags 入库记录
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouseRecord/getWarehouseRecordById [get]
func (w *warehouserecord) getWarehouseRecordById(ctx ack.Context) {
	ack.Get(ctx, service.WarehouseRecord.GetWarehouseRecordById)
}

// 审核入库记录
// @Summary 审核入库记录
// @Tags 入库记录
// @Accept application/json
// @Produce application/json
// @Param data body dto.AuditWarehouseRecordReq true "AuditWarehouseRecordReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouseRecord/auditWarehouseRecord [post]
func (w *warehouserecord) auditWarehouseRecord(ctx ack.Context) {
	ack.Post(ctx, service.WarehouseRecord.AuditWarehouseRecord)
}

// 删除入库记录
// @Summary 删除入库记录
// @Tags 入库记录
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouseRecord/deleteWarehouseRecord [post]
func (w *warehouserecord) deleteWarehouseRecord(ctx ack.Context) {
	ack.Post(ctx, service.WarehouseRecord.DeleteWarehouseRecord)
}
