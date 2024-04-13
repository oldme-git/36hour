package main

import (
	_ "user/internal/packed"

	_ "seat/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"user/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
