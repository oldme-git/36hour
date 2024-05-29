package location_test

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
	_ "github.com/oldme-git/36hour/app/lib-manager/internal/logic/lib"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/location"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx          = gctx.New()
			locationData = new(entity.Location)
			locationIn   = &entity.Location{
				LibId:        1,
				FloorId:      1,
				LocationName: "locationTest",
			}
		)

		// 创建一个 lib
		libId, err := lib.Create(ctx, &entity.Lib{
			LibName: "libTest",
			Address: "libTestAddress",
			Active:  true,
		})
		t.AssertNil(err)
		locationIn.LibId = libId
		defer lib.Delete(ctx, libId)

		// 创建一个 floor
		floorId, err := floor.Create(ctx, &entity.Floor{
			LibId:     libId,
			FloorName: "floorTest",
		})
		t.AssertNil(err)
		locationIn.FloorId = floorId

		// Create
		id, err := location.Create(ctx, locationIn)
		t.AssertNil(err)

		// GetList
		condition := &dao.LocationSearchCondition{
			Page:     1,
			PageSize: 1,
			LibId:    libId,
			FloorId:  floorId,
		}
		locations, err := location.GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(locations), 1)

		// GetOne
		locationData, err = location.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(locationData.Id, id)
		t.Assert(locationData.LibId, locationIn.LibId)
		t.Assert(locationData.FloorId, locationIn.FloorId)
		t.Assert(locationData.LocationName, locationIn.LocationName)

		// Update
		var locationUptIn = &entity.Location{
			Id:           id,
			LibId:        libId,
			FloorId:      floorId,
			LocationName: "locationTestUpt",
		}
		err = location.Update(ctx, locationUptIn)
		t.AssertNil(err)
		locationData, err = location.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(locationData.LibId, locationUptIn.LibId)
		t.Assert(locationData.FloorId, locationUptIn.FloorId)
		t.Assert(locationData.LocationName, locationUptIn.LocationName)

		// Delete
		err = location.Delete(ctx, id)
		t.AssertNil(err)
		_, err = location.GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
	})
}
