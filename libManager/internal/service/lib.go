// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ILib interface {
		Create(ctx context.Context) (err error)
		GetOne(ctx context.Context) (err error)
		GetList(ctx context.Context) (err error)
		Update(ctx context.Context) (err error)
		Delete(ctx context.Context) (err error)
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
