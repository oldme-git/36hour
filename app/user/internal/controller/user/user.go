package user

import (
	"context"

	"github.com/oldme-git/36hour/app/user/api/pbentity"
	v1 "github.com/oldme-git/36hour/app/user/api/user/v1"
	"github.com/oldme-git/36hour/app/user/internal/logic/user"
	"github.com/oldme-git/36hour/app/user/internal/model"
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
	id, err := user.Create(ctx, &model.User{
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
	userOne, err := user.GetOne(ctx, model.Id(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		User: &pbentity.User{
			Id:        int64(userOne.Id),
			Username:  userOne.Username,
			Phone:     userOne.Phone,
			CreatedAt: timestamppb.New(userOne.CreatedAt.Time),
			UpdatedAt: timestamppb.New(userOne.UpdatedAt.Time),
		},
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	users, err := user.GetList(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	list := make([]*pbentity.User, len(users))
	for k, v := range users {
		list[k] = &pbentity.User{
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
	err = user.Update(ctx, &model.User{
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
	err = user.Delete(ctx, model.Id(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{
		Id: req.Id,
	}, nil
}

func (*Controller) BindLib(ctx context.Context, req *v1.BindLibRequest) (res *v1.BindLibResponse, err error) {
	err = user.BindLib(ctx, model.Id(req.UserId), model.Id(req.LibId))
	if err != nil {
		return nil, err
	}
	return &v1.BindLibResponse{}, nil
}
