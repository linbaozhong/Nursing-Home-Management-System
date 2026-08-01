package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type roomtype struct{}

func init() {
	ack.RegisterRoute(&roomtype{})
}

func (r *roomtype) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/roomType")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageRoomTypeByKey", r.pageRoomTypeByKey)
	_g.Get("/getRoomTypeById", r.getRoomTypeById)
	_g.Post("/addRoomType", r.addRoomType)
	_g.Put("/editRoomType", r.editRoomType)
	_g.Delete("/deleteRoomType", r.deleteRoomType)
}

func (r *roomtype) pageRoomTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.RoomType.PageRoomTypeByKey)
}

func (r *roomtype) getRoomTypeById(ctx ack.Context) {
	ack.Get(ctx, service.RoomType.GetRoomTypeById)
}

func (r *roomtype) addRoomType(ctx ack.Context) {
	ack.Post(ctx, service.RoomType.AddRoomType)
}

func (r *roomtype) editRoomType(ctx ack.Context) {
	ack.Put(ctx, service.RoomType.EditRoomType)
}

func (r *roomtype) deleteRoomType(ctx ack.Context) {
	ack.Delete(ctx, service.RoomType.DeleteRoomType)
}
