package policy

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"seat/internal/model/policy"
	"seat/internal/service"
)

type sPolicy struct {
}

func init() {
	service.RegisterPolicy(New())
}

func New() *sPolicy {
	return &sPolicy{}
}

// New 从 json 字符串创建一个 Policy 对象
func (s *sPolicy) New(ctx context.Context, str string) (p *policy.Policy, err error) {
	p = new(policy.Policy)
	if err := gjson.DecodeTo(str, p); err != nil {
		return nil, err
	}
	return p, nil
}
