package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
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
	_g.Put("/auditRetreat", r.auditRetreat)
	_g.Get("/pageSearchElderByKey", r.pageSearchElderByKey)
}

func (r *retreataudit) pageRetreatAuditByKey(ctx ack.Context) {
	ack.Get(ctx, service.RetreatAudit.PageRetreatAuditByKey)
}

func (r *retreataudit) getRetreatAuditById(ctx ack.Context) {
	ack.Get(ctx, service.RetreatAudit.GetRetreatAuditById)
}

func (r *retreataudit) auditRetreat(ctx ack.Context) {
	ack.Put(ctx, service.RetreatAudit.AuditRetreat)
}

func (r *retreataudit) pageSearchElderByKey(ctx ack.Context) {
	ack.Get(ctx, service.RetreatAudit.PageSearchElderByKey)
}
