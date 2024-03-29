package snowflake

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"user/internal/service"
)

func TestGet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		id := service.Snowflake().Get()
		t.AssertNQ(id, 0)
	})
}
