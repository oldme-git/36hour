package packed

import (
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type err struct {
	maps map[int]string
}

var Err = &err{
	maps: map[int]string{
		0:     "请求成功",
		99999: "未知错误",
	},
}

// GetMsg 获取code码对应的msg
func (c *err) GetMsg(code int) string {
	return c.maps[code]
}

// Skip 抛出一个业务级别的错误，不会打印错误堆栈信息
func (c *err) Skip(code int, msg ...string) (err error) {
	var msgStr string
	if len(msg) == 0 {
		msgStr = c.GetMsg(code)
	} else {
		msg = append([]string{c.GetMsg(code)}, msg...)
		msgStr = strings.Join(msg, ", ")
	}
	return gerror.NewWithOption(gerror.Option{
		Stack: false,
		Text:  msgStr,
		Code:  gcode.New(code, "", nil),
	})
}

// Sys 抛出一个系统级别的错误，使用code码：99999，会打印错误堆栈信息
// msg 接受string和error类型
// !!! 使用该方法传入error类型时，一定要注意不要泄露系统信息
func (c *err) Sys(msg ...interface{}) error {
	var (
		code     = 99999
		msgSlice = []string{
			c.GetMsg(code),
		}
	)

	if len(msg) != 0 {
		for _, v := range msg {
			switch a := v.(type) {
			case error:
				msgSlice = append(msgSlice, a.Error())
			case string:
				msgSlice = append(msgSlice, a)
			}
		}
	}

	msgStr := strings.Join(msgSlice, ", ")
	return gerror.NewCode(gcode.New(code, "", nil), msgStr)
}

// SysDb 所有关于数据库操作的错误都要使用此方法抛出错误，防止数据表结构泄露
func (c *err) SysDb(handel string, table string) error {
	return c.Sys(handel + " " + table + " err")
}
