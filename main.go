package main

import (
	_ "goBack/internal/packed"

	_ "goBack/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"goBack/internal/cmd"
	// 使用mysql的init()
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
