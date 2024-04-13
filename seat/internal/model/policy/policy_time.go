package policy

// Week 按week为周期管理开闭关时间
type Week int

const (
	Sun Week = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
)

// Time 时间策略
type Time struct {
	// 时间策略
	PolicyTimeConf map[Week]TimeConf
	// 提前预约时间，以开馆时间做为标准
	ReservePrepareMinute int
	// 保留座位的时间
	ReserveKeepMinute int
	// 节假日策略
	Holiday []Holiday
}

// TimeConf 时间策略配置
type TimeConf struct {
	// 开馆时间，例如：08:00
	OpenTime string
	// 闭馆时间，例如：18:00
	CloseTime string
}

// Holiday 节假日策略
type Holiday struct {
	// 开始时间
	StartDate string
	// 结束时间
	EndDate string
}
