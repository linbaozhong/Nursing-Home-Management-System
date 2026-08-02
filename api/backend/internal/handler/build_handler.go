package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/editBuilding", b.editBuilding)
	_g.Post("/deleteBuilding", b.deleteBuilding)
	_g.Get("/pageFloorByKey", b.pageFloorByKey)
	_g.Get("/getFloorById", b.getFloorById)
	_g.Post("/addFloor", b.addFloor)
	_g.Post("/editFloor", b.editFloor)
	_g.Post("/deleteFloor", b.deleteFloor)
	_g.Get("/pageRoomByKey", b.pageRoomByKey)
	_g.Get("/getRoomById", b.getRoomById)
	_g.Post("/addRoom", b.addRoom)
	_g.Post("/editRoom", b.editRoom)
	_g.Post("/deleteRoom", b.deleteRoom)
	_g.Get("/listRoomType", b.listRoomType)
}

// 查询房型
// @Summary 查询房型
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /build/listRoomType [get]
func (b *build) listRoomType(ctx ack.Context) {
	ack.Get(ctx, service.Build.ListRoomType)
}

// 分页查询楼栋
// @Summary 分页查询楼栋
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageBuildingByKeyQuery true "PageBuildingByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /build/pageBuildingByKey [get]
func (b *build) pageBuildingByKey(ctx ack.Context) {
	ack.Get(ctx, service.Build.PageBuildingByKey)
}

// 获取楼栋
// @Summary 获取楼栋
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /build/getBuildingById [get]
func (b *build) getBuildingById(ctx ack.Context) {
	ack.Get(ctx, service.Build.GetBuildingById)
}

// 新增楼栋
// @Summary 新增楼栋
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddBuildingQuery true "AddBuildingQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /build/addBuilding [post]
func (b *build) addBuilding(ctx ack.Context) {
	ack.Post(ctx, service.Build.AddBuilding)
}

// 编辑楼栋
// @Summary 编辑楼栋
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditBuildingQuery true "EditBuildingQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /build/editBuilding [post]
func (b *build) editBuilding(ctx ack.Context) {
	ack.Post(ctx, service.Build.EditBuilding)
}

// 删除楼栋
// @Summary 删除楼栋
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /build/deleteBuilding [post]
func (b *build) deleteBuilding(ctx ack.Context) {
	ack.Post(ctx, service.Build.DeleteBuilding)
}

// 分页查询楼层
// @Summary 分页查询楼层
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageFloorByKeyQuery true "PageFloorByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /build/pageFloorByKey [get]
func (b *build) pageFloorByKey(ctx ack.Context) {
	ack.Get(ctx, service.Build.PageFloorByKey)
}

// 获取楼层
// @Summary 获取楼层
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /build/getFloorById [get]
func (b *build) getFloorById(ctx ack.Context) {
	ack.Get(ctx, service.Build.GetFloorById)
}

// 新增楼层
// @Summary 新增楼层
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddFloorQuery true "AddFloorQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /build/addFloor [post]
func (b *build) addFloor(ctx ack.Context) {
	ack.Post(ctx, service.Build.AddFloor)
}

// 编辑楼层
// @Summary 编辑楼层
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditFloorQuery true "EditFloorQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /build/editFloor [post]
func (b *build) editFloor(ctx ack.Context) {
	ack.Post(ctx, service.Build.EditFloor)
}

// 删除楼层
// @Summary 删除楼层
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /build/deleteFloor [post]
func (b *build) deleteFloor(ctx ack.Context) {
	ack.Post(ctx, service.Build.DeleteFloor)
}

// 分页查询房型
// @Summary 分页查询房型
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageRoomByKeyQuery true "PageRoomByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /build/pageRoomByKey [get]
func (b *build) pageRoomByKey(ctx ack.Context) {
	ack.Get(ctx, service.Build.PageRoomByKey)
}

// 获取房型
// @Summary 获取房型
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /build/getRoomById [get]
func (b *build) getRoomById(ctx ack.Context) {
	ack.Get(ctx, service.Build.GetRoomById)
}

// 新增房型
// @Summary 新增房型
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddRoomQuery true "AddRoomQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /build/addRoom [post]
func (b *build) addRoom(ctx ack.Context) {
	ack.Post(ctx, service.Build.AddRoom)
}

// 编辑房型
// @Summary 编辑房型
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditRoomQuery true "EditRoomQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /build/editRoom [post]
func (b *build) editRoom(ctx ack.Context) {
	ack.Post(ctx, service.Build.EditRoom)
}

// 删除房型
// @Summary 删除房型
// @Tags 楼栋
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /build/deleteRoom [post]
func (b *build) deleteRoom(ctx ack.Context) {
	ack.Post(ctx, service.Build.DeleteRoom)
}
