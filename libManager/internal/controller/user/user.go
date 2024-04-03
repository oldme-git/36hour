package user

import (
	"context"

	"github.com/gogf/gf/v2/util/gutil"
	v1 "user/api/user/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	gutil.Dump(req.GetUsername())
	gutil.Dump(req.GetPassword())
	return &v1.CreateRes{Id: 3}, nil
}
