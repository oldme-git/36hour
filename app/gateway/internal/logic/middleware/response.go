// 统一响应处理

package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/oldme-git/36hour/utility"
	"google.golang.org/grpc/codes"
)

type data struct {
	Code    int         `json:"code"    dc:"业务码"`
	Message string      `json:"message" dc:"业务码说明"`
	Data    interface{} `json:"data"    dc:"返回的数据"`
}

func Response(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	// 先过滤掉服务器内部错误
	if r.Response.Status >= http.StatusInternalServerError {
		// 清除掉缓存区，防止服务器信息泄露到客户端
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器打盹了，请稍后再来找他！")
	}

	var (
		res  = r.GetHandlerResponse()
		err  = r.GetError()
		code = gerror.Code(err).Code()
		msg  string
	)

	if err != nil {
		// 如果是 grpc Unknown 的错误码，则转换为系统错误
		if code == int(codes.Unknown) {
			code = utility.CodeErrSys
		}
		// 如果是系统错误，不要把错误信息抛出到客户端，防止泄露系统信息
		if code == utility.CodeErrSys {
			msg = utility.Err.GetSysMsg()
		} else {
			msg = err.Error()
		}
	} else {
		code = utility.CodeOk
		msg = utility.Err.GetOkMsg()
	}

	r.Response.WriteJson(data{
		Code:    code,
		Message: msg,
		Data:    res,
	})
}
