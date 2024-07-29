package seat_handle_test

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/seat"
	"github.com/oldme-git/36hour/app/seat/internal/logic/seat_handle"
	"github.com/oldme-git/36hour/utility"
)

func TestReserve(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx      = context.TODO()
			userSeat = seat.UserSeat{
				LocationId: 9,
				CellNo:     1,
				Uid:        1805425223755894,
				Type:       seat.TypeReserve,
				Status:     seat_handle.GetCellStatus(ctx, seat.TypeReserve),
			}
		)
		err := seat_handle.Reserve(ctx, &userSeat)
		t.AssertNil(err)

		err = seat_handle.Reserve(ctx, &userSeat)
		t.Assert(err, utility.Err.New(3003))
	})
}
