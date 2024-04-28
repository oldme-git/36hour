package user

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/user/internal/model"
	"github.com/oldme-git/36hour/app/user/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/oldme-git/36hour/app/user/internal/logic/snowflake"
)

var ctx = gctx.New()

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			testId model.Id = 1
			user            = new(model.User)
			userIn          = &model.User{
				Id:       testId,
				Username: "test",
				Password: "123456",
				Phone:    "12345678901",
			}
		)

		// Delete
		err := service.User().Delete(ctx, userIn.Id)
		t.AssertNil(err)

		// Create
		id, err := service.User().Create(ctx, userIn)
		t.AssertNil(err)
		t.Assert(id, userIn.Id)

		// GetList
		_, err = service.User().GetList(ctx, 1, 10)
		t.AssertNil(err)

		// GetOne
		user, err = service.User().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(user.Id, userIn.Id)
		t.Assert(user.Username, userIn.Username)
		t.Assert(user.Password, userIn.Password)
		t.Assert(user.Phone, userIn.Phone)

		// Update
		var userUpt = &model.User{
			Id:       testId,
			Username: "test2",
			Phone:    "12345678901",
		}
		err = service.User().Update(ctx, userUpt)
		t.AssertNil(err)
		user, err = service.User().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(user.Id, userUpt.Id)
		t.Assert(user.Username, userUpt.Username)
		t.Assert(user.Phone, userUpt.Phone)
	})
}
