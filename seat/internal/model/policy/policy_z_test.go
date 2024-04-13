package policy

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gutil"
)

func TestNew(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			p = &Policy{
				Black: Black{
					BlackOpen: true,
				},
			}
		)
		s, err := p.String()
		t.AssertNil(err)
		ps, err := NewStr(s)
		t.AssertNil(err)

		gutil.Dump(p)
		gutil.Dump(ps)
		t.Assert(p, ps)
	})
}
