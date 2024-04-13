package policy

// ReservePrepare 提前预约，意思是提前多少天可以预约
type ReservePrepare struct {
	ReservePrepareOpen bool
	// 允许多少天提前预约
	ReservePrepareDay int
	// 允许预约的比例
	ReservePrepareRate int
}

// Reserve 当天预约
type Reserve struct {
	// 闭馆前多少分钟不允许预约
	ReserveCloseMinute int
	// 多少分钟内只能预约一次
	ReserveMinute int
	// 闭馆前多少分钟通知
	ReserveNoticeMinute int
}
