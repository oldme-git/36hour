package main

import (
	"github.com/oldme-git/36hour/app/user/internal/cmd"
	"github.com/oldme-git/36hour/utility/svc_disc"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/oldme-git/36hour/app/user/internal/logic/snowflake"
)

func main() {
	var ctx = gctx.GetInitCtx()
	svc_disc.Init(ctx)
	cmd.Main.Run(ctx)
}
