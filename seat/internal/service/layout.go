// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"seat/internal/dao"
	"seat/internal/model/entity"
)

type (
	ILayout interface {
		Create(ctx context.Context, layout *entity.Layout) (id int, err error)
		GetOne(ctx context.Context, id int) (layout *entity.Layout, err error)
		GetList(ctx context.Context, condition *dao.LayoutSearchCondition) (layouts []*entity.Layout, err error)
		Update(ctx context.Context, layout *entity.Layout) (err error)
		Delete(ctx context.Context, id int) (err error)
	}
)

var (
	localLayout ILayout
)

func Layout() ILayout {
	if localLayout == nil {
		panic("implement not found for interface ILayout, forgot register?")
	}
	return localLayout
}

func RegisterLayout(i ILayout) {
	localLayout = i
}
