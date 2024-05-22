package lib_test

import (
	"database/sql"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
	"github.com/oldme-git/36hour/app/lib-manager/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/oldme-git/36hour/app/lib-manager/internal/logic/floor"
	_ "github.com/oldme-git/36hour/app/lib-manager/internal/logic/location"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx   = gctx.New()
			lib   = new(entity.Lib)
			libIn = &entity.Lib{
				LibName: "libTest",
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
			LibName:  "lib",
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
		t.Assert(lib.LibName, libIn.LibName)
		t.Assert(lib.Address, libIn.Address)
		t.Assert(lib.Active, libIn.Active)

		// Update
		var libUptIn = &entity.Lib{
			Id:      id,
			LibName: "libTestUpt",
			Address: "libTestAddressUpt",
			Active:  false,
		}
		err = service.Lib().Update(ctx, libUptIn)
		t.AssertNil(err)
		lib, err = service.Lib().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(lib.LibName, libUptIn.LibName)
		t.Assert(lib.Address, libUptIn.Address)
		t.Assert(lib.Active, libUptIn.Active)

		// Delete
		// 创建一些 floor 以便测试删除 lib 时的级联删除
		floorId, err := service.Floor().Create(ctx, &entity.Floor{
			LibId:     id,
			FloorName: "floorTest",
		})
		t.AssertNil(err)
		// 创建一些 location 以便测试删除 lib 时的级联删除
		_, err = service.Location().Create(ctx, &entity.Location{
			LibId:        id,
			FloorId:      floorId,
			LocationName: "locationTest",
		})
		t.AssertNil(err)
		_, err = service.Location().Create(ctx, &entity.Location{
			LibId:        id,
			FloorId:      floorId,
			LocationName: "locationTest2",
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
		// 获取 location 列表，验证 lib 删除时的级联删除
		_, err = service.Location().GetList(ctx, &dao.LocationSearchCondition{
			LibId: id,
		})
		t.Assert(err, sql.ErrNoRows)
	})
}

func TestExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ctx = gctx.New()

		// 创建一个 lib
		id, err := service.Lib().Create(ctx, &entity.Lib{
			LibName: "libTest",
			Address: "libTestAddress",
			Active:  true,
		})
		t.AssertNil(err)

		// 验证存在
		err = service.Lib().Exist(ctx, id)
		t.AssertNil(err)

		// 验证不存在
		err = service.Lib().Exist(ctx, id+10000)
		t.AssertNE(err, nil)

		// 删除
		err = service.Lib().Delete(ctx, id)
		t.AssertNil(err)
	})
}
