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
	_ "libManager/internal/logic/hall"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx     = gctx.New()
			floor   = new(entity.Floor)
			floorIn = &entity.Floor{
				LibId:     1,
				FloorName: "floorTest",
			}
		)

		// Create
		id, err := service.Floor().Create(ctx, floorIn)
		t.AssertNil(err)

		// GetList
		condition := &dao.FloorSearchCondition{
			Page:     1,
			PageSize: 1,
			LibId:    1,
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
			LibId:     2,
			FloorName: "floorTestUpt",
		}
		err = service.Floor().Update(ctx, floorUptIn)
		t.AssertNil(err)
		floor, err = service.Floor().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(floor.LibId, floorUptIn.LibId)
		t.Assert(floor.FloorName, floorUptIn.FloorName)

		// Delete
		// 创建一些 hall 以便测试删除 floor 时的级联删除
		_, err = service.Hall().Create(ctx, &entity.Hall{
			LibId:    2,
			FloorId:  id,
			HallName: "hallTest",
		})
		t.AssertNil(err)
		_, err = service.Hall().Create(ctx, &entity.Hall{
			LibId:    2,
			FloorId:  id,
			HallName: "hallTest2",
		})
		t.AssertNil(err)

		err = service.Floor().Delete(ctx, id)
		t.AssertNil(err)
		_, err = service.Floor().GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
		// 获取 hall 列表，验证 floor 删除时的级联删除
		_, err = service.Hall().GetList(ctx, &dao.HallSearchCondition{
			FloorId: id,
		})
		t.Assert(err, sql.ErrNoRows)
	})
}
