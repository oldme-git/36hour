package layout

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"seat/internal/model/layout"
)

func (s *sLayout) JsonToLayoutCells(ctx context.Context, json string) ([]layout.Cell, error) {
	var cells []layout.Cell
	err := gconv.Structs(json, &cells)
	if err != nil {
		return nil, err
	}
	return cells, nil
}
