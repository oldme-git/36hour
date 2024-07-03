package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/36hour/app/gateway/utility/data"
)

type (
	GetLayoutsReq struct {
		g.Meta `path:"seat/layouts" method:"get" sm:"获取layout列表" tags:"seat"`
	}

	GetLayoutsRes struct {
		Layouts []Layout `json:"layouts" sm:"layout列表"`
	}
)

type (
	GetLayoutReq struct {
		g.Meta `path:"seat/layout/{id}" method:"get" sm:"获取layout详情" tags:"seat"`
		*data.IdInput
	}

	GetLayoutRes struct {
		Id           data.Id `json:"id"`
		LocationId   data.Id `json:"locationId"`
		LocationName string  `json:"locationName"`
		LayoutName   string  `json:"layoutName"`
		Map          string  `json:"map"`
	}
)
