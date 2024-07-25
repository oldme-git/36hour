package seat_handle

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	v1 "github.com/oldme-git/36hour/app/seat/api/seat_handle/v1"
	"github.com/oldme-git/36hour/app/seat/internal/imodel"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/seat"
	"github.com/oldme-git/36hour/app/seat/internal/logic/seat_handle"
)

type Controller struct {
	v1.UnimplementedSeatHandleServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterSeatHandleServer(s.Server, &Controller{})
}

func (*Controller) Reserve(ctx context.Context, req *v1.ReserveReq) (res *v1.ReserveRes, err error) {
	err = seat_handle.Reserve(ctx, &seat.UserSeat{
		Uid:        imodel.Id(req.Uid),
		LocationId: int(req.LocationId),
		CellNo:     int(req.CellNo),
		Type:       seat.TypeReserve,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
