package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"seat/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "seat/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
