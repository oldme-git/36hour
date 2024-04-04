package hall

import (
	"database/sql"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"libManager/internal/model/entity"
	"libManager/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx    = gctx.New()
			hall   = new(entity.Hall)
			hallIn = &entity.Hall{
				LibId:    1,
				FloorId:  1,
				HallName: "hallTest",
			}
		)

		// Create
		id, err := service.Hall().Create(ctx, hallIn)
		t.AssertNil(err)

		// GetList
		halls, err := service.Hall().GetList(ctx, 1, 1)
		t.AssertNil(err)
		t.Assert(len(halls), 1)

		// GetOne
		hall, err = service.Hall().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(hall.Id, id)
		t.Assert(hall.LibId, hallIn.LibId)
		t.Assert(hall.FloorId, hallIn.FloorId)
		t.Assert(hall.HallName, hallIn.HallName)

		// Update
		var hallUptIn = &entity.Hall{
			Id:       id,
			LibId:    2,
			FloorId:  2,
			HallName: "hallTestUpt",
		}
		err = service.Hall().Update(ctx, hallUptIn)
		t.AssertNil(err)
		hall, err = service.Hall().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(hall.LibId, hallUptIn.LibId)
		t.Assert(hall.FloorId, hallUptIn.FloorId)
		t.Assert(hall.HallName, hallUptIn.HallName)

		// Delete
		err = service.Hall().Delete(ctx, id)
		t.AssertNil(err)
		_, err = service.Hall().GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
	})
}
