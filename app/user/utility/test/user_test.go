// 生成测试数据文件
package test

import (
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/user/internal/logic/user"
	"github.com/oldme-git/36hour/app/user/internal/model"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			testId model.Id = 1
			userIn          = &model.User{
				Id:       testId,
				Username: "oldme",
				Password: "12345678",
				Phone:    "17122334455",
			}
		)

		// Create
		id, err := user.Create(gctx.New(), userIn)
		t.AssertNil(err)
		t.Assert(id, userIn.Id)
	})
}
