package cmd

import (
	"context"
	"goBack/internal/controller"
	"goBack/internal/service"

	"github.com/goflyfox/gtoken/gtoken"
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
			// 启动gtoken
			gfAdminToken := &gtoken.GfToken{
				CacheMode:        2,
				ServerName:       "shop",
				LoginPath:        "/login",
				LoginBeforeFunc:  loginBeforeFunc,
				LoginAfterFunc:   loginAfterFunc,
				LogoutPath:       "/user/logout",
				AuthPaths:        g.SliceStr{"/admin"},
				AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径
				AuthAfterFunc:    authAfterFunc,
				MultiLogin:       true,
			}
			s.Group("/api", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Admin.Create, // 管理员
					controller.Admin.Update, // 管理员
					controller.Admin.Delete, // 管理员
					controller.Admin.List,   // 管理员
					controller.Rotation,     // 轮播图
					controller.Position,     // 手工位
					controller.Login,        // 登录
					controller.Data,         // 数据
					controller.Role,         //管理角色
					controller.Permission,   // 权限

				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					// group.Middleware(service.Middleware().Auth)
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.ALLMap(g.Map{
						"/admin/info": controller.Admin.Info,
					})
					group.Bind(
						controller.File,   // 本地文件
						controller.Upload, // 云平台
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
