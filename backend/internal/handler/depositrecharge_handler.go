package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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
	_g.Post("/editDepositRecharge", d.editDepositRecharge)
	_g.Post("/deleteDepositRecharge", d.deleteDepositRecharge)
	_g.Get("/pageSearchElderByKey", d.pageSearchElderByKey)
	_g.Get("/pageSearchStaffByKey", d.pageSearchStaffByKey)
	_g.Get("/getElderFeeById", d.getElderFeeById)
	_g.Post("/auditElderFee", d.auditElderFee)
}

// 分页查询押金充值
// @Summary 分页查询押金充值
// @Tags 押金充值
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageDepositRechargeByKeyQuery true "PageDepositRechargeByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /depositRecharge/pageDepositRechargeByKey [get]
func (d *depositrecharge) pageDepositRechargeByKey(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.PageDepositRechargeByKey)
}

// 获取押金充值
// @Summary 获取押金充值
// @Tags 押金充值
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /depositRecharge/getDepositRechargeById [get]
func (d *depositrecharge) getDepositRechargeById(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.GetDepositRechargeById)
}

// 新增押金充值
// @Summary 新增押金充值
// @Tags 押金充值
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddDepositRechargeQuery true "AddDepositRechargeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /depositRecharge/addDepositRecharge [post]
func (d *depositrecharge) addDepositRecharge(ctx ack.Context) {
	ack.Post(ctx, service.DepositRecharge.AddDepositRecharge)
}

// 编辑押金充值
// @Summary 编辑押金充值
// @Tags 押金充值
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditDepositRechargeQuery true "EditDepositRechargeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /depositRecharge/editDepositRecharge [post]
func (d *depositrecharge) editDepositRecharge(ctx ack.Context) {
	ack.Post(ctx, service.DepositRecharge.EditDepositRecharge)
}

// 删除押金充值
// @Summary 删除押金充值
// @Tags 押金充值
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /depositRecharge/deleteDepositRecharge [post]
func (d *depositrecharge) deleteDepositRecharge(ctx ack.Context) {
	ack.Post(ctx, service.DepositRecharge.DeleteDepositRecharge)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 押金充值
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyQuery true "PageSearchElderByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /depositRecharge/pageSearchElderByKey [get]
func (d *depositrecharge) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.PageSearchElderByKey)
}

// 分页查询员工
// @Summary 分页查询员工
// @Tags 押金充值
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchStaffByKeyQuery true "PageSearchStaffByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /depositRecharge/pageSearchStaffByKey [get]
func (d *depositrecharge) pageSearchStaffByKey(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.PageSearchStaffByKey)
}

// 获取老人费用
// @Summary 获取老人费用
// @Tags 押金充值
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /depositRecharge/getElderFeeById [get]
func (d *depositrecharge) getElderFeeById(ctx ack.Context) {
	ack.Get(ctx, service.DepositRecharge.GetElderFeeById)
}

// 审核老人费用
// @Summary 审核老人费用
// @Tags 押金充值
// @Accept application/json
// @Produce application/json
// @Param data body dto.AuditElderFeeQuery true "AuditElderFeeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /depositRecharge/auditElderFee [post]
func (d *depositrecharge) auditElderFee(ctx ack.Context) {
	ack.Post(ctx, service.DepositRecharge.AuditElderFee)
}
