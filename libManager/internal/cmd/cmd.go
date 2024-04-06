package cmd

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
	"libManager/internal/controller/floor"
	"libManager/internal/controller/hall"
	"libManager/internal/controller/lib"
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
			hall.Register(s)
			s.Run()
			return nil
		},
	}
)
