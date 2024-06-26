package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/oldme-git/36hour/app/gateway/internal/controller/account"
	"github.com/oldme-git/36hour/app/gateway/internal/controller/login"
	"github.com/oldme-git/36hour/app/gateway/internal/controller/seat"
	"github.com/oldme-git/36hour/app/gateway/internal/logic/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 允许跨域
				// group.Middleware(ghttp.MiddlewareCORS)
				group.Middleware(middleware.Response)

				group.Group("/v1", func(group *ghttp.RouterGroup) {
					// App v1 路由
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Bind(login.NewV1())
						group.Bind(seat.NewV1())
					})

					// Auth Middleware
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.Bind(account.NewV1())
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
