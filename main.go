package main

import (
	_ "36hour/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"36hour/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
