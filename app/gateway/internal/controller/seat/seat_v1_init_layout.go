package seat

import (
	"context"

	"github.com/oldme-git/36hour/utility/svc_disc"

	"github.com/oldme-git/36hour/app/gateway/api/seat/v1"
	layoutSvc "github.com/oldme-git/36hour/app/seat/api/layout/v1"
)

func (c *ControllerV1) InitLayout(ctx context.Context, req *v1.InitLayoutReq) (res *v1.InitLayoutRes, err error) {
	var (
		seatConn     = svc_disc.SeatClientConn(ctx)
		layoutClient = layoutSvc.NewLayoutClient(seatConn)
	)

	ctx, cancel := context.WithTimeout(ctx, svc_disc.Timeout)
	defer cancel()

	_, err = layoutClient.InitLayout(ctx, &layoutSvc.InitLayoutReq{
		LayoutId: int64(req.Id),
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
