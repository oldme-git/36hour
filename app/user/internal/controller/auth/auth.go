package auth

import (
	"context"

	libPb "github.com/oldme-git/36hour/app/lib-manager/api/lib/v1"
	v1 "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/app/user/api/pbentity"
	userPb "github.com/oldme-git/36hour/app/user/api/user/v1"
	"github.com/oldme-git/36hour/app/user/internal/logic/auth"
	"github.com/oldme-git/36hour/app/user/internal/logic/user"
	"github.com/oldme-git/36hour/app/user/internal/model"
	"github.com/oldme-git/36hour/utility/svc_disc"
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
	var (
		userOne   *model.User
		userLibId model.Id
	)
	userOne, err = auth.GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	userLibId, err = user.GetUserLibId(ctx, userOne.Id)
	if err != nil {
		return nil, err
	}

	// 读取场馆信息
	var (
		conn   = svc_disc.LibManagerClientConn(ctx)
		client = libPb.NewLibClient(conn)
		lib    *libPb.GetOneRes
	)

	lib, err = client.GetOne(ctx, &libPb.GetOneReq{
		Id: int32(userLibId),
	})
	if err != nil {
		return nil, err
	}

	res = &v1.GetUserInfoRes{
		User: &pbentity.User{
			Id:        int64(userOne.Id),
			Username:  userOne.Username,
			Phone:     userOne.Phone,
			CreatedAt: timestamppb.New(userOne.CreatedAt.Time),
			UpdatedAt: timestamppb.New(userOne.UpdatedAt.Time),
		},
		Lib: &userPb.Lib{
			LibId:   int64(lib.GetLib().GetId()),
			LibName: lib.GetLib().GetLibName(),
		},
	}
	return res, nil
}
