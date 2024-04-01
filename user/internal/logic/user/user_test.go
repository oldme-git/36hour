package user

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"user/internal/model"
	"user/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "user/internal/logic/snowflake"
)

var ctx = gctx.New()

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			user   = new(model.User)
			userIn = &model.User{
				Id:       1,
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
		users, err := service.User().GetList(ctx, 1, 10)
		t.AssertNil(err)
		t.Assert(users[0].Id, userIn.Id)

		// GetOne
		user, err = service.User().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(user.Id, userIn.Id)
		t.Assert(user.Username, userIn.Username)
		t.Assert(user.Password, userIn.Password)
		t.Assert(user.Phone, userIn.Phone)

		// Update
		var userUpt = &model.User{
			Id:       1,
			Username: "test2",
			Password: "123456",
			Phone:    "12345678901",
		}
		err = service.User().Update(ctx, userUpt)
		t.AssertNil(err)
		user, err = service.User().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(user.Id, userUpt.Id)
		t.Assert(user.Username, userUpt.Username)
		t.Assert(user.Password, userUpt.Password)
		t.Assert(user.Phone, userUpt.Phone)
	})
}
