package snowflake_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/user/internal/service"
)

func TestGet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		id := service.Snowflake().Get()
		t.AssertNE(id, 0)
	})
}
