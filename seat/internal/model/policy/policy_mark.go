package policy

type MarkType int

const (
	// MarkTypeRelease 释放座位
	MarkTypeRelease MarkType = iota + 1
	// MarkTypeGet 获取座位
	MarkTypeGet
)

// Mark 标记模式，监督占座
type Mark struct {
	MarkOpen bool
	// 空座率大于多少时，允许标记
	MarkRate int
	// 标记类型
	MarkType MarkType
	// 允许回签的时间，单位分钟
	MarkBack int
	// 限制每个人标记的次数/天
	MarkLimit int
}
