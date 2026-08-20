package lib

import (
	"context"
	"strconv"
	"strings"

	"api/internal/constant"

	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"github.com/linbaozhong/gentity/pkg/token"
)

// ctxKey context key 类型，避免与其它包冲突
type ctxKey int

const (
	ctxUserID   ctxKey = iota // 全局用户编号
	ctxTenantID               // 租户编号
	ctxMemberID               // 成员编号
	ctxRole                   // 角色标识
)

// AuthMiddleware 登录鉴权中间件：
// 解析 token，注入 user_id/tenant_id/member_id/role 到请求上下文。
// 兼容旧 token（无租户信息，视为平台租户管理员，强制后续重新登录时更新）。
func AuthMiddleware() ack.Handler {
	return func(c ack.Context) {
		tk := c.GetHeader(constant.Authorization)
		if tk == "" {
			ack.Fail(c, constant.ErrAuthorizationNotFound)
			return
		}
		id, role, e := token.GetIDAndTokenFromCipher(tk)
		if e != nil {
			ack.Fail(c, constant.ErrTokenError)
			return
		}
		userID, tenantID, memberID := parseTokenID(id)

		ctx := context.WithValue(c.Request().Context(), ctxUserID, userID)
		ctx = context.WithValue(ctx, ctxTenantID, tenantID)
		ctx = context.WithValue(ctx, ctxMemberID, memberID)
		ctx = context.WithValue(ctx, ctxRole, role)
		c.ResetRequest(c.Request().WithContext(ctx))
		c.Next()
	}
}

// parseTokenID 解析 token 中的 id 段：格式为 userID[_tenantID[_memberID]]
// 旧 token 为纯 staff id（单段），无租户信息。
func parseTokenID(id string) (userID, tenantID, memberID int64) {
	parts := strings.Split(id, "_")
	userID = atoi64(parts[0])
	if len(parts) > 1 {
		tenantID = atoi64(parts[1])
	}
	if len(parts) > 2 {
		memberID = atoi64(parts[2])
	}
	return
}

func atoi64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

// UserID 从 context 读取当前登录全局用户编号
func UserID(ctx context.Context) int64 {
	v, _ := ctx.Value(ctxUserID).(int64)
	return v
}

// TenantID 从 context 读取当前租户编号
func TenantID(ctx context.Context) int64 {
	v, _ := ctx.Value(ctxTenantID).(int64)
	return v
}

// MemberID 从 context 读取当前成员编号
func MemberID(ctx context.Context) int64 {
	v, _ := ctx.Value(ctxMemberID).(int64)
	return v
}

// Role 从 context 读取当前角色标识
func Role(ctx context.Context) string {
	v, _ := ctx.Value(ctxRole).(string)
	return v
}
