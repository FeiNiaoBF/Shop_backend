package cmd

import (
	"context"
	controller "goBack/internal/controller/rotation"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// root routes
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					// subRouter
					controller.Rotation, // 内容
				)
			})
			s.Run()
			return nil
		},
	}
)
