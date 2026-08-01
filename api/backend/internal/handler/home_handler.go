package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type home struct{}

func init() {
	ack.RegisterRoute(&home{})
}

func (h *home) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/home")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/todayOverview", h.todayOverview)
	_g.Get("/availableBed", h.availableBed)
	_g.Get("/todaySaleFollow", h.todaySaleFollow)
	_g.Get("/monthPerformanceRank", h.monthPerformanceRank)
	_g.Get("/clientSource", h.clientSource)
	_g.Get("/businessTrend", h.businessTrend)
}

func (h *home) todayOverview(ctx ack.Context) {
	ack.Get(ctx, service.Home.TodayOverview)
}

func (h *home) availableBed(ctx ack.Context) {
	ack.Get(ctx, service.Home.AvailableBed)
}

func (h *home) todaySaleFollow(ctx ack.Context) {
	ack.Get(ctx, service.Home.TodaySaleFollow)
}

func (h *home) monthPerformanceRank(ctx ack.Context) {
	ack.Get(ctx, service.Home.MonthPerformanceRank)
}

func (h *home) clientSource(ctx ack.Context) {
	ack.Get(ctx, service.Home.ClientSource)
}

func (h *home) businessTrend(ctx ack.Context) {
	ack.Get(ctx, service.Home.BusinessTrend)
}
