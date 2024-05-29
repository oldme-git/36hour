package floor

import (
	"context"

	v1 "github.com/oldme-git/36hour/app/lib-manager/api/floor/v1"
	"github.com/oldme-git/36hour/app/lib-manager/api/pbentity"
	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/floor"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
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
	id, err := floor.Create(ctx, &entity.Floor{
		LibId:     int(req.LibId),
		FloorName: req.FloorName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: int32(id)}, nil
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	floorOne, err := floor.GetOne(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		Floor: &pbentity.Floor{
			Id:        int32(floorOne.Id),
			LibId:     int32(floorOne.LibId),
			FloorName: floorOne.FloorName,
			CreatedAt: timestamppb.New(floorOne.CreatedAt.Time),
			UpdatedAt: timestamppb.New(floorOne.UpdatedAt.Time),
		},
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	floors, err := floor.GetList(ctx, &dao.FloorSearchCondition{
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
	err = floor.Update(ctx, &entity.Floor{
		Id:        int(req.Id),
		FloorName: req.FloorName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = floor.Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{Id: req.Id}, nil
}
