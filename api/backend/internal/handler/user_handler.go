package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"github.com/linbaozhong/gentity/pkg/cachego/mmap"
	"api/internal/lib"
	"api/internal/service"
)

type user struct{}

var (
	idempotencyConfig = ack.DefaultIdempotencyConfig(mmap.New())
)

func init() {
	ack.RegisterRoute(&user{})
}

func (u *user) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/user")

	// 不需要登录
	_g.Post("/user_register", u.userRegister)

	// 需要登录
	_g.Use(lib.AuthMiddleware())
	_g.Get("/get", u.get)
}

func (u *user) userRegister(ctx ack.Context) {
	// 如果不需要保持接口的幂等性，则使用下面的 Post 方法
	ack.Post(ctx, service.User.Register)
	// 如果需要保持接口的幂等性，则使用下面的 PostIdempotent 方法
	ack.PostIdempotent(ctx, idempotencyConfig, service.User.Register)
}

func (u *user) get(ctx ack.Context) {
	ack.Get(ctx, service.User.Get)
}