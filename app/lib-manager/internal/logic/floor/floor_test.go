package floor_test

import (
	"database/sql"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/floor"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/lib"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/location"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/oldme-git/36hour/app/lib-manager/internal/logic/lib"
	_ "github.com/oldme-git/36hour/app/lib-manager/internal/logic/location"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx       = gctx.New()
			floorData = new(entity.Floor)
			floorIn   = &entity.Floor{
				FloorName: "floorTest",
			}
		)

		// 创建一个 lib
		libId, err := lib.Create(ctx, &entity.Lib{
			LibName: "libTest",
			Address: "libTestAddress",
			Active:  true,
		})
		floorIn.LibId = libId
		defer lib.Delete(ctx, libId)

		// Create
		id, err := floor.Create(ctx, floorIn)
		t.AssertNil(err)

		// GetList
		condition := &dao.FloorSearchCondition{
			Page:     1,
			PageSize: 1,
			LibId:    libId,
		}
		floors, err := floor.GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(floors), 1)

		// GetOne
		floorData, err = floor.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(floorData.Id, id)
		t.Assert(floorData.LibId, floorIn.LibId)
		t.Assert(floorData.FloorName, floorIn.FloorName)

		// Update
		var floorUptIn = &entity.Floor{
			Id:        id,
			LibId:     libId,
			FloorName: "floorTestUpt",
		}
		err = floor.Update(ctx, floorUptIn)
		t.AssertNil(err)
		floorData, err = floor.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(floorData.LibId, floorUptIn.LibId)
		t.Assert(floorData.FloorName, floorUptIn.FloorName)

		// Delete
		// 创建一些 location 以便测试删除 floor 时的级联删除
		_, err = location.Create(ctx, &entity.Location{
			LibId:        libId,
			FloorId:      id,
			LocationName: "locationTest",
		})
		t.AssertNil(err)
		_, err = location.Create(ctx, &entity.Location{
			LibId:        libId,
			FloorId:      id,
			LocationName: "locationTest2",
		})
		t.AssertNil(err)

		err = floor.Delete(ctx, id)
		t.AssertNil(err)
		_, err = floor.GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
		// 获取 location 列表，验证 floor 删除时的级联删除
		_, err = location.GetList(ctx, &dao.LocationSearchCondition{
			FloorId: id,
		})
		t.Assert(err, sql.ErrNoRows)
	})
}

func TestExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ctx = gctx.New()

		// 创建一个 lib
		libId, err := lib.Create(ctx, &entity.Lib{
			LibName: "libTest",
			Address: "libTestAddress",
			Active:  true,
		})
		t.AssertNil(err)

		// 创建一个 floor
		id, err := floor.Create(ctx, &entity.Floor{
			LibId:     libId,
			FloorName: "floorTest",
		})
		t.AssertNil(err)

		// 验证存在
		err = floor.Exist(ctx, id)
		t.AssertNil(err)

		// 验证不存在
		err = floor.Exist(ctx, id+10000)
		t.AssertNE(err, nil)

		// 删除
		err = lib.Delete(ctx, libId)
		t.AssertNil(err)
	})
}
