package policy

// Renew 续时模式
type Renew struct {
	RenewOpen bool
	// 续时周期，单位分钟
	RenewCycle int
}
