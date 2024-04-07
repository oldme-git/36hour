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
	IFloor interface {
		Create(ctx context.Context, floor *entity.Floor) (id int, err error)
		GetOne(ctx context.Context, id int) (floor *entity.Floor, err error)
		GetList(ctx context.Context, condition *dao.FloorSearchCondition) (floors []*entity.Floor, err error)
		Update(ctx context.Context, floor *entity.Floor) (err error)
		Delete(ctx context.Context, id int) (err error)
		Exist(ctx context.Context, id int) error
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
