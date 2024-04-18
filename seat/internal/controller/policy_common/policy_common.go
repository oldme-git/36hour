package policy_common

import (
	"context"
	v1 "seat/api/policy_common/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedPolicyCommonServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPolicyCommonServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreatePolicyCommonReq) (res *v1.CreatePolicyCommonRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOnePolicyCommonReq) (res *v1.GetOnePolicyCommonRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListPolicyCommonReq) (res *v1.GetListPolicyCommonRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Update(ctx context.Context, req *v1.UpdatePolicyCommonReq) (res *v1.UpdatePolicyCommonRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Delete(ctx context.Context, req *v1.DeletePolicyCommonReq) (res *v1.DeletePolicyCommonRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
