// 生成测试数据文件
package test

import (
	"fmt"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/user/internal/logic/user"
	"github.com/oldme-git/36hour/app/user/internal/model"
)

// TestCreate 测试创建用户
func TestCreate(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			userIn = &model.User{
				Username: "oldme",
				Password: "12345678",
				Phone:    "17122334455",
			}
		)

		_, err := user.Create(gctx.New(), userIn)
		t.AssertNil(err)
	})
}

// TestCreateMulti 测试批量创建用户
func TestCreateMulti(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			num         = 10
			usernamePre = "oldme"
			phonePre    = "171"
			userLibId   = model.Id(1)
			userIn      = new(model.User)
		)
		for i := 1; i <= num; i++ {
			userIn = &model.User{
				Username: fmt.Sprintf("%s_%d", usernamePre, i),
				Password: "12345678",
				Phone:    fmt.Sprintf("%s%08d", phonePre, i),
			}
			userId, err := user.Create(gctx.New(), userIn)
			t.AssertNil(err)

			err = user.BindLib(gctx.New(), userId, userLibId)
			t.AssertNil(err)
		}
	})
}
