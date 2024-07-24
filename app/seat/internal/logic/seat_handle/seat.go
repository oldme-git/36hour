package seat_handle

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/layout"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/seat"
	"github.com/oldme-git/36hour/app/seat/utility/cache"
	"github.com/oldme-git/36hour/utility"
)

// GetCell 获取座位信息
func GetCell(ctx context.Context, layoutId, cellNo int) (cell layout.Cell, err error) {
	var (
		redis     = g.Redis()
		cacheKey  = cache.LayoutMapKey(layoutId)
		cellNoStr = strconv.Itoa(cellNo)
	)
	cellGvar, err := redis.HGet(ctx, cacheKey, cellNoStr)
	if err != nil {
		return cell, utility.Err.NewSys(err)
	}
	if err = gconv.Struct(cellGvar.Val(), &cell); err != nil {
		return cell, utility.Err.NewSys(err)
	}
	return
}

// updateCellStatus 更新座位状态
func updateCellStatus(ctx context.Context, layoutId, cellNo int, status layout.CellStatus) (err error) {
	var (
		redis     = g.Redis()
		cacheKey  = cache.LayoutMapKey(layoutId)
		cell      layout.Cell
		cellNoStr = strconv.Itoa(cellNo)
	)
	cell, err = GetCell(ctx, layoutId, cellNo)
	if err != nil {
		return
	}

	_, err = redis.HSet(ctx, cacheKey, map[string]interface{}{
		cellNoStr: layout.Cell{
			X:       cell.X,
			Y:       cell.Y,
			VectorX: cell.VectorX,
			VectorY: cell.VectorY,
			No:      cell.No,
			Label:   cell.Label,
			Type:    cell.Type,
			Status:  status,
		},
	})
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return nil
}

// GetCellStatus 根据操作类型获取座位状态
func GetCellStatus(ctx context.Context, t seat.Type) (status layout.CellStatus) {
	switch t {
	case seat.TypeReserve:
	case seat.TypeSign:
		status = layout.CellStatusUsing
	case seat.TypeCancel:
		status = layout.CellStatusFree
	}
	return
}
