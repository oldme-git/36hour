package user

import (
	"context"

	"user/api/pbentity"
	v1 "user/api/user/v1"
	"user/internal/model"
	"user/internal/service"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := service.User().Create(ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.CreateRes{
		Id: int64(id),
	}
	return res, nil
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	user, err := service.User().GetOne(ctx, model.Id(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		User: &pbentity.UserMain{
			Id:        int64(user.Id),
			Username:  user.Username,
			Phone:     user.Phone,
			CreatedAt: timestamppb.New(user.CreatedAt.Time),
			UpdatedAt: timestamppb.New(user.UpdatedAt.Time),
		},
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	users, err := service.User().GetList(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	list := make([]*pbentity.UserMain, len(users))
	for k, v := range users {
		list[k] = &pbentity.UserMain{
			Id:        int64(v.Id),
			Username:  v.Username,
			Phone:     v.Phone,
			CreatedAt: timestamppb.New(v.CreatedAt.Time),
			UpdatedAt: timestamppb.New(v.UpdatedAt.Time),
		}
	}
	return &v1.GetListRes{
		Users: list,
	}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.User().Update(ctx, &model.User{
		Id:       model.Id(req.Id),
		Username: req.Username,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRes{
		Id: req.Id,
	}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.User().Delete(ctx, model.Id(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{
		Id: req.Id,
	}, nil
}
