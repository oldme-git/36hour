package seat

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/api/seat/v1"
	"github.com/oldme-git/36hour/app/gateway/internal/logic/seat"
)

func (c *ControllerV1) GetLayoutRuntime(ctx context.Context, req *v1.GetLayoutRuntimeReq) (res *v1.GetLayoutRuntimeRes, err error) {
	layout, err := seat.GetLayoutRuntimeMap(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetLayoutRuntimeRes{
		Map: layout.GetMap(),
	}, nil
}
