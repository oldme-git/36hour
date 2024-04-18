package layout

import (
	"database/sql"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"seat/internal/dao"
	"seat/internal/model/entity"
	"seat/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "seat/internal/logic/policy_layout"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx      = gctx.New()
			layout   = new(entity.Layout)
			layoutIn = &entity.Layout{
				LocationId: 1,
				PolicyCId:  1,
				PolicyLId:  1,
				LayoutName: "layoutTest",
				Map:        `{"x":1,"y":1}`,
				Status:     1,
				Sort:       1,
				Seats:      1,
			}
		)

		// Create
		id, err := service.Layout().Create(ctx, layoutIn)
		t.AssertNil(err)

		// GetList
		condition := &dao.LayoutSearchCondition{
			Page:     1,
			PageSize: 1,
		}
		layouts, err := service.Layout().GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(layouts), 1)

		// GetOne
		layout, err = service.Layout().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(layout.Id, id)
		t.Assert(layout.LocationId, layoutIn.LocationId)
		t.Assert(layout.PolicyCId, layoutIn.PolicyCId)
		t.Assert(layout.PolicyLId, layoutIn.PolicyLId)
		t.Assert(layout.LayoutName, layoutIn.LayoutName)
		t.Assert(layout.Map, layoutIn.Map)
		t.Assert(layout.Status, layoutIn.Status)
		t.Assert(layout.Sort, layoutIn.Sort)
		t.Assert(layout.Seats, layoutIn.Seats)

		// Update
		var layoutUptIn = &entity.Layout{
			Id:         id,
			LocationId: 2,
			PolicyCId:  2,
			PolicyLId:  2,
			LayoutName: "layoutTestUpt",
			Map:        `{"x":2,"y":2}`,
			Status:     2,
			Sort:       2,
			Seats:      2,
		}
		err = service.Layout().Update(ctx, layoutUptIn)
		t.AssertNil(err)
		layout, err = service.Layout().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(layout.LocationId, layoutUptIn.LocationId)
		t.Assert(layout.PolicyCId, layoutUptIn.PolicyCId)
		t.Assert(layout.PolicyLId, layoutUptIn.PolicyLId)
		t.Assert(layout.LayoutName, layoutUptIn.LayoutName)
		t.Assert(layout.Map, layoutUptIn.Map)
		t.Assert(layout.Status, layoutUptIn.Status)
		t.Assert(layout.Sort, layoutUptIn.Sort)
		t.Assert(layout.Seats, layoutUptIn.Seats)

		// Delete
		err = service.Layout().Delete(ctx, id)
		t.AssertNil(err)
		_, err = service.Layout().GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
		_, err = service.PolicyLayout().GetOne(ctx, layoutUptIn.PolicyLId)
		t.Assert(err, sql.ErrNoRows)
	})
}
