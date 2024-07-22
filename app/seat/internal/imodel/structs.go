package imodel

import "github.com/oldme-git/36hour/app/seat/internal/imodel/layout"

type Id int64

// UserSeat 用户座位信息
type UserSeat struct {
	Uid        Id                `json:"uid"`
	LocationId uint              `json:"location_id"`
	CellNo     int               `json:"cell_no"`
	Status     layout.CellStatus `json:"status"`
}
