package layout_test

import (
	"database/sql"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/seat/internal/imodel"
	layoutModel "github.com/oldme-git/36hour/app/seat/internal/imodel/layout"
	"github.com/oldme-git/36hour/app/seat/internal/logic/layout"
	"github.com/oldme-git/36hour/app/seat/internal/logic/policy_layout"
	_ "github.com/oldme-git/36hour/app/seat/internal/logic/policy_layout"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx        = gctx.New()
			layoutData = new(entity.Layout)
			layoutIn   = &entity.Layout{
				LocationId: 1,
				PolicyCId:  1,
				PolicyLId:  1,
				LayoutName: "layoutTest",
				Map:        `[{"x":1,"y":1}]`,
				Status:     1,
				Sort:       1,
				Seats:      1,
			}
		)

		// Create
		id, err := layout.Create(ctx, layoutIn)
		t.AssertNil(err)

		// GetList
		condition := &imodel.LayoutSearchCondition{
			Page:     1,
			PageSize: 1,
		}
		layouts, _, err := layout.GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(layouts), 1)

		// GetOne
		layoutData, err = layout.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(layoutData.Id, id)
		t.Assert(layoutData.LocationId, layoutIn.LocationId)
		t.Assert(layoutData.PolicyCId, layoutIn.PolicyCId)
		t.Assert(layoutData.PolicyLId, layoutIn.PolicyLId)
		t.Assert(layoutData.LayoutName, layoutIn.LayoutName)
		t.Assert(layoutData.Map, layoutIn.Map)
		t.Assert(layoutData.Status, layoutIn.Status)
		t.Assert(layoutData.Sort, layoutIn.Sort)
		t.Assert(layoutData.Seats, layoutIn.Seats)

		// Update
		var layoutUptIn = &entity.Layout{
			Id:         id,
			LocationId: 2,
			PolicyCId:  2,
			PolicyLId:  2,
			LayoutName: "layoutTestUpt",
			Map:        `[{"x":2,"y":2}]`,
			Status:     2,
			Sort:       2,
			Seats:      2,
		}
		err = layout.Update(ctx, layoutUptIn)
		t.AssertNil(err)
		layoutData, err = layout.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(layoutData.LocationId, layoutUptIn.LocationId)
		t.Assert(layoutData.PolicyCId, layoutUptIn.PolicyCId)
		t.Assert(layoutData.PolicyLId, layoutUptIn.PolicyLId)
		t.Assert(layoutData.LayoutName, layoutUptIn.LayoutName)
		t.Assert(layoutData.Map, layoutUptIn.Map)
		t.Assert(layoutData.Status, layoutUptIn.Status)
		t.Assert(layoutData.Sort, layoutUptIn.Sort)
		t.Assert(layoutData.Seats, layoutUptIn.Seats)

		// Delete
		err = layout.Delete(ctx, id)
		t.AssertNil(err)
		_, err = layout.GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
		_, err = policy_layout.GetOne(ctx, layoutUptIn.PolicyLId)
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
		cells, err := layout.JsonToLayoutCells(ctx, jsonStr)
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
			cells = []layoutModel.Cell{
				{X: 1, Y: 1, VectorX: 3, VectorY: 3, No: 1, Label: "1号座位", Type: layoutModel.CellSeat},
				{X: 5, Y: 5, VectorX: 0, VectorY: 0, No: 2, Label: "2号座位", Type: layoutModel.CellSeat},
			}
		)
		jsonStr, err := layout.LayoutCellsToJson(ctx, cells)
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
		seats, err := layout.CalculateSeatsByJson(ctx, jsonStr)
		t.AssertNil(err)
		t.Assert(seats, 2)
	})
}
