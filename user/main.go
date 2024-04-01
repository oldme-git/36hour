package main

import (
	"user/internal/cmd"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "user/internal/logic"
)

func main() {
	// TODO 注册到etcd，应该封装一下使用
	grpcx.Resolver.Register(etcd.New("srv.com:2379"))
	cmd.Main.Run(gctx.GetInitCtx())
}
