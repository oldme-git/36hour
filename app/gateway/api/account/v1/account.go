package v1

import "github.com/gogf/gf/v2/frame/g"

type (
	InfoReq struct {
		g.Meta `path:"account/info" method:"get" sm:"获取账户信息" tags:"账户"`
	}

	InfoRes struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Phone    string `json:"phone"`
		Lib      *Lib   `json:"lib"`
	}
)
