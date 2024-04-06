package hall

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "libManager/api/hall/v1"
	"libManager/api/pbentity"
	"libManager/internal/dao"
	"libManager/internal/model/entity"
	"libManager/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedHallServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterHallServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := service.Hall().Create(ctx, &entity.Hall{
		LibId:    int(req.LibId),
		FloorId:  int(req.FloorId),
		HallName: req.HallName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: int32(id)}, nil
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	hall, err := service.Hall().GetOne(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		Hall: &pbentity.Hall{
			Id:        int32(hall.Id),
			LibId:     int32(hall.LibId),
			FloorId:   int32(hall.FloorId),
			HallName:  hall.HallName,
			CreatedAt: timestamppb.New(hall.CreatedAt.Time),
			UpdatedAt: timestamppb.New(hall.UpdatedAt.Time),
		},
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	halls, err := service.Hall().GetList(ctx, &dao.HallSearchCondition{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
		LibId:    int(req.LibId),
		FloorId:  int(req.FloorId),
	})
	if err != nil {
		return nil, err
	}
	list := make([]*pbentity.Hall, len(halls))
	for k, v := range halls {
		list[k] = &pbentity.Hall{
			Id:        int32(v.Id),
			LibId:     int32(v.LibId),
			FloorId:   int32(v.FloorId),
			HallName:  v.HallName,
			CreatedAt: timestamppb.New(v.CreatedAt.Time),
			UpdatedAt: timestamppb.New(v.UpdatedAt.Time),
		}
	}
	return &v1.GetListRes{Halls: list}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.Hall().Update(ctx, &entity.Hall{
		Id:       int(req.Id),
		FloorId:  int(req.FloorId),
		HallName: req.HallName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Hall().Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{Id: req.Id}, nil
}
