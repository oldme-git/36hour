package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/oldme-git/36hour/app/gateway/internal/cmd"
	"github.com/oldme-git/36hour/utility/svc_disc"

	_ "github.com/oldme-git/36hour/app/gateway/internal/packed"
)

func main() {
	var ctx = gctx.GetInitCtx()
	svc_disc.Init(ctx)
	cmd.Main.Run(ctx)
}
