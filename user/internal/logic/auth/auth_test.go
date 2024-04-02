package auth

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"user/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "user/internal/logic/snowflake"
)

var ctx = gctx.New()

func TestLogin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		token, err := service.Auth().Login(ctx)
		t.AssertNil(err)
		fmt.Println(token)
	})
}
