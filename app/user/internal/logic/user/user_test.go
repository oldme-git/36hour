package user_test

import (
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	_ "github.com/oldme-git/36hour/app/user/internal/logic/snowflake"
	"github.com/oldme-git/36hour/app/user/internal/logic/user"
	"github.com/oldme-git/36hour/app/user/internal/model"
)

var ctx = gctx.New()

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			testId  model.Id = 1
			userOne          = new(model.User)
			userIn           = &model.User{
				Id:       testId,
				Username: "test",
				Password: "123456",
				Phone:    "12345678901",
			}
		)

		// Delete
		err := user.Delete(ctx, userIn.Id)
		t.AssertNil(err)

		// Create
		id, err := user.Create(ctx, userIn)
		t.AssertNil(err)
		t.Assert(id, userIn.Id)

		// GetList
		_, err = user.GetList(ctx, 1, 10)
		t.AssertNil(err)

		// GetOne
		userOne, err = user.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(userOne.Id, userIn.Id)
		t.Assert(userOne.Username, userIn.Username)
		t.Assert(userOne.Password, userIn.Password)
		t.Assert(userOne.Phone, userIn.Phone)

		// Update
		var userUpt = &model.User{
			Id:       testId,
			Username: "test2",
			Phone:    "12345678901",
		}
		err = user.Update(ctx, userUpt)
		t.AssertNil(err)
		userOne, err = user.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(userOne.Id, userUpt.Id)
		t.Assert(userOne.Username, userUpt.Username)
		t.Assert(userOne.Phone, userUpt.Phone)
	})
}
