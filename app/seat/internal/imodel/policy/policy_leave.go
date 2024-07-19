package policy

// Leave 暂离模式
type Leave struct {
	LeaveOpen bool
	// 暂离配置
	LeaveConf []LeaveConf
}
type LeaveConf struct {
	// 暂离名称，如午休
	LeaveName string
	// 暂离时长，单位分钟
	LeaveTime int
	// 暂离次数
	LeaveCount int
	// 开启时间，例如 08:00
	LeaveStart string
	// 关闭时间，例如 18:00
	LeaveEnd string
}
