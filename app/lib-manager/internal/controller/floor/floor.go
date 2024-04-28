package floor

import (
	"context"

	v1 "github.com/oldme-git/36hour/app/lib-manager/api/floor/v1"
	"github.com/oldme-git/36hour/app/lib-manager/api/pbentity"
	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
	"github.com/oldme-git/36hour/app/lib-manager/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedFloorServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterFloorServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := service.Floor().Create(ctx, &entity.Floor{
		LibId:     int(req.LibId),
		FloorName: req.FloorName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: int32(id)}, nil
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	floor, err := service.Floor().GetOne(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		Floor: &pbentity.Floor{
			Id:        int32(floor.Id),
			LibId:     int32(floor.LibId),
			FloorName: floor.FloorName,
			CreatedAt: timestamppb.New(floor.CreatedAt.Time),
			UpdatedAt: timestamppb.New(floor.UpdatedAt.Time),
		},
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	floors, err := service.Floor().GetList(ctx, &dao.FloorSearchCondition{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
		LibId:    int(req.LibId),
	})
	if err != nil {
		return nil, err
	}
	list := make([]*pbentity.Floor, len(floors))
	for k, v := range floors {
		list[k] = &pbentity.Floor{
			Id:        int32(v.Id),
			LibId:     int32(v.LibId),
			FloorName: v.FloorName,
			CreatedAt: timestamppb.New(v.CreatedAt.Time),
			UpdatedAt: timestamppb.New(v.UpdatedAt.Time),
		}
	}
	return &v1.GetListRes{Floors: list}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.Floor().Update(ctx, &entity.Floor{
		Id:        int(req.Id),
		FloorName: req.FloorName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Floor().Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{Id: req.Id}, nil
}
