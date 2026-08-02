package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type bedpanorama struct{}

func init() {
	ack.RegisterRoute(&bedpanorama{})
}

func (b *bedpanorama) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/bedPanorama")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/getBedPanorama", b.getBedPanorama)
}

// 获取床位全景
// @Summary 获取床位全景
// @Tags 床位全景
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /bedPanorama/getBedPanorama [get]
func (b *bedpanorama) getBedPanorama(ctx ack.Context) {
	ack.Get(ctx, service.BedPanorama.GetBedPanorama)
}
