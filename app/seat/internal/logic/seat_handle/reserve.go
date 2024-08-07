package seat_handle

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/layout"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/seat"
	"github.com/oldme-git/36hour/app/seat/utility/cache"
	"github.com/oldme-git/36hour/utility"
)

// Reserve 预定座位
// 1. 检查座位是否已经被预定
// 2. 预定座位
// 3. 返回预定结果
func Reserve(ctx context.Context, userSeat *seat.UserSeat) (err error) {
	userSeat.Status = GetCellStatus(ctx, userSeat.Type)
	isFree, err := seatIsFree(ctx, userSeat)
	if err != nil {
		return
	}
	if !isFree {
		return utility.Err.New(3003)
	}
	is, err := reserveLock(ctx, userSeat)
	if err != nil {
		return
	}
	if !is {
		return utility.Err.New(3003)
	}
	defer reserveUnlock(ctx, userSeat)

	err = writeCacheReserve(ctx, userSeat)
	if err != nil {
		return err
	}

	return
}

// reserveLock 预约锁
// 只有拿到锁的用户才能进行预约
func reserveLock(ctx context.Context, userSeat *seat.UserSeat) (is bool, err error) {
	var (
		redis = g.Redis()
		key   = cache.ReserveLock(userSeat.LocationId, userSeat.CellNo)
	)

	is, err = redis.SetNX(ctx, key, int64(userSeat.Uid))
	if err != nil {
		return false, utility.Err.NewSys(err)
	}
	if !is {
		return false, nil
	}
	_, err = redis.Expire(ctx, key, 5)
	if err != nil {
		return false, utility.Err.NewSys(err)
	}
	return is, nil
}

// reserveUnlock 预约释放锁
func reserveUnlock(ctx context.Context, userSeat *seat.UserSeat) (err error) {
	var (
		redis = g.Redis()
		key   = cache.ReserveLock(userSeat.LocationId, userSeat.CellNo)
	)

	_, err = redis.Del(ctx, key)
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return nil
}

// seatIsFree 检查座位是否是空闲状态
func seatIsFree(ctx context.Context, userSeat *seat.UserSeat) (isFree bool, err error) {
	var (
		redis     = g.Redis()
		key       = cache.LayoutMapKey(userSeat.LocationId)
		cellNoStr = strconv.Itoa(userSeat.CellNo)
		cell      = new(layout.Cell)
	)

	data, err := redis.HGet(ctx, key, cellNoStr)
	if err != nil {
		return false, utility.Err.NewSys(err)
	}
	if err := data.Scan(&cell); err != nil {
		return false, utility.Err.NewSys(err)
	}

	return cell.Status == layout.CellStatusFree, nil
}

// writeCacheReserve 将预约信息写入缓存
func writeCacheReserve(ctx context.Context, userSeat *seat.UserSeat) (err error) {
	var (
		redis     = g.Redis()
		key       = cache.LayoutSeatStatusKey(userSeat.LocationId)
		cellNoStr = strconv.Itoa(userSeat.CellNo)
	)

	_, err = redis.HSet(ctx, key, map[string]interface{}{
		cellNoStr: map[string]interface{}{
			"uid":    userSeat.Uid,
			"type":   userSeat.Type,
			"status": userSeat.Status,
		},
	})
	if err != nil {
		return utility.Err.NewSys(err)
	}
	err = cache.ExpireAtEndOfDay(ctx, key)
	if err != nil {
		return utility.Err.NewSys(err)
	}
	// 更新座位状态
	err = updateCellStatus(ctx, userSeat.LocationId, userSeat.CellNo, userSeat.Status)
	if err != nil {
		return
	}
	// 写入日志到数据库
	err = writeLog(ctx, 0, userSeat)
	if err != nil {
		return
	}

	return
}
