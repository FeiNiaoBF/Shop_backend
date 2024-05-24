package cmd

import (
	"context"
	"goBack/internal/controller"

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
			s.Group("/api", func(group *ghttp.RouterGroup) {
				// 中间件
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					// subRouter
					controller.Rotation, // 轮播图
					controller.Position, // 手工位
				)
			})
			s.Run()
			return nil
		},
	}
)
