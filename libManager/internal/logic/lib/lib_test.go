package lib

import (
	"database/sql"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"libManager/internal/dao"
	"libManager/internal/model/entity"
	"libManager/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "libManager/internal/logic/floor"
	_ "libManager/internal/logic/hall"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx   = gctx.New()
			lib   = new(entity.Lib)
			libIn = &entity.Lib{
				Name:    "libTest",
				Address: "libTestAddress",
				Active:  true,
			}
		)

		// Create
		id, err := service.Lib().Create(ctx, libIn)
		t.AssertNil(err)

		// GetList
		condition := &dao.LibSearchCondition{
			Page:     1,
			PageSize: 1,
			Name:     "lib",
			Address:  "Add",
			Active:   true,
		}
		libs, err := service.Lib().GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(libs), 1)

		// GetOne
		lib, err = service.Lib().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(lib.Id, id)
		t.Assert(lib.Name, libIn.Name)
		t.Assert(lib.Address, libIn.Address)
		t.Assert(lib.Active, libIn.Active)

		// Update
		var libUptIn = &entity.Lib{
			Id:      id,
			Name:    "libTestUpt",
			Address: "libTestAddressUpt",
			Active:  false,
		}
		err = service.Lib().Update(ctx, libUptIn)
		t.AssertNil(err)
		lib, err = service.Lib().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(lib.Name, libUptIn.Name)
		t.Assert(lib.Address, libUptIn.Address)
		t.Assert(lib.Active, libUptIn.Active)

		// Delete
		// 创建一些 floor 以便测试删除 lib 时的级联删除
		floorId, err := service.Floor().Create(ctx, &entity.Floor{
			LibId:     id,
			FloorName: "floorTest",
		})
		t.AssertNil(err)
		// 创建一些 hall 以便测试删除 lib 时的级联删除
		_, err = service.Hall().Create(ctx, &entity.Hall{
			LibId:    id,
			FloorId:  floorId,
			HallName: "hallTest",
		})
		t.AssertNil(err)
		_, err = service.Hall().Create(ctx, &entity.Hall{
			LibId:    id,
			FloorId:  floorId,
			HallName: "hallTest2",
		})

		err = service.Lib().Delete(ctx, id)
		t.AssertNil(err)
		_, err = service.Lib().GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
		// 获取 floor 列表，验证 lib 删除时的级联删除
		_, err = service.Floor().GetList(ctx, &dao.FloorSearchCondition{
			LibId: id,
		})
		t.Assert(err, sql.ErrNoRows)
		// 获取 hall 列表，验证 lib 删除时的级联删除
		_, err = service.Hall().GetList(ctx, &dao.HallSearchCondition{
			LibId: id,
		})
		t.Assert(err, sql.ErrNoRows)
	})
}
