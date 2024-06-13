package layout

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/oldme-git/36hour/app/seat/internal/dao"
	"github.com/oldme-git/36hour/app/seat/internal/logic/policy_common"
	"github.com/oldme-git/36hour/app/seat/internal/logic/policy_layout"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
	"github.com/oldme-git/36hour/app/seat/internal/model/layout"
	"github.com/oldme-git/36hour/app/seat/utility/cache_key"
	"github.com/oldme-git/36hour/utility"
)

// GetRuntimePolicy 获取运行中的策略
// 如果有公共策略，获取公共策略信息
// 如果没有，则返回自己的策略信息
func GetRuntimePolicy(ctx context.Context, layout *entity.Layout) (policyInfo string, err error) {
	if layout.PolicyCId != 0 {
		policyCommon, err := policy_common.GetOne(ctx, layout.PolicyCId)
		if err != nil {
			return "", err
		}
		policyInfo = policyCommon.Info
	} else {
		policyLayout, err := policy_layout.GetOne(ctx, layout.PolicyLId)
		if err != nil {
			return "", err
		}
		policyInfo = policyLayout.Info
	}
	return
}

// GetRuntimeLayout 获取运行中的布局列表
func GetRuntimeLayout(ctx context.Context, locationIds ...int) (layouts []entity.Layout, err error) {
	layouts = make([]entity.Layout, 0)
	err = dao.Layout.Ctx(ctx).
		FieldsEx(dao.Layout.Columns().Map).
		Where("location_id in (?)", locationIds).
		Scan(&layouts)
	if err != nil {
		return
	}
	return
}

// GetRuntimeLayoutMap 获取运行中的布局地图
func GetRuntimeLayoutMap(ctx context.Context, layoutId int) (cells []layout.Cell, err error) {
	var (
		mapGvar  = gvar.New(nil)
		redis    = g.Redis()
		cacheKey = cache_key.LayoutMapKey(layoutId)
	)
	mapGvar, err = redis.Get(ctx, cacheKey)
	if err != nil {
		err = utility.Err.NewSys(err)
		return
	}
	err = mapGvar.Scan(&cells)
	if err != nil {
		err = utility.Err.NewSys(err)
		return
	}
	if len(cells) != 0 {
		return
	}

	// 读库并载入缓存
	var expireAt = gtime.Now().EndOfDay().Time
	layoutData, err := GetOne(ctx, layoutId)
	if err != nil {
		return nil, err
	}
	cells, err = JsonToLayoutCells(ctx, layoutData.Map)
	if err != nil {
		return nil, err
	}
	if len(cells) == 0 {
		// 预防缓存穿透
		layoutData.Map = `[{"x":0,"y":0}]`
		expireAt = gtime.Now().
			Add(5 * time.Second).Time
	}

	_, err = redis.Set(ctx, cacheKey, layoutData.Map)
	if err != nil {
		return nil, err
	}
	_, err = redis.ExpireAt(ctx, cacheKey, expireAt)
	if err != nil {
		return nil, err
	}

	return
}
