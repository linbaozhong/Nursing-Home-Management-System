package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type build struct{}

func init() {
	ack.RegisterRoute(&build{})
}

func (b *build) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/build")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageBuildingByKey", b.pageBuildingByKey)
	_g.Get("/getBuildingById", b.getBuildingById)
	_g.Post("/addBuilding", b.addBuilding)
	_g.Put("/editBuilding", b.editBuilding)
	_g.Delete("/deleteBuilding", b.deleteBuilding)
	_g.Get("/pageFloorByKey", b.pageFloorByKey)
	_g.Get("/getFloorById", b.getFloorById)
	_g.Post("/addFloor", b.addFloor)
	_g.Put("/editFloor", b.editFloor)
	_g.Delete("/deleteFloor", b.deleteFloor)
	_g.Get("/pageRoomByKey", b.pageRoomByKey)
	_g.Get("/getRoomById", b.getRoomById)
	_g.Post("/addRoom", b.addRoom)
	_g.Put("/editRoom", b.editRoom)
	_g.Delete("/deleteRoom", b.deleteRoom)
	_g.Get("/listRoomType", b.listRoomType)
}

func (b *build) listRoomType(ctx ack.Context) {
	ack.Get(ctx, service.Build.ListRoomType)
}

func (b *build) pageBuildingByKey(ctx ack.Context) {
	ack.Get(ctx, service.Build.PageBuildingByKey)
}

func (b *build) getBuildingById(ctx ack.Context) {
	ack.Get(ctx, service.Build.GetBuildingById)
}

func (b *build) addBuilding(ctx ack.Context) {
	ack.Post(ctx, service.Build.AddBuilding)
}

func (b *build) editBuilding(ctx ack.Context) {
	ack.Put(ctx, service.Build.EditBuilding)
}

func (b *build) deleteBuilding(ctx ack.Context) {
	ack.Delete(ctx, service.Build.DeleteBuilding)
}

func (b *build) pageFloorByKey(ctx ack.Context) {
	ack.Get(ctx, service.Build.PageFloorByKey)
}

func (b *build) getFloorById(ctx ack.Context) {
	ack.Get(ctx, service.Build.GetFloorById)
}

func (b *build) addFloor(ctx ack.Context) {
	ack.Post(ctx, service.Build.AddFloor)
}

func (b *build) editFloor(ctx ack.Context) {
	ack.Put(ctx, service.Build.EditFloor)
}

func (b *build) deleteFloor(ctx ack.Context) {
	ack.Delete(ctx, service.Build.DeleteFloor)
}

func (b *build) pageRoomByKey(ctx ack.Context) {
	ack.Get(ctx, service.Build.PageRoomByKey)
}

func (b *build) getRoomById(ctx ack.Context) {
	ack.Get(ctx, service.Build.GetRoomById)
}

func (b *build) addRoom(ctx ack.Context) {
	ack.Post(ctx, service.Build.AddRoom)
}

func (b *build) editRoom(ctx ack.Context) {
	ack.Put(ctx, service.Build.EditRoom)
}

func (b *build) deleteRoom(ctx ack.Context) {
	ack.Delete(ctx, service.Build.DeleteRoom)
}
