package main

import (
	_ "github.com/oldme-git/36hour/app/gateway/internal/packed"
	"github.com/oldme-git/36hour/utility/svc_disc"

	_ "github.com/oldme-git/36hour/app/gateway/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/oldme-git/36hour/app/gateway/internal/cmd"
)

func main() {
	var ctx = gctx.GetInitCtx()
	svc_disc.Init(ctx)
	cmd.Main.Run(ctx)
}
