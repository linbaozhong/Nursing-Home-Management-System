package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type material struct{}

func init() {
	ack.RegisterRoute(&material{})
}

func (m *material) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/material")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageMaterialByKey", m.pageMaterialByKey)
	_g.Get("/getMaterialById", m.getMaterialById)
	_g.Post("/addMaterial", m.addMaterial)
	_g.Post("/editMaterial", m.editMaterial)
	_g.Post("/deleteMaterial", m.deleteMaterial)
	_g.Get("/pageMaterialTypeByKey", m.pageMaterialTypeByKey)
	_g.Get("/getMaterialTypeById", m.getMaterialTypeById)
	_g.Post("/addMaterialType", m.addMaterialType)
	_g.Post("/editMaterialType", m.editMaterialType)
	_g.Post("/deleteMaterialType", m.deleteMaterialType)
}

// 分页查询物料
// @Summary 分页查询物料
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageMaterialByKeyReq true "PageMaterialByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/pageMaterialByKey [get]
func (m *material) pageMaterialByKey(ctx ack.Context) {
	ack.Get(ctx, service.Material.PageMaterialByKey)
}

// 获取物料
// @Summary 获取物料
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/getMaterialById [get]
func (m *material) getMaterialById(ctx ack.Context) {
	ack.Get(ctx, service.Material.GetMaterialById)
}

// 新增物料
// @Summary 新增物料
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddMaterialReq true "AddMaterialReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/addMaterial [post]
func (m *material) addMaterial(ctx ack.Context) {
	ack.Post(ctx, service.Material.AddMaterial)
}

// 编辑物料
// @Summary 编辑物料
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditMaterialReq true "EditMaterialReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/editMaterial [post]
func (m *material) editMaterial(ctx ack.Context) {
	ack.Post(ctx, service.Material.EditMaterial)
}

// 删除物料
// @Summary 删除物料
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/deleteMaterial [post]
func (m *material) deleteMaterial(ctx ack.Context) {
	ack.Post(ctx, service.Material.DeleteMaterial)
}

// 分页查询物料分类
// @Summary 分页查询物料分类
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageMaterialTypeByKeyReq true "PageMaterialTypeByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/pageMaterialTypeByKey [get]
func (m *material) pageMaterialTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.Material.PageMaterialTypeByKey)
}

// 获取物料分类
// @Summary 获取物料分类
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/getMaterialTypeById [get]
func (m *material) getMaterialTypeById(ctx ack.Context) {
	ack.Get(ctx, service.Material.GetMaterialTypeById)
}

// 新增物料分类
// @Summary 新增物料分类
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddMaterialTypeReq true "AddMaterialTypeReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/addMaterialType [post]
func (m *material) addMaterialType(ctx ack.Context) {
	ack.Post(ctx, service.Material.AddMaterialType)
}

// 编辑物料分类
// @Summary 编辑物料分类
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditMaterialTypeReq true "EditMaterialTypeReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/editMaterialType [post]
func (m *material) editMaterialType(ctx ack.Context) {
	ack.Post(ctx, service.Material.EditMaterialType)
}

// 删除物料分类
// @Summary 删除物料分类
// @Tags 物料
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /material/deleteMaterialType [post]
func (m *material) deleteMaterialType(ctx ack.Context) {
	ack.Post(ctx, service.Material.DeleteMaterialType)
}
