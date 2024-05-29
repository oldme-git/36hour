package lib_test

import (
	"database/sql"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/floor"
	_ "github.com/oldme-git/36hour/app/lib-manager/internal/logic/floor"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/lib"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/location"
	_ "github.com/oldme-git/36hour/app/lib-manager/internal/logic/location"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx     = gctx.New()
			libData = new(entity.Lib)
			libIn   = &entity.Lib{
				LibName: "libTest",
				Address: "libTestAddress",
				Active:  true,
			}
		)

		// Create
		id, err := lib.Create(ctx, libIn)
		t.AssertNil(err)

		// GetList
		condition := &dao.LibSearchCondition{
			Page:     1,
			PageSize: 1,
			LibName:  "lib",
			Address:  "Add",
			Active:   true,
		}
		libs, err := lib.GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(libs), 1)

		// GetOne
		libData, err = lib.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(libData.Id, id)
		t.Assert(libData.LibName, libIn.LibName)
		t.Assert(libData.Address, libIn.Address)
		t.Assert(libData.Active, libIn.Active)

		// Update
		var libUptIn = &entity.Lib{
			Id:      id,
			LibName: "libTestUpt",
			Address: "libTestAddressUpt",
			Active:  false,
		}
		err = lib.Update(ctx, libUptIn)
		t.AssertNil(err)
		libData, err = lib.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(libData.LibName, libUptIn.LibName)
		t.Assert(libData.Address, libUptIn.Address)
		t.Assert(libData.Active, libUptIn.Active)

		// Delete
		// 创建一些 floor 以便测试删除 lib 时的级联删除
		floorId, err := floor.Create(ctx, &entity.Floor{
			LibId:     id,
			FloorName: "floorTest",
		})
		t.AssertNil(err)
		// 创建一些 location 以便测试删除 lib 时的级联删除
		_, err = location.Create(ctx, &entity.Location{
			LibId:        id,
			FloorId:      floorId,
			LocationName: "locationTest",
		})
		t.AssertNil(err)
		_, err = location.Create(ctx, &entity.Location{
			LibId:        id,
			FloorId:      floorId,
			LocationName: "locationTest2",
		})

		err = lib.Delete(ctx, id)
		t.AssertNil(err)
		_, err = lib.GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
		// 获取 floor 列表，验证 lib 删除时的级联删除
		_, err = floor.GetList(ctx, &dao.FloorSearchCondition{
			LibId: id,
		})
		t.Assert(err, sql.ErrNoRows)
		// 获取 location 列表，验证 lib 删除时的级联删除
		_, err = location.GetList(ctx, &dao.LocationSearchCondition{
			LibId: id,
		})
		t.Assert(err, sql.ErrNoRows)
	})
}

func TestExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ctx = gctx.New()

		// 创建一个 lib
		id, err := lib.Create(ctx, &entity.Lib{
			LibName: "libTest",
			Address: "libTestAddress",
			Active:  true,
		})
		t.AssertNil(err)

		// 验证存在
		err = lib.Exist(ctx, id)
		t.AssertNil(err)

		// 验证不存在
		err = lib.Exist(ctx, id+10000)
		t.AssertNE(err, nil)

		// 删除
		err = lib.Delete(ctx, id)
		t.AssertNil(err)
	})
}
