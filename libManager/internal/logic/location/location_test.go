package location

import (
	"database/sql"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"libManager/internal/dao"
	"libManager/internal/model/entity"
	"libManager/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
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

		// Create
		id, err := service.Location().Create(ctx, locationIn)
		t.AssertNil(err)

		// GetList
		condition := &dao.LocationSearchCondition{
			Page:     1,
			PageSize: 1,
			LibId:    1,
			FloorId:  1,
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
			LibId:        2,
			FloorId:      2,
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
