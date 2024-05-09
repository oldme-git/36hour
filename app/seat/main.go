package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/oldme-git/36hour/app/seat/internal/cmd"
	"github.com/oldme-git/36hour/utility/svc_disc"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/oldme-git/36hour/app/seat/internal/logic"
)

func main() {
	var ctx = gctx.GetInitCtx()
	svc_disc.Init(ctx)
	cmd.Main.Run(gctx.GetInitCtx())
}
