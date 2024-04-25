package layout

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"seat/internal/model/layout"
)

func (s *sLayout) JsonToLayoutCells(ctx context.Context, jsonStr string) ([]layout.Cell, error) {
	var cells []layout.Cell
	if err := gjson.DecodeTo(jsonStr, &cells); err != nil {
		return nil, err
	}
	return cells, nil
}

func (s *sLayout) LayoutCellsToJson(ctx context.Context, cells []layout.Cell) (jsonStr string, err error) {
	return gjson.EncodeString(cells)
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
