package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/oldme-git/36hour/app/seat/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/oldme-git/36hour/app/seat/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
