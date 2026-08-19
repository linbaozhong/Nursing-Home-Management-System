package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/editRoomType", r.editRoomType)
	_g.Post("/deleteRoomType", r.deleteRoomType)
}

// 分页查询房型分类
// @Summary 分页查询房型分类
// @Tags 房型
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageRoomTypeByKeyQuery true "PageRoomTypeByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /roomType/pageRoomTypeByKey [get]
func (r *roomtype) pageRoomTypeByKey(ctx ack.Context) {
	ack.Get(ctx, service.RoomType.PageRoomTypeByKey)
}

// 获取房型分类
// @Summary 获取房型分类
// @Tags 房型
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /roomType/getRoomTypeById [get]
func (r *roomtype) getRoomTypeById(ctx ack.Context) {
	ack.Get(ctx, service.RoomType.GetRoomTypeById)
}

// 新增房型
// @Summary 新增房型
// @Tags 房型
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateRoomTypeQuery true "OperateRoomTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /roomType/addRoomType [post]
func (r *roomtype) addRoomType(ctx ack.Context) {
	ack.Post(ctx, service.RoomType.AddRoomType)
}

// 编辑房型
// @Summary 编辑房型
// @Tags 房型
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateRoomTypeQuery true "OperateRoomTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /roomType/editRoomType [post]
func (r *roomtype) editRoomType(ctx ack.Context) {
	ack.Post(ctx, service.RoomType.EditRoomType)
}

// 删除房型
// @Summary 删除房型
// @Tags 房型
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /roomType/deleteRoomType [post]
func (r *roomtype) deleteRoomType(ctx ack.Context) {
	ack.Post(ctx, service.RoomType.DeleteRoomType)
}
