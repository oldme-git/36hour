package utility

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type err struct {
	maps map[int]string
}

const CodeOk = 0

// CodeErrSys 系统级别的错误码，不应该把错误信息抛出到客户端
const CodeErrSys = -1

var Err = &err{
	maps: map[int]string{
		// general
		CodeOk:     "成功",
		CodeErrSys: "未知错误",

		// user
		1001: "用户名或密码错误",
		1002: "不正确的或者过期的token",

		// lib-manager
		2001: "图书馆不存在",

		// seat
		3001: "座位布局不是合法的 json 格式",
		3002: "请为场馆设置一个策略",
	},
}

// GetMsg 获取code码对应的msg
func (c *err) GetMsg(code int) string {
	return c.maps[code]
}

// New 生成一个新的业务级别的错误，不会打印错误堆栈信息
func (c *err) New(code int) error {
	var msgStr = c.GetMsg(code)
	return gerror.NewWithOption(gerror.Option{
		Stack: false,
		Text:  msgStr,
		Code:  gcode.New(code, "", nil),
	})
}

// NewSys 生成一个新的系统级别的错误，使用特殊的code码：CodeErrSys
// !!! 使用该方法时，它会打印错误堆栈信息到日志，但是一定不要把任何错误信息抛出到客户端，防止泄露系统信息
// !!! 推荐做法是在后置中间件中捕获 code -1 的错误，然后返回给客户端一个统一的错误提示
func (c *err) NewSys(err error) error {
	return gerror.NewCode(gcode.New(CodeErrSys, "", nil), err.Error())
}

// GetOkMsg 获取成功的msg
func (c *err) GetOkMsg() string {
	return c.GetMsg(CodeOk)
}

// GetSysMsg 获取系统错误的msg
func (c *err) GetSysMsg() string {
	return c.GetMsg(CodeErrSys)
}
