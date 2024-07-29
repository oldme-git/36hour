package seat

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/internal/logic/account"
	seatHandle "github.com/oldme-git/36hour/app/seat/api/seat_handle/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"

	"github.com/oldme-git/36hour/app/gateway/api/seat/v1"
)

func (c *ControllerV1) Reserve(ctx context.Context, req *v1.ReserveReq) (res *v1.ReserveRes, err error) {
	var (
		seatConn         = svc_disc.SeatClientConn(ctx)
		seatHandleClient = seatHandle.NewSeatHandleClient(seatConn)
	)
	ctx, cancel := context.WithTimeout(ctx, svc_disc.Timeout)
	defer cancel()

	var (
		token = account.GetToken(ctx)
	)
	userInfo, err := account.GetInfo(ctx, token)
	if err != nil {
		return nil, err
	}

	_, err = seatHandleClient.Reserve(ctx, &seatHandle.ReserveReq{
		Uid:        userInfo.GetUser().GetId(),
		LocationId: int64(req.LocationId),
		CellNo:     int64(req.CellNo),
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
