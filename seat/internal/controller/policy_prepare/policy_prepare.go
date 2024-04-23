package policy_prepare

import (
	"context"

	"seat/api/pbentity"
	v1 "seat/api/policy_prepare/v1"
	"seat/internal/model"
	"seat/internal/model/entity"
	"seat/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedPolicyPrepareServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPolicyPrepareServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := service.PolicyPrepare().Create(ctx, &entity.PolicyPrepare{
		Name: req.Name,
		Info: req.Policy,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: int32(id)}, nil
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	policy, err := service.PolicyPrepare().GetOne(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		PolicyPrepare: &pbentity.PolicyPrepare{
			Id:   int32(policy.Id),
			Name: policy.Name,
			Info: policy.Info,
		},
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	policies, err := service.PolicyPrepare().GetList(ctx, &model.PolicyPrepareSearchCondition{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	total, err := service.PolicyPrepare().GetTotal(ctx, &model.PolicyPrepareSearchCondition{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	list := make([]*pbentity.PolicyPrepare, len(policies))
	for k, v := range policies {
		list[k] = &pbentity.PolicyPrepare{
			Id:   int32(v.Id),
			Name: v.Name,
			Info: v.Info,
		}
	}
	return &v1.GetListRes{PolicyPrepares: list, Total: int32(total)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.PolicyPrepare().Update(ctx, &entity.PolicyPrepare{
		Id:   int(req.Id),
		Name: req.Name,
		Info: req.Policy,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.PolicyPrepare().Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return
}
