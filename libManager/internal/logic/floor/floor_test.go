package floor

import (
	"database/sql"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"libManager/internal/dao"
	"libManager/internal/model/entity"
	"libManager/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "libManager/internal/logic/lib"
	_ "libManager/internal/logic/location"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx     = gctx.New()
			floor   = new(entity.Floor)
			floorIn = &entity.Floor{
				FloorName: "floorTest",
			}
		)

		// 创建一个 lib
		libId, err := service.Lib().Create(ctx, &entity.Lib{
			LibName: "libTest",
			Address: "libTestAddress",
			Active:  true,
		})
		floorIn.LibId = libId
		defer service.Lib().Delete(ctx, libId)

		// Create
		id, err := service.Floor().Create(ctx, floorIn)
		t.AssertNil(err)

		// GetList
		condition := &dao.FloorSearchCondition{
			Page:     1,
			PageSize: 1,
			LibId:    libId,
		}
		floors, err := service.Floor().GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(floors), 1)

		// GetOne
		floor, err = service.Floor().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(floor.Id, id)
		t.Assert(floor.LibId, floorIn.LibId)
		t.Assert(floor.FloorName, floorIn.FloorName)

		// Update
		var floorUptIn = &entity.Floor{
			Id:        id,
			LibId:     libId,
			FloorName: "floorTestUpt",
		}
		err = service.Floor().Update(ctx, floorUptIn)
		t.AssertNil(err)
		floor, err = service.Floor().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(floor.LibId, floorUptIn.LibId)
		t.Assert(floor.FloorName, floorUptIn.FloorName)

		// Delete
		// 创建一些 location 以便测试删除 floor 时的级联删除
		_, err = service.Location().Create(ctx, &entity.Location{
			LibId:        libId,
			FloorId:      id,
			LocationName: "locationTest",
		})
		t.AssertNil(err)
		_, err = service.Location().Create(ctx, &entity.Location{
			LibId:        libId,
			FloorId:      id,
			LocationName: "locationTest2",
		})
		t.AssertNil(err)

		err = service.Floor().Delete(ctx, id)
		t.AssertNil(err)
		_, err = service.Floor().GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
		// 获取 location 列表，验证 floor 删除时的级联删除
		_, err = service.Location().GetList(ctx, &dao.LocationSearchCondition{
			FloorId: id,
		})
		t.Assert(err, sql.ErrNoRows)
	})
}

func TestExist(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ctx = gctx.New()

		// 创建一个 lib
		libId, err := service.Lib().Create(ctx, &entity.Lib{
			LibName: "libTest",
			Address: "libTestAddress",
			Active:  true,
		})
		t.AssertNil(err)

		// 创建一个 floor
		id, err := service.Floor().Create(ctx, &entity.Floor{
			LibId:     libId,
			FloorName: "floorTest",
		})
		t.AssertNil(err)

		// 验证存在
		err = service.Floor().Exist(ctx, id)
		t.AssertNil(err)

		// 验证不存在
		err = service.Floor().Exist(ctx, id+10000)
		t.AssertNE(err, nil)

		// 删除
		err = service.Lib().Delete(ctx, libId)
		t.AssertNil(err)
	})
}
