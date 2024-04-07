// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"libManager/internal/dao"
	"libManager/internal/model/entity"
)

type (
	ILib interface {
		Create(ctx context.Context, lib *entity.Lib) (id int, err error)
		GetOne(ctx context.Context, id int) (lib *entity.Lib, err error)
		GetList(ctx context.Context, condition *dao.LibSearchCondition) (libs []*entity.Lib, err error)
		Update(ctx context.Context, lib *entity.Lib) (err error)
		Delete(ctx context.Context, id int) (err error)
		Exist(ctx context.Context, id int) error
	}
)

var (
	localLib ILib
)

func Lib() ILib {
	if localLib == nil {
		panic("implement not found for interface ILib, forgot register?")
	}
	return localLib
}

func RegisterLib(i ILib) {
	localLib = i
}
