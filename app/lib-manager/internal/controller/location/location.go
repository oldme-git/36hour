package location

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	v1 "github.com/oldme-git/36hour/app/lib-manager/api/location/v1"
	"github.com/oldme-git/36hour/app/lib-manager/api/pbentity"
	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/location"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Controller struct {
	v1.UnimplementedLocationServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterLocationServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := location.Create(ctx, &entity.Location{
		LibId:        int(req.LibId),
		FloorId:      int(req.FloorId),
		LocationName: req.LocationName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: int32(id)}, nil
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	locationOne, err := location.GetOne(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		Location: &pbentity.Location{
			Id:           int32(locationOne.Id),
			LibId:        int32(locationOne.LibId),
			FloorId:      int32(locationOne.FloorId),
			LocationName: locationOne.LocationName,
			CreatedAt:    timestamppb.New(locationOne.CreatedAt.Time),
			UpdatedAt:    timestamppb.New(locationOne.UpdatedAt.Time),
		},
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	locations, err := location.GetList(ctx, &dao.LocationSearchCondition{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
		LibId:    int(req.LibId),
		FloorId:  int(req.FloorId),
	})
	if err != nil {
		return nil, err
	}
	list := make([]*pbentity.Location, len(locations))
	for k, v := range locations {
		list[k] = &pbentity.Location{
			Id:           int32(v.Id),
			LibId:        int32(v.LibId),
			FloorId:      int32(v.FloorId),
			LocationName: v.LocationName,
			CreatedAt:    timestamppb.New(v.CreatedAt.Time),
			UpdatedAt:    timestamppb.New(v.UpdatedAt.Time),
		}
	}
	return &v1.GetListRes{Locations: list}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = location.Update(ctx, &entity.Location{
		Id:           int(req.Id),
		FloorId:      int(req.FloorId),
		LocationName: req.LocationName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = location.Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{Id: req.Id}, nil
}
