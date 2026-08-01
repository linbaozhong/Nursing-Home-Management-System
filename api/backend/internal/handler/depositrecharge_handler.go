package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type depositrecharge struct{}

func init() {
	ack.RegisterRoute(&depositrecharge{})
}

func (d *depositrecharge) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/depositRecharge")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageDepositRechargeByKey", d.pageDepositRechargeByKey)
	_g.Get("/getDepositRechargeById", d.getDepositRechargeById)
	_g.Post("/addDepositRecharge", d.addDepositRecharge)
	_g.Put("/editDepositRecharge", d.editDepositRecharge)
	_g.Delete("/deleteDepositRecharge", d.deleteDepositRecharge)
	_g.Get("/pageSearchElderByKey", d.pageSearchElderByKey)
	_g.Get("/pageSearchStaffByKey", d.pageSearchStaffByKey)
	_g.Get("/getElderFeeById", d.getElderFeeById)
	_g.Put("/auditElderFee", d.auditElderFee)
}

func (d *depositrecharge) pageDepositRechargeByKey(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.PageDepositRechargeByKey)
}

func (d *depositrecharge) getDepositRechargeById(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.GetDepositRechargeById)
}

func (d *depositrecharge) addDepositRecharge(ctx ack.Context) {
	ack.Post(ctx, service.DepositRecharge.AddDepositRecharge)
}

func (d *depositrecharge) editDepositRecharge(ctx ack.Context) {
	ack.Put(ctx, service.DepositRecharge.EditDepositRecharge)
}

func (d *depositrecharge) deleteDepositRecharge(ctx ack.Context) {
	ack.Delete(ctx, service.DepositRecharge.DeleteDepositRecharge)
}

func (d *depositrecharge) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.PageSearchElderByKey)
}

func (d *depositrecharge) pageSearchStaffByKey(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.PageSearchStaffByKey)
}

func (d *depositrecharge) getElderFeeById(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.GetElderFeeById)
}

func (d *depositrecharge) auditElderFee(ctx ack.Context) {
	ack.Put(ctx, service.DepositRecharge.AuditElderFee)
}
