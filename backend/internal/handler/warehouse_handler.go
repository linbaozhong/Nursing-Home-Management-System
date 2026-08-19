package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/editWarehouse", w.editWarehouse)
	_g.Post("/deleteWarehouse", w.deleteWarehouse)
}

// 分页查询仓库
// @Summary 分页查询仓库
// @Tags 仓库
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageWarehouseByKeyQuery true "PageWarehouseByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouse/pageWarehouseByKey [get]
func (w *warehouse) pageWarehouseByKey(ctx ack.Context) {
	ack.Get(ctx, service.Warehouse.PageWarehouseByKey)
}

// 查询仓库员工
// @Summary 查询仓库员工
// @Tags 仓库
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouse/listWarehouseStaff [get]
func (w *warehouse) listWarehouseStaff(ctx ack.Context) {
	ack.Get(ctx, service.Warehouse.ListWarehouseStaff)
}

// 新增仓库
// @Summary 新增仓库
// @Tags 仓库
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateWarehouseQuery true "OperateWarehouseQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouse/addWarehouse [post]
func (w *warehouse) addWarehouse(ctx ack.Context) {
	ack.Post(ctx, service.Warehouse.AddWarehouse)
}

// 获取仓库
// @Summary 获取仓库
// @Tags 仓库
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouse/getWarehouseById [get]
func (w *warehouse) getWarehouseById(ctx ack.Context) {
	ack.Get(ctx, service.Warehouse.GetWarehouseById)
}

// 编辑仓库
// @Summary 编辑仓库
// @Tags 仓库
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateWarehouseQuery true "OperateWarehouseQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouse/editWarehouse [post]
func (w *warehouse) editWarehouse(ctx ack.Context) {
	ack.Post(ctx, service.Warehouse.EditWarehouse)
}

// 删除仓库
// @Summary 删除仓库
// @Tags 仓库
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /warehouse/deleteWarehouse [post]
func (w *warehouse) deleteWarehouse(ctx ack.Context) {
	ack.Post(ctx, service.Warehouse.DeleteWarehouse)
}
