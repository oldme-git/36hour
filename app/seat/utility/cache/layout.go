package cache

import "fmt"

const (
	// layoutMap 运行中的布局座位图
	layoutMap = "layout_runtime_map:%d"

	// layoutSeatStatus 座位使用情况
	layoutSeatStatus = "layout_seat_status:%d"

	// reserveLock 预约锁
	reserveLock = "reserve_lock:%d:%d"
)

// LayoutMapKey 运行中的布局座位图
func LayoutMapKey(layoutId int) string {
	return fmt.Sprintf(layoutMap, layoutId)
}

// LayoutSeatStatusKey 布局座位使用情况
func LayoutSeatStatusKey(locationId int) string {
	return fmt.Sprintf(layoutSeatStatus, locationId)
}

// ReserveLock 预约锁
func ReserveLock(locationId, seatNo int) string {
	return fmt.Sprintf(reserveLock, locationId, seatNo)
}
