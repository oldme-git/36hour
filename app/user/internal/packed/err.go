package packed

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type err struct {
	maps map[int]string
}

var Err = &err{
	maps: map[int]string{
		1001: "密码错误",
		1002: "不正确的或者过期的token",
	},
}

// GetMsg 获取code码对应的msg
func (c *err) GetMsg(code int) string {
	return c.maps[code]
}

func (c *err) New(code int) error {
	var msgStr = c.GetMsg(code)
	return gerror.NewCode(gcode.New(code, "", nil), msgStr)
}
