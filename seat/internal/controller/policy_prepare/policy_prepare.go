package policy_prepare

import (
	"context"
	v1 "seat/api/policy_prepare/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedPolicyPrepareServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPolicyPrepareServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreatePolicyPrepareReq) (res *v1.CreatePolicyPrepareRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOnePolicyPrepareReq) (res *v1.GetOnePolicyPrepareRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListPolicyPrepareReq) (res *v1.GetListPolicyPrepareRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Update(ctx context.Context, req *v1.UpdatePolicyPrepareReq) (res *v1.UpdatePolicyPrepareRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Delete(ctx context.Context, req *v1.DeletePolicyPrepareReq) (res *v1.DeletePolicyPrepareRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
