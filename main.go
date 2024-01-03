package main

import (
	_ "gfdemo/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	_ "gfdemo/internal/logic"

	"gfdemo/internal/cmd"
)

func main() {

	cmd.Main.Run(gctx.GetInitCtx())
}
