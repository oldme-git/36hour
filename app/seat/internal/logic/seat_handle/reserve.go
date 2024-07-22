package seat_handle

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/36hour/app/seat/internal/imodel"
	"github.com/oldme-git/36hour/app/seat/utility/cache"
	"github.com/oldme-git/36hour/utility"
)

// Reserve 预定座位
// 1. 检查座位是否已经被预定
// 2. 预定座位
// 3. 返回预定结果
func Reserve(ctx context.Context, userSeat *imodel.UserSeat) (err error) {
	// 检查座位是否已经被预定
	isExist, err := checkReserve(ctx, userSeat)
	if err != nil {
		return
	}
	if isExist {
		return utility.Err.New(3003)
	}
	err = writeCacheReserve(ctx, userSeat)
	if err != nil {
		return err
	}
	return
}

// checkReserve 检查座位是否已经被预定
func checkReserve(ctx context.Context, userSeat *imodel.UserSeat) (isReserve bool, err error) {
	var (
		redis     = g.Redis()
		key       = cache.SeatStatusKey(userSeat.LocationId)
		cellNoStr = strconv.Itoa(userSeat.CellNo)
	)

	isExist, err := redis.HExists(ctx, key, cellNoStr)
	if err != nil {
		return false, utility.Err.NewSys(err)
	}

	return isExist == 1, nil
}

// writeCacheReserve 将预约信息写入缓存
func writeCacheReserve(ctx context.Context, userSeat *imodel.UserSeat) (err error) {
	var (
		redis     = g.Redis()
		key       = cache.SeatStatusKey(userSeat.LocationId)
		cellNoStr = strconv.Itoa(userSeat.CellNo)
	)

	_, err = redis.HSetNX(ctx, key, cellNoStr, map[string]interface{}{
		"uid":    userSeat.Uid,
		"status": userSeat.Status,
	})
	if err != nil {
		return utility.Err.NewSys(err)
	}
	err = cache.ExpireAtEndOfDay(ctx, key)
	if err != nil {
		return utility.Err.NewSys(err)
	}

	return
}
