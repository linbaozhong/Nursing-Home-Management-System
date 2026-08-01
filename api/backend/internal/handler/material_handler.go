package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
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
	_g.Put("/editMaterial", m.editMaterial)
	_g.Delete("/deleteMaterial", m.deleteMaterial)
	_g.Get("/pageMaterialTypeByKey", m.pageMaterialTypeByKey)
	_g.Get("/getMaterialTypeById", m.getMaterialTypeById)
	_g.Post("/addMaterialType", m.addMaterialType)
	_g.Put("/editMaterialType", m.editMaterialType)
	_g.Delete("/deleteMaterialType", m.deleteMaterialType)
}

func (m *material) pageMaterialByKey(ctx ack.Context) {
	ack.Get(ctx, service.Material.PageMaterialByKey)
}

func (m *material) getMaterialById(ctx ack.Context) {
	ack.Get(ctx, service.Material.GetMaterialById)
}

func (m *material) addMaterial(ctx ack.Context) {
	ack.Post(ctx, service.Material.AddMaterial)
}

func (m *material) editMaterial(ctx ack.Context) {
	ack.Put(ctx, service.Material.EditMaterial)
}

func (m *material) deleteMaterial(ctx ack.Context) {
	ack.Delete(ctx, service.Material.DeleteMaterial)
}

func (m *material) pageMaterialTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.Material.PageMaterialTypeByKey)
}

func (m *material) getMaterialTypeById(ctx ack.Context) {
	ack.Get(ctx, service.Material.GetMaterialTypeById)
}

func (m *material) addMaterialType(ctx ack.Context) {
	ack.Post(ctx, service.Material.AddMaterialType)
}

func (m *material) editMaterialType(ctx ack.Context) {
	ack.Put(ctx, service.Material.EditMaterialType)
}

func (m *material) deleteMaterialType(ctx ack.Context) {
	ack.Delete(ctx, service.Material.DeleteMaterialType)
}
