package layout

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"seat/internal/model/layout"
	"seat/internal/packed"
)

func (s *sLayout) JsonToLayoutCells(ctx context.Context, jsonStr string) ([]layout.Cell, error) {
	var cells []layout.Cell
	err := gconv.Structs(jsonStr, &cells)
	if err != nil {
		return nil, err
	}
	if jsonStr != "" && len(cells) == 0 {
		return nil, packed.Err.New(3003)
	}
	return cells, nil
}

// CalculateSeatsByJson 根据 layout.Map 的 json 数据，计算出座位数
func (s *sLayout) CalculateSeatsByJson(ctx context.Context, jsonStr string) (seats int, err error) {
	cells, err := s.JsonToLayoutCells(ctx, jsonStr)
	if err != nil {
		return
	}
	for _, cell := range cells {
		if cell.Type == layout.CellSeat {
			seats++
		}
	}
	return
}
