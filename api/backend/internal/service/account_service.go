package service

import (
	"context"

	"api/internal/model/dto"
)

type account struct{}

var Account = &account{}

// Login 登录
func (a *account) Login(ctx context.Context, in *dto.LoginQuery, out *dto.LoginUserVO) error {
	// TODO: 实现登录逻辑
	return nil
}

// SendCode 发送验证码
func (a *account) SendCode(ctx context.Context, in *dto.SendCodeQuery, out *dto.EmptyResp) error {
	// TODO: 实现发送验证码逻辑
	return nil
}

// Forget 忘记密码
func (a *account) Forget(ctx context.Context, in *dto.ForgetQuery, out *dto.LoginUserVO) error {
	// TODO: 实现忘记密码逻辑
	return nil
}

// Edit 修改账户
func (a *account) Edit(ctx context.Context, in *dto.EditQuery, out *dto.EmptyResp) error {
	// TODO: 实现修改账户逻辑
	return nil
}

// Logout 退出登录
func (a *account) Logout(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// TODO: 实现退出登录逻辑
	return nil
}
