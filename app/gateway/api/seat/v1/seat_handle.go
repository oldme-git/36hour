package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/36hour/app/gateway/utility/data"
)

type (
	ReserveReq struct {
		g.Meta     `path:"seat/reserve" method:"post" sm:"预约座位" tags:"seat"`
		LocationId data.Id `json:"locationId"`
		CellNo     int     `json:"cellNo"`
	}

	ReserveRes struct {
	}
)
