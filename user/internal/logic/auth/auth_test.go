package auth

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"user/internal/model"
	"user/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "user/internal/logic/snowflake"
	_ "user/internal/logic/user"
)

var ctx = gctx.New()

func TestLogin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			plainPwd = "12345678"
			user     = &model.User{
				Username: "test",
				Password: plainPwd,
			}
		)

		id, err := service.User().Create(ctx, user)
		t.AssertNil(err)

		token, err := service.Auth().Login(ctx, user.Username, plainPwd)
		t.AssertNil(err)
		fmt.Println(token)

		err = service.User().Delete(ctx, id)
	})
}
