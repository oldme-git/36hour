// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IFloor interface {
		Create(ctx context.Context) (err error)
		GetOne(ctx context.Context) (err error)
		GetList(ctx context.Context) (err error)
		Update(ctx context.Context) (err error)
		Delete(ctx context.Context) (err error)
	}
)

var (
	localFloor IFloor
)

func Floor() IFloor {
	if localFloor == nil {
		panic("implement not found for interface IFloor, forgot register?")
	}
	return localFloor
}

func RegisterFloor(i IFloor) {
	localFloor = i
}
