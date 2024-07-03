package seat

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/api/seat/v1"
	"github.com/oldme-git/36hour/app/gateway/internal/logic/seat"
	"github.com/oldme-git/36hour/app/gateway/utility/data"
)

func (c *ControllerV1) GetLayout(ctx context.Context, req *v1.GetLayoutReq) (res *v1.GetLayoutRes, err error) {
	layout, err := seat.GetLayout(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetLayoutRes{
		Id:           data.Id(layout.Layout.GetId()),
		LocationId:   data.Id(layout.Layout.GetLocationId()),
		LocationName: layout.GetLocationName(),
		PolicyCId:    data.Id(layout.Layout.GetPolicyCId()),
		LayoutName:   layout.Layout.GetLayoutName(),
		Map:          layout.Layout.GetMap(),
	}, nil
}
