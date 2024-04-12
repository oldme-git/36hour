package policy

// Policy 策略
type Policy struct {
	// 黑名单策略
	Black
	// 暂离策略
	Leave
	// 标记策略
	Mark
	// 续时策略
	Renew
	// 预约策略
	Reserve
	// 时间策略
	Time
}
