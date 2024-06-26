package auth

import (
	"context"

	v1 "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/app/user/api/pbentity"
	"github.com/oldme-git/36hour/app/user/internal/logic/auth"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedAuthServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAuthServer(s.Server, &Controller{})
}

func (*Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := auth.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	res = &v1.LoginRes{
		Token: token,
	}
	return res, nil
}

func (*Controller) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	user, err := auth.GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUserInfoRes{
		User: &pbentity.User{
			Id:        int64(user.Id),
			Username:  user.Username,
			Phone:     user.Phone,
			CreatedAt: timestamppb.New(user.CreatedAt.Time),
			UpdatedAt: timestamppb.New(user.UpdatedAt.Time),
		},
	}
	return res, nil
}
