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

func (b *bedpanorama) getBedPanorama(ctx ack.Context) {
	ack.Get(ctx, service.BedPanorama.GetBedPanorama)
}
