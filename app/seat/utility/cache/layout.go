package cache

import "fmt"

const (
	// LayoutMap 运行中的布局座位图
	LayoutMap = "layout_runtime_map:%d"

	// SeatStatus 座位使用情况
	SeatStatus = "seat_status:%d"
)

// LayoutMapKey 运行中的布局座位图
func LayoutMapKey(layoutId int) string {
	return fmt.Sprintf(LayoutMap, layoutId)
}

// SeatStatusKey 座位使用情况
func SeatStatusKey(locationId uint) string {
	return fmt.Sprintf(SeatStatus, locationId)
}
