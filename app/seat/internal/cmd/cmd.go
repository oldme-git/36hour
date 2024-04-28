package cmd

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/oldme-git/36hour/app/seat/internal/controller/layout"
	"github.com/oldme-git/36hour/app/seat/internal/controller/policy_common"
	"github.com/oldme-git/36hour/app/seat/internal/controller/policy_prepare"
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
			policy_prepare.Register(s)
			policy_common.Register(s)
			layout.Register(s)
			s.Run()
			return nil
		},
	}
)
