package seat_handle

import (
	"context"

	"github.com/oldme-git/36hour/app/seat/internal/dao"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/seat"
	"github.com/oldme-git/36hour/app/seat/internal/model/do"
	"github.com/oldme-git/36hour/utility"
)

// writeLog 写日志到数据库
func writeLog(ctx context.Context, layoutId int, userSeat *seat.UserSeat) (err error) {
	_, err = dao.SeatLog.Ctx(ctx).Insert(do.SeatLog{
		LocationId: 0,
		LayoutId:   layoutId,
		No:         userSeat.CellNo,
		Type:       int(userSeat.Type),
		Uid:        int(userSeat.Uid),
	})
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return
}
