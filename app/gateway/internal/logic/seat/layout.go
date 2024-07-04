package seat

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/internal/logic/account"
	"github.com/oldme-git/36hour/app/gateway/utility/data"
	locationSvc "github.com/oldme-git/36hour/app/lib-manager/api/location/v1"
	layoutSvc "github.com/oldme-git/36hour/app/seat/api/layout/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"
)

// GetLayouts 获取座位布局列表
func GetLayouts(ctx context.Context, token string) (*layoutSvc.GetRuntimeLayoutRes, error) {
	userInfo, err := account.GetInfo(ctx, token)
	if err != nil {
		return nil, err
	}

	// 获取location列表
	var (
		libConn        = svc_disc.LibManagerClientConn(ctx)
		locationClient = locationSvc.NewLocationClient(libConn)
		locationIds    []uint64
	)
	locations, err := locationClient.GetList(ctx, &locationSvc.GetListReq{
		Page:     1,
		PageSize: 0,
		LibId:    int32(userInfo.GetLib().GetLibId()),
		FloorId:  0,
	})
	if err != nil {
		return nil, err
	}
	for _, location := range locations.GetLocations() {
		locationIds = append(locationIds, uint64(location.Id))
	}

	// 根据location列表获取座位布局列表
	var (
		seatConn     = svc_disc.SeatClientConn(ctx)
		layoutClient = layoutSvc.NewLayoutClient(seatConn)
	)
	return layoutClient.GetRuntimeLayouts(ctx, &layoutSvc.GetRuntimeLayoutReq{
		LocationIds: locationIds,
	})
}

// GetLayout 获取座位布局详情
func GetLayout(ctx context.Context, layoutId data.Id) (*layoutSvc.GetOneRes, error) {
	var (
		seatConn     = svc_disc.SeatClientConn(ctx)
		layoutClient = layoutSvc.NewLayoutClient(seatConn)
	)
	return layoutClient.GetOne(ctx, &layoutSvc.GetOneReq{
		Id: int32(layoutId),
	})
}

// GetLayoutRuntimeMap 获取座位布局运行时的 Map
func GetLayoutRuntimeMap(ctx context.Context, layoutId data.Id) (*layoutSvc.GetRuntimeLayoutMapRes, error) {
	var (
		seatConn     = svc_disc.SeatClientConn(ctx)
		layoutClient = layoutSvc.NewLayoutClient(seatConn)
	)
	return layoutClient.GetRuntimeLayoutMap(ctx, &layoutSvc.GetRuntimeLayoutMapReq{
		LayoutId: int32(layoutId),
	})
}
