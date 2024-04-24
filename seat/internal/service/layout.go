// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"seat/internal/model"
	"seat/internal/model/entity"
	"seat/internal/model/layout"
)

type (
	ILayout interface {
		Create(ctx context.Context, layout *entity.Layout) (id int, err error)
		GetOne(ctx context.Context, id int) (layout *entity.Layout, err error)
		GetList(ctx context.Context, condition *model.LayoutSearchCondition) (layouts []*entity.Layout, err error)
		// GetTotal 获取 Layout 总数
		GetTotal(ctx context.Context, condition *model.LayoutSearchCondition) (total int, err error)
		Update(ctx context.Context, layout *entity.Layout) (err error)
		Delete(ctx context.Context, id int) (err error)
		JsonToLayoutCells(ctx context.Context, jsonStr string) ([]layout.Cell, error)
		// CalculateSeatsByJson 根据 layout.Map 的 json 数据，计算出座位数
		CalculateSeatsByJson(ctx context.Context, jsonStr string) (seats int, err error)
		// GetRuntimePolicy 获取运行中的策略
		// 如果有公共策略，获取公共策略信息
		// 如果没有，则返回自己的策略信息
		GetRuntimePolicy(ctx context.Context, layout *entity.Layout) (policyInfo string, err error)
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
