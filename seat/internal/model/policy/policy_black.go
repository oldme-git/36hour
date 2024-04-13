package policy

// Black 黑名单
type Black struct {
	BlackOpen bool
	// 多少天内不允许预约
	BlackDay int
	// 多少次违约后加入黑名单
	BlackCount int
	// 清零周期，单位天
	BlackCycle int
}
