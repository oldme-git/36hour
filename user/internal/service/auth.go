// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"user/internal/model"
)

type (
	IAuth interface {
		Login(ctx context.Context, Username, Password string) (token string, err error)
		Logout(ctx context.Context, id model.Id) (err error)
		Register(ctx context.Context) (err error)
		ChangePassword(ctx context.Context) (err error)
		ResetPassword(ctx context.Context) (err error)
		GetUserInfo(ctx context.Context, token string) (user *model.User, err error)
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
