package seat

import (
	"github.com/oldme-git/36hour/app/seat/internal/imodel"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/layout"
)

// Type 操作类型，1: 预定，2: 签到，3: 退坐
type Type int

const (
	TypeReserve Type = iota + 1
	TypeSign
	TypeCancel
)

// UserSeat 用户座位信息
type UserSeat struct {
	Uid        imodel.Id         `json:"uid"`
	LocationId int               `json:"location_id"`
	CellNo     int               `json:"cell_no"`
	Type       Type              `json:"type"`
	Status     layout.CellStatus `json:"status"`
}
