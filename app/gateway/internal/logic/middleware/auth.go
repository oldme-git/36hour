package middleware

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	auth "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"
)

func Auth(r *ghttp.Request) {
	var (
		ctx         = r.GetCtx()
		tokenString = r.Header.Get("Authorization")
		conn        = svc_disc.UserClientConn(ctx)
		client      = auth.NewAuthClient(conn)
	)

	ctx, cancel := context.WithTimeout(ctx, svc_disc.Timeout)
	defer cancel()

	info, err := client.GetUserInfo(ctx, &auth.GetUserInfoReq{
		Token: tokenString,
	})
	if err != nil || info.GetUser().Id == 0 {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}

	r.Middleware.Next()
}
