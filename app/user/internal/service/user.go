// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/oldme-git/36hour/app/user/internal/model"
)

type (
	IUser interface {
		Create(ctx context.Context, user *model.User) (id model.Id, err error)
		GetList(ctx context.Context, page int, pageSize int) (users []*model.User, err error)
		GetOne(ctx context.Context, id model.Id) (user *model.User, err error)
		GetOneByUsername(ctx context.Context, username string) (user *model.User, err error)
		// Update 不会修改密码
		Update(ctx context.Context, user *model.User) (err error)
		Delete(ctx context.Context, id model.Id) (err error)
		// CheckPassword 输入明文和密码，判断密码是否正确
		CheckPassword(ctx context.Context, plainPwd, hashedPwd string) (bool, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
