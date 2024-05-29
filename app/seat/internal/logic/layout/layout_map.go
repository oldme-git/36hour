package layout

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/oldme-git/36hour/app/seat/internal/model/layout"
	"github.com/oldme-git/36hour/utility"
)

func JsonToLayoutCells(ctx context.Context, jsonStr string) ([]layout.Cell, error) {
	var cells []layout.Cell
	if err := gjson.DecodeTo(jsonStr, &cells); err != nil {
		return nil, utility.Err.NewSys(err)
	}
	return cells, nil
}

func LayoutCellsToJson(ctx context.Context, cells []layout.Cell) (jsonStr string, err error) {
	return gjson.EncodeString(cells)
}

// CalculateSeatsByJson 根据 layout.Map 的 json 数据，计算出座位数
func CalculateSeatsByJson(ctx context.Context, jsonStr string) (seats int, err error) {
	cells, err := JsonToLayoutCells(ctx, jsonStr)
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
