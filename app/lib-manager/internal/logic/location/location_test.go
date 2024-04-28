package location

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
	_ "github.com/oldme-git/36hour/app/lib-manager/internal/logic/lib"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx        = gctx.New()
			location   = new(entity.Location)
			locationIn = &entity.Location{
				LibId:        1,
				FloorId:      1,
				LocationName: "locationTest",
			}
		)

		// 创建一个 lib
		libId, err := service.Lib().Create(ctx, &entity.Lib{
			LibName: "libTest",
			Address: "libTestAddress",
			Active:  true,
		})
		t.AssertNil(err)
		locationIn.LibId = libId
		defer service.Lib().Delete(ctx, libId)

		// 创建一个 floor
		floorId, err := service.Floor().Create(ctx, &entity.Floor{
			LibId:     libId,
			FloorName: "floorTest",
		})
		t.AssertNil(err)
		locationIn.FloorId = floorId

		// Create
		id, err := service.Location().Create(ctx, locationIn)
		t.AssertNil(err)

		// GetList
		condition := &dao.LocationSearchCondition{
			Page:     1,
			PageSize: 1,
			LibId:    libId,
			FloorId:  floorId,
		}
		locations, err := service.Location().GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(locations), 1)

		// GetOne
		location, err = service.Location().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(location.Id, id)
		t.Assert(location.LibId, locationIn.LibId)
		t.Assert(location.FloorId, locationIn.FloorId)
		t.Assert(location.LocationName, locationIn.LocationName)

		// Update
		var locationUptIn = &entity.Location{
			Id:           id,
			LibId:        libId,
			FloorId:      floorId,
			LocationName: "locationTestUpt",
		}
		err = service.Location().Update(ctx, locationUptIn)
		t.AssertNil(err)
		location, err = service.Location().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(location.LibId, locationUptIn.LibId)
		t.Assert(location.FloorId, locationUptIn.FloorId)
		t.Assert(location.LocationName, locationUptIn.LocationName)

		// Delete
		err = service.Location().Delete(ctx, id)
		t.AssertNil(err)
		_, err = service.Location().GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
	})
}
