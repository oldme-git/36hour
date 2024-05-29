package auth_test

import (
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/user/internal/logic/auth"
	_ "github.com/oldme-git/36hour/app/user/internal/logic/snowflake"
	"github.com/oldme-git/36hour/app/user/internal/logic/user"
	_ "github.com/oldme-git/36hour/app/user/internal/logic/user"
	"github.com/oldme-git/36hour/app/user/internal/model"
)

var ctx = gctx.New()

func TestLogin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			plainPwd = "12345678"
			userOne  = &model.User{
				Username: "test",
				Password: plainPwd,
			}
		)

		id, err := user.Create(ctx, userOne)
		t.AssertNil(err)

		token, err := auth.Login(ctx, userOne.Username, plainPwd)
		t.AssertNil(err)

		user2, err := auth.GetUserInfo(ctx, token)
		t.AssertNil(err)
		t.Assert(userOne.Id, user2.Id)
		t.Assert(userOne.Username, user2.Username)

		err = user.Delete(ctx, id)
	})
}
