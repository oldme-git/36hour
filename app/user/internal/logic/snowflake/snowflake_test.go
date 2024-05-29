package snowflake_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/user/internal/logic/snowflake"
)

func TestGet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		id := snowflake.Get()
		t.AssertNE(id, 0)
	})
}
