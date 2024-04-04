package main

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "libManager/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
	"libManager/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
