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
	IHall interface {
		Create(ctx context.Context, hall *entity.Hall) (id int, err error)
		GetOne(ctx context.Context, id int) (hall *entity.Hall, err error)
		GetList(ctx context.Context, condition *dao.HallSearchCondition) (halls []*entity.Hall, err error)
		Update(ctx context.Context, hall *entity.Hall) (err error)
		Delete(ctx context.Context, id int) (err error)
	}
)

var (
	localHall IHall
)

func Hall() IHall {
	if localHall == nil {
		panic("implement not found for interface IHall, forgot register?")
	}
	return localHall
}

func RegisterHall(i IHall) {
	localHall = i
}
