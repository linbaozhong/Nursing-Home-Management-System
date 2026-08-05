package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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

// Overview
// @Summary Overview
// @Tags 首页
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.TodayOverviewVO
// @Router /home/todayOverview [get]
func (h *home) todayOverview(ctx ack.Context) {
	ack.Get(ctx, service.Home.TodayOverview)
}

// 床位
// @Summary 床位
// @Tags 首页
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.AvailableBedVO
// @Router /home/availableBed [get]
func (h *home) availableBed(ctx ack.Context) {
	ack.Get(ctx, service.Home.AvailableBed)
}

// SaleFollow
// @Summary SaleFollow
// @Tags 首页
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.TodaySaleFollowVO
// @Router /home/todaySaleFollow [get]
func (h *home) todaySaleFollow(ctx ack.Context) {
	ack.Get(ctx, service.Home.TodaySaleFollow)
}

// PerformanceRank
// @Summary PerformanceRank
// @Tags 首页
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} dto.MonthPerformanceRankVO
// @Router /home/monthPerformanceRank [get]
func (h *home) monthPerformanceRank(ctx ack.Context) {
	ack.Get(ctx, service.Home.MonthPerformanceRank)
}

// 来源
// @Summary 来源
// @Tags 首页
// @Accept application/json
// @Produce application/json
// @Param data query dto.ClientSourceQuery true "ClientSourceQuery"
// @Success 200 {object} []dto.ClientSourceVO
// @Router /home/clientSource [get]
func (h *home) clientSource(ctx ack.Context) {
	ack.Get(ctx, service.Home.ClientSource)
}

// Trend
// @Summary Trend
// @Tags 首页
// @Accept application/json
// @Produce application/json
// @Param data query dto.EmptyReq true "EmptyReq"
// @Success 200 {object} []dto.BusinessTrendVO
// @Router /home/businessTrend [get]
func (h *home) businessTrend(ctx ack.Context) {
	ack.Get(ctx, service.Home.BusinessTrend)
}
