package policy

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"seat/internal/model/policy"
	"seat/internal/service"
)

func TestNew(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.New()
			p   = &policy.Policy{
				Black: policy.Black{
					BlackOpen: true,
				},
			}
		)
		s, err := p.String()
		t.AssertNil(err)
		ps, err := service.Policy().New(ctx, s)
		t.AssertNil(err)
		t.Assert(p, ps)
	})
}
