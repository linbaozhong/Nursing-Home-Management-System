package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type accountledger struct{}

func init() {
	ack.RegisterRoute(&accountledger{})
}

// RegisterRoute 老人资金明细台账路由
func (a *accountledger) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/account/ledger")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageAccountLedgerByKey", a.pageAccountLedgerByKey)
	_g.Get("/getBalanceSummary", a.getBalanceSummary)
	_g.Post("/changeBalance", a.changeBalance)
}

// 分页查询老人资金明细台账
// @Summary 分页查询老人资金明细台账
// @Tags 资金台账
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageAccountLedgerByKeyReq true "PageAccountLedgerByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /account/ledger/pageAccountLedgerByKey [get]
func (a *accountledger) pageAccountLedgerByKey(ctx ack.Context) {
	ack.Get(ctx, service.AccountLedger.PageAccountLedgerByKey)
}

// 老人资金账户汇总
// @Summary 老人资金账户汇总
// @Tags 资金台账
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.PageAccountLedgerBalanceResp
// @Router /account/ledger/getBalanceSummary [get]
func (a *accountledger) getBalanceSummary(ctx ack.Context) {
	ack.Get(ctx, service.AccountLedger.GetBalanceSummary)
}

// 老人余额增减记账（事务+幂等）
// @Summary 老人余额增减记账
// @Tags 资金台账
// @Accept application/json
// @Produce application/json
// @Param data body dto.ChangeBalanceReq true "ChangeBalanceReq"
// @Success 200 {object} dto.EmptyResp
// @Router /account/ledger/changeBalance [post]
func (a *accountledger) changeBalance(ctx ack.Context) {
	ack.Post(ctx, service.AccountLedger.ChangeBalance)
}
