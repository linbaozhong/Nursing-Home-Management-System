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
	_g.Post("/editInventory", i.editInventory)
	_g.Post("/auditInventory", i.auditInventory)
}

// 分页查询库存
// @Summary 分页查询库存
// @Tags 库存
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageInventoryByKeyQuery true "PageInventoryByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /inventory/pageInventoryByKey [get]
func (i *inventory) pageInventoryByKey(ctx ack.Context) {
	ack.Get(ctx, service.Inventory.PageInventoryByKey)
}

// 获取库存
// @Summary 获取库存
// @Tags 库存
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /inventory/getInventoryById [get]
func (i *inventory) getInventoryById(ctx ack.Context) {
	ack.Get(ctx, service.Inventory.GetInventoryById)
}

// 分页查询库存记录
// @Summary 分页查询库存记录
// @Tags 库存
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageInventoryRecordByKeyQuery true "PageInventoryRecordByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /inventory/pageInventoryRecordByKey [get]
func (i *inventory) pageInventoryRecordByKey(ctx ack.Context) {
	ack.Get(ctx, service.Inventory.PageInventoryRecordByKey)
}

// 获取库存记录
// @Summary 获取库存记录
// @Tags 库存
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /inventory/getInventoryRecordById [get]
func (i *inventory) getInventoryRecordById(ctx ack.Context) {
	ack.Get(ctx, service.Inventory.GetInventoryRecordById)
}

// 新增库存
// @Summary 新增库存
// @Tags 库存
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddInventoryQuery true "AddInventoryQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /inventory/addInventory [post]
func (i *inventory) addInventory(ctx ack.Context) {
	ack.Post(ctx, service.Inventory.AddInventory)
}

// 编辑库存
// @Summary 编辑库存
// @Tags 库存
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditInventoryQuery true "EditInventoryQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /inventory/editInventory [post]
func (i *inventory) editInventory(ctx ack.Context) {
	ack.Post(ctx, service.Inventory.EditInventory)
}

// 审核库存
// @Summary 审核库存
// @Tags 库存
// @Accept application/json
// @Produce application/json
// @Param data body dto.AuditInventoryQuery true "AuditInventoryQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /inventory/auditInventory [post]
func (i *inventory) auditInventory(ctx ack.Context) {
	ack.Post(ctx, service.Inventory.AuditInventory)
}
