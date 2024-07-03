package seat

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/internal/logic/account"
	"github.com/oldme-git/36hour/app/gateway/internal/logic/seat"

	"github.com/oldme-git/36hour/app/gateway/api/seat/v1"
)

func (c *ControllerV1) GetLayouts(ctx context.Context, req *v1.GetLayoutsReq) (res *v1.GetLayoutsRes, err error) {
	token := account.GetToken(ctx)
	layoutData, err := seat.GetLayouts(ctx, token)
	if err != nil {
		return nil, err
	}
	var layouts []v1.Layout
	for _, layout := range layoutData.Layouts {
		layouts = append(layouts, v1.Layout{
			Id:    int(layout.Id),
			Name:  layout.LayoutName,
			Seats: int(layout.Seats),
		})
	}

	return &v1.GetLayoutsRes{Layouts: layouts}, nil
}
