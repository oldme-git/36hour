package v1

import "github.com/gogf/gf/v2/frame/g"

type GetLayoutsReq struct {
	g.Meta `path:"seat/layouts" method:"get" sm:"获取layout列表" tags:"seat"`
}

type GetLayoutsRes struct {
	Layouts []Layout `json:"layouts" sm:"layout列表"`
}
