package lib

import (
	"context"

	v1 "github.com/oldme-git/36hour/app/lib-manager/api/lib/v1"
	"github.com/oldme-git/36hour/app/lib-manager/api/pbentity"
	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
	"github.com/oldme-git/36hour/app/lib-manager/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedLibServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterLibServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := service.Lib().Create(ctx, &entity.Lib{
		LibName: req.LibName,
		Address: req.Address,
		Active:  req.Active,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: int32(id)}, nil
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	lib, err := service.Lib().GetOne(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		Lib: &pbentity.Lib{
			Id:        int32(lib.Id),
			LibName:   lib.LibName,
			Address:   lib.Address,
			Active:    lib.Active,
			CreatedAt: timestamppb.New(lib.CreatedAt.Time),
			UpdatedAt: timestamppb.New(lib.UpdatedAt.Time),
		},
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	libs, err := service.Lib().GetList(ctx, &dao.LibSearchCondition{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
		LibName:  req.LibName,
		Address:  req.Address,
		Active:   req.Active,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*pbentity.Lib, len(libs))
	for k, v := range libs {
		list[k] = &pbentity.Lib{
			Id:        int32(v.Id),
			LibName:   v.LibName,
			Address:   v.Address,
			Active:    v.Active,
			CreatedAt: timestamppb.New(v.CreatedAt.Time),
			UpdatedAt: timestamppb.New(v.UpdatedAt.Time),
		}
	}
	return &v1.GetListRes{Libs: list}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.Lib().Update(ctx, &entity.Lib{
		Id:      int(req.Id),
		LibName: req.LibName,
		Address: req.Address,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Lib().Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{Id: req.Id}, nil
}
