package main

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/oldme-git/36hour/app/lib-manager/internal/cmd"
	"github.com/oldme-git/36hour/utility/svc_disc"
)

func main() {
	var ctx = gctx.GetInitCtx()
	svc_disc.Init(ctx)
	cmd.Main.Run(ctx)
}
