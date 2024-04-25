package layout

import (
	"database/sql"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"seat/internal/model"
	"seat/internal/model/entity"
	"seat/internal/model/layout"
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
		condition := &model.LayoutSearchCondition{
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

func TestJsonToLayoutCells(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx     = gctx.New()
			jsonStr = `[
				{"x":1,"y":1,"vx":3,"vy":3,"n":1,"l":"1号座位","t":1},
				{"x":5,"y":5,"vx":0,"vy":0,"n":2,"l":"2号座位","t":1}
			]`
		)
		cells, err := service.Layout().JsonToLayoutCells(ctx, jsonStr)
		t.AssertNil(err)
		t.Assert(len(cells), 2)

		t.Assert(cells[0].X, 1)
		t.Assert(cells[0].Y, 1)
		t.Assert(cells[0].VectorX, 3)
		t.Assert(cells[0].VectorY, 3)
		t.Assert(cells[0].No, 1)
		t.Assert(cells[0].Label, "1号座位")
		t.Assert(cells[0].Type, 1)

		t.Assert(cells[1].X, 5)
		t.Assert(cells[1].Y, 5)
		t.Assert(cells[1].VectorX, 0)
		t.Assert(cells[1].VectorY, 0)
		t.Assert(cells[1].No, 2)
		t.Assert(cells[1].Label, "2号座位")
		t.Assert(cells[1].Type, 1)
	})
}

func TestLayoutCellsToJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx   = gctx.New()
			cells = []layout.Cell{
				{X: 1, Y: 1, VectorX: 3, VectorY: 3, No: 1, Label: "1号座位", Type: layout.CellSeat},
				{X: 5, Y: 5, VectorX: 0, VectorY: 0, No: 2, Label: "2号座位", Type: layout.CellSeat},
			}
		)
		jsonStr, err := service.Layout().LayoutCellsToJson(ctx, cells)
		t.AssertNil(err)
		t.Assert(jsonStr, `[{"x":1,"y":1,"vx":3,"vy":3,"n":1,"l":"1号座位","t":1},{"x":5,"y":5,"vx":0,"vy":0,"n":2,"l":"2号座位","t":1}]`)
	})
}

func TestCalculateSeatsByJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx     = gctx.New()
			jsonStr = `[
				{"x":1,"y":1,"vx":3,"vy":3,"n":1,"l":"1号座位","t":1},
				{"x":5,"y":5,"vx":0,"vy":0,"n":2,"l":"2号座位","t":1},
				{"x":6,"y":6,"vx":0,"vy":0,"n":0,"l":"hello","t":2}
			]`
		)
		seats, err := service.Layout().CalculateSeatsByJson(ctx, jsonStr)
		t.AssertNil(err)
		t.Assert(seats, 2)
	})
}
