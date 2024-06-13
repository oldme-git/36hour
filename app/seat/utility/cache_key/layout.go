package cache_key

import "fmt"

const (
	// LayoutMap 运行中的布局座位图
	LayoutMap = "layout_map:%d"
)

// LayoutMapKey 运行中的布局座位图
func LayoutMapKey(layoutId int) string {
	return fmt.Sprintf(LayoutMap, layoutId)
}
