package cmd

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/oldme-git/36hour/app/lib-manager/internal/controller/floor"
	"github.com/oldme-git/36hour/app/lib-manager/internal/controller/lib"
	"github.com/oldme-git/36hour/app/lib-manager/internal/controller/location"
	"google.golang.org/grpc"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "user grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			lib.Register(s)
			floor.Register(s)
			location.Register(s)
			s.Run()
			return nil
		},
	}
)
