package v1

import "github.com/gogf/gf/v2/frame/g"

type GetLayoutListReq struct {
	g.Meta `path:"seat/layout/list" method:"get" sm:"获取layout列表" tags:"seat_layout"`
}

type GetLayoutListRes struct {
	Id    int    `json:"id" sm:"id"`
	Name  string `json:"name" sm:"名称"`
	Seats int    `json:"seats" sm:"座位数"`
}
