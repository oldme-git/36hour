package layout

import (
	"context"
	v1 "seat/api/layout/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedLayoutServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterLayoutServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateLayoutReq) (res *v1.CreateLayoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
