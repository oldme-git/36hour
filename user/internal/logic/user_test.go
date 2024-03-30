package logic

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"user/internal/model"
	"user/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

var ctx = gctx.New()

func TestCre(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		user := &model.User{
			Username: "test",
			Password: "123456",
			Phone:    "12345678901",
		}
		id, err := service.User().Cre(ctx, user)
		t.AssertNil(err)
		t.AssertNE(id, 0)
	})
}
