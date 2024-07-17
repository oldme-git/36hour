package layout

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	libLocation "github.com/oldme-git/36hour/app/lib-manager/api/location/v1"
	v1 "github.com/oldme-git/36hour/app/seat/api/layout/v1"
	"github.com/oldme-git/36hour/app/seat/api/pbentity"
	"github.com/oldme-git/36hour/app/seat/internal/logic/layout"
	"github.com/oldme-git/36hour/app/seat/internal/logic/policy_layout"
	"github.com/oldme-git/36hour/app/seat/internal/model"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
	"github.com/oldme-git/36hour/utility"

	"github.com/oldme-git/36hour/utility/svc_disc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedLayoutServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterLayoutServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	var policyLid int

	if req.PolicyInfo != "" {
		// 保存自己的策略信息
		policyLid, err = policy_layout.Create(ctx, &entity.PolicyLayout{
			Info: req.PolicyInfo,
		})
		if err != nil {
			return nil, err
		}
	}
	id, err := layout.Create(ctx, &entity.Layout{
		LocationId: int(req.LocationId),
		PolicyCId:  int(req.PolicyCId),
		PolicyLId:  policyLid,
		LayoutName: req.LayoutName,
		Map:        req.Map,
		Status:     int(req.Status),
		Sort:       int(req.Sort),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: int32(id)}, nil
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	layoutOne, err := layout.GetOne(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	// 获取运行中的策略
	policyInfo, err := layout.GetRuntimePolicy(ctx, layoutOne)
	if err != nil {
		return nil, err
	}

	// 获取场馆名称
	var (
		conn   = svc_disc.LibManagerClientConn(ctx)
		client = libLocation.NewLocationClient(conn)
	)

	ctx, cancel := context.WithTimeout(ctx, svc_disc.Timeout)
	defer cancel()
	lib, err := client.GetOne(ctx, &libLocation.GetOneReq{Id: int32(layoutOne.LocationId)})
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		Layout: &pbentity.Layout{
			Id:         int32(layoutOne.Id),
			LocationId: int32(layoutOne.LocationId),
			PolicyCId:  int32(layoutOne.PolicyCId),
			PolicyLId:  int32(layoutOne.PolicyLId),
			LayoutName: layoutOne.LayoutName,
			Map:        layoutOne.Map,
			Status:     int32(layoutOne.Status),
			Sort:       int32(layoutOne.Sort),
			Seats:      int32(layoutOne.Seats),
			CreatedAt:  timestamppb.New(layoutOne.CreatedAt.Time),
			UpdatedAt:  timestamppb.New(layoutOne.UpdatedAt.Time),
		},
		PolicyInfo:   policyInfo,
		LocationName: lib.GetLocation().GetLocationName(),
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	layouts, total, err := layout.GetList(ctx, &model.LayoutSearchCondition{
		Page:       int(req.Page),
		PageSize:   int(req.PageSize),
		LayoutName: req.LayoutName,
		Status:     int(req.Status),
		SeatsMin:   int(req.SeatsMin),
		SeatsMax:   int(req.SeatsMax),
	})
	if err != nil {
		return nil, err
	}

	list := make([]*v1.LayoutList, len(layouts))
	for k, v := range layouts {
		list[k] = &v1.LayoutList{
			Id:         int32(v.Id),
			LocationId: int32(v.LocationId),
			LayoutName: v.LayoutName,
			Status:     int32(v.Status),
			Sort:       int32(v.Sort),
			Seats:      int32(v.Seats),
		}
	}
	return &v1.GetListRes{Layouts: list, Total: int32(total)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	var policyLid int

	if req.PolicyInfo != "" {
		// TODO 自己的策略信息不能使用更新创建，需要调整
		// 保存自己的策略信息
		policyLid, err = policy_layout.Create(ctx, &entity.PolicyLayout{
			Info: req.PolicyInfo,
		})
		if err != nil {
			return nil, err
		}
	}
	err = layout.Update(ctx, &entity.Layout{
		Id:         int(req.Id),
		LocationId: int(req.LocationId),
		PolicyCId:  int(req.PolicyCId),
		PolicyLId:  policyLid,
		LayoutName: req.LayoutName,
		Map:        req.Map,
		Status:     int(req.Status),
		Sort:       int(req.Sort),
	})
	if err != nil {
		return nil, err
	}
	return
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = layout.Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return
}

func (*Controller) GetRuntimeLayouts(ctx context.Context, req *v1.GetRuntimeLayoutReq) (res *v1.GetRuntimeLayoutRes, err error) {
	var locationIds = make([]int, len(req.LocationIds))
	for k, v := range req.LocationIds {
		locationIds[k] = int(v)
	}
	layouts, err := layout.GetRuntimeLayout(ctx, locationIds...)
	if err != nil {
		return nil, err
	}

	var layoutsRes = make([]*pbentity.Layout, len(layouts))
	for k, v := range layouts {
		layoutsRes[k] = &pbentity.Layout{
			Id:         int32(v.Id),
			LocationId: int32(v.LocationId),
			PolicyCId:  int32(v.PolicyCId),
			PolicyLId:  int32(v.PolicyLId),
			LayoutName: v.LayoutName,
			Map:        v.Map,
			Status:     int32(v.Status),
			Sort:       int32(v.Sort),
			Seats:      int32(v.Seats),
			CreatedAt:  timestamppb.New(v.CreatedAt.Time),
			UpdatedAt:  timestamppb.New(v.UpdatedAt.Time),
		}
	}
	return &v1.GetRuntimeLayoutRes{
		Layouts: layoutsRes,
	}, nil
}

func (*Controller) GetRuntimeLayoutMap(ctx context.Context, req *v1.GetRuntimeLayoutMapReq) (res *v1.GetRuntimeLayoutMapRes, err error) {
	var layoutId = int(req.LayoutId)
	cells, err := layout.GetRuntimeLayoutMap(ctx, layoutId)
	if err != nil {
		return nil, err
	}
	cellsJson, err := gjson.EncodeString(cells)
	if err != nil {
		return nil, utility.Err.NewSys(err)
	}
	return &v1.GetRuntimeLayoutMapRes{Map: cellsJson}, nil
}
