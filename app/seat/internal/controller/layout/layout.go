package layout

import (
	"context"

	v1 "github.com/oldme-git/36hour/app/seat/api/layout/v1"
	"github.com/oldme-git/36hour/app/seat/api/pbentity"
	"github.com/oldme-git/36hour/app/seat/internal/model"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
	"github.com/oldme-git/36hour/app/seat/internal/service"
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
		policyLid, err = service.PolicyLayout().Create(ctx, &entity.PolicyLayout{
			Info: req.PolicyInfo,
		})
		if err != nil {
			return nil, err
		}
	}
	id, err := service.Layout().Create(ctx, &entity.Layout{
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
	layout, err := service.Layout().GetOne(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	// 获取运行中的策略
	policyInfo, err := service.Layout().GetRuntimePolicy(ctx, layout)
	if err != nil {
		return nil, err
	}

	return &v1.GetOneRes{
		Layout: &pbentity.Layout{
			Id:         int32(layout.Id),
			LocationId: int32(layout.LocationId),
			PolicyCId:  int32(layout.PolicyCId),
			PolicyLId:  int32(layout.PolicyLId),
			LayoutName: layout.LayoutName,
			Map:        layout.Map,
			Status:     int32(layout.Status),
			Sort:       int32(layout.Sort),
			Seats:      int32(layout.Seats),
			CreatedAt:  timestamppb.New(layout.CreatedAt.Time),
			UpdatedAt:  timestamppb.New(layout.UpdatedAt.Time),
		},
		PolicyInfo: policyInfo,
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	layouts, err := service.Layout().GetList(ctx, &model.LayoutSearchCondition{
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
	total, err := service.Layout().GetTotal(ctx, &model.LayoutSearchCondition{
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
			// TODO, 需要去libManager中获取LocationName
			LocationName: "",
			LayoutName:   v.LayoutName,
			Status:       int32(v.Status),
			Sort:         int32(v.Sort),
			Seats:        int32(v.Seats),
		}
	}
	return &v1.GetListRes{Layouts: list, Total: int32(total)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	var policyLid int

	if req.PolicyInfo != "" {
		// TODO 自己的策略信息不能使用更新创建，需要调整
		// 保存自己的策略信息
		policyLid, err = service.PolicyLayout().Create(ctx, &entity.PolicyLayout{
			Info: req.PolicyInfo,
		})
		if err != nil {
			return nil, err
		}
	}
	err = service.Layout().Update(ctx, &entity.Layout{
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
	err = service.Layout().Delete(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return
}
