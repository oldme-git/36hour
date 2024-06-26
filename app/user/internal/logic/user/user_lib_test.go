package user_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/user/internal/logic/user"
	"github.com/oldme-git/36hour/app/user/internal/model"
)

func TestBindLib(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			userId     model.Id = 1
			libId      model.Id = 1
			BoundLibId model.Id
		)
		err := user.BindLib(ctx, userId, libId)
		t.AssertNil(err)

		BoundLibId, err = user.GetUserLib(ctx, userId)
		t.AssertNil(err)
		t.Assert(BoundLibId, libId)

		libId = 2
		err = user.BindLib(ctx, userId, libId)
		t.AssertNil(err)

		BoundLibId, err = user.GetUserLib(ctx, userId)
		t.AssertNil(err)
		t.Assert(BoundLibId, libId)
	})
}
