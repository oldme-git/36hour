package policy

import "github.com/gogf/gf/v2/encoding/gjson"

// New 从 json 字符串创建一个 Policy 对象
func New(s string) (*Policy, error) {
	p := new(Policy)
	if err := gjson.DecodeTo(s, p); err != nil {
		return nil, err
	}
	return p, nil
}

// Policy 策略
type Policy struct {
	// 黑名单策略
	Black Black
	// 暂离策略
	Leave Leave
	// 标记策略
	Mark Mark
	// 续时策略
	Renew Renew
	// 预约策略
	Reserve Reserve
	// 时间策略
	Time Time
}

// String 转换为 json 字符串
func (p *Policy) String() (string, error) {
	return gjson.EncodeString(p)
}
