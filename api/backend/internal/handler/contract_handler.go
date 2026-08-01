package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type contract struct{}

func init() {
	ack.RegisterRoute(&contract{})
}

func (c *contract) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/contract")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/contractExpireJob", c.contractExpireJob)
}

func (c *contract) contractExpireJob(ctx ack.Context) {
	ack.Get(ctx, service.Contract.ContractExpireJob)
}
