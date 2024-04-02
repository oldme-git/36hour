// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IAuth interface {
		Login(ctx context.Context) (token string, err error)
		Logout(ctx context.Context) (err error)
		Register(ctx context.Context) (err error)
		ChangePassword(ctx context.Context) (err error)
		ResetPassword(ctx context.Context) (err error)
		GetUserInfo(ctx context.Context) (err error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
