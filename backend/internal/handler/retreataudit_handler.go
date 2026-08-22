package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type retreataudit struct{}

func init() {
	ack.RegisterRoute(&retreataudit{})
}

func (r *retreataudit) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/retreatAudit")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageRetreatAuditByKey", r.pageRetreatAuditByKey)
	_g.Get("/getRetreatAuditById", r.getRetreatAuditById)
	_g.Post("/auditRetreat", r.auditRetreat)
	_g.Get("/pageSearchElderByKey", r.pageSearchElderByKey)
}

// 分页查询退住Audit
// @Summary 分页查询退住Audit
// @Tags 退住审核
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageRetreatAuditByKeyReq true "PageRetreatAuditByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatAudit/pageRetreatAuditByKey [get]
func (r *retreataudit) pageRetreatAuditByKey(ctx ack.Context) {
	ack.Get(ctx, service.RetreatAudit.PageRetreatAuditByKey)
}

// 获取退住Audit
// @Summary 获取退住Audit
// @Tags 退住审核
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatAudit/getRetreatAuditById [get]
func (r *retreataudit) getRetreatAuditById(ctx ack.Context) {
	ack.Get(ctx, service.RetreatAudit.GetRetreatAuditById)
}

// 审核退住
// @Summary 审核退住
// @Tags 退住审核
// @Accept application/json
// @Produce application/json
// @Param data body dto.AuditRetreatReq true "AuditRetreatReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatAudit/auditRetreat [post]
func (r *retreataudit) auditRetreat(ctx ack.Context) {
	ack.Post(ctx, service.RetreatAudit.AuditRetreat)
}

// 分页查询老人
// @Summary 分页查询老人
// @Tags 退住审核
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageSearchElderByKeyReq true "PageSearchElderByKeyReq"
// @Success 200 {object} dto.EmptyResp
// @Router /retreatAudit/pageSearchElderByKey [get]
func (r *retreataudit) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.RetreatAudit.PageSearchElderByKey)
}
