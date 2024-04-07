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
	ILocation interface {
		Create(ctx context.Context, location *entity.Location) (id int, err error)
		GetOne(ctx context.Context, id int) (location *entity.Location, err error)
		GetList(ctx context.Context, condition *dao.LocationSearchCondition) (locations []*entity.Location, err error)
		Update(ctx context.Context, location *entity.Location) (err error)
		Delete(ctx context.Context, id int) (err error)
	}
)

var (
	localLocation ILocation
)

func Location() ILocation {
	if localLocation == nil {
		panic("implement not found for interface ILocation, forgot register?")
	}
	return localLocation
}

func RegisterLocation(i ILocation) {
	localLocation = i
}
