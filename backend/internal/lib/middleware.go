package lib

import (
	"api/internal/constant"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

func AuthMiddleware() ack.Handler {
	return func(c ack.Context) {
		token := c.GetHeader(constant.Authorization)
		if token == "" {
			ack.Fail(c, constant.ErrAuthorizationNotFound)
			return
		}
		c.Next()
	}
}
