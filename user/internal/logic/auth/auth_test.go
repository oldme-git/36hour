package auth

import (
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

		user2, err := service.Auth().GetUserInfo(ctx, token)
		t.AssertNil(err)
		t.Assert(user.Id, user2.Id)
		t.Assert(user.Username, user2.Username)

		err = service.User().Delete(ctx, id)
	})
}
