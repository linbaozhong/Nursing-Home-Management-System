package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type bedpanorama struct{}

func init() {
	ack.RegisterRoute(&bedpanorama{})
}

func (b *bedpanorama) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/bedPanorama")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/listBuilding", b.listBuilding)
	_g.Get("/listFloorByBuildingId", b.listFloorByBuildingId)
	_g.Get("/listRoomByKey", b.listRoomByKey)
}

// 获取楼栋列表
// @Summary 获取楼栋列表
// @Tags 床位全景
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} []dto.DropDown
// @Router /bedPanorama/listBuilding [get]
func (b *bedpanorama) listBuilding(ctx ack.Context) {
	ack.Get(ctx, service.BedPanorama.ListBuilding)
}

// 获取楼层列表
// @Summary 获取楼层列表
// @Tags 床位全景
// @Accept application/json
// @Produce application/json
// @Param data query dto.ListFloorByBuildingIdReq true "ListFloorByBuildingIdReq"
// @Success 200 {object} []dto.DropDown
// @Router /bedPanorama/listFloorByBuildingId [get]
func (b *bedpanorama) listFloorByBuildingId(ctx ack.Context) {
	ack.Get(ctx, service.BedPanorama.ListFloorByBuildingId)
}

// 获取房间列表（含床位与入住老人）
// @Summary 获取房间列表
// @Tags 床位全景
// @Accept application/json
// @Produce application/json
// @Param data query dto.ListRoomByKeyReq true "ListRoomByKeyReq"
// @Success 200 {object} []dto.FloorItemResp
// @Router /bedPanorama/listRoomByKey [get]
func (b *bedpanorama) listRoomByKey(ctx ack.Context) {
	ack.Get(ctx, service.BedPanorama.ListRoomByKey)
}
