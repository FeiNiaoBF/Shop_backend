package cmd

import (
	"context"
	"goBack/internal/consts"
	"goBack/internal/controller"
	"goBack/internal/service"
	"log"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动gtoken
			gfAdminToken, err := SetGtoken()
			if err != nil {
				log.Println(err)
			}
			// main URL
			s.Group("/", func(group *ghttp.RouterGroup) {
				// URL: /api/
				s.Group("/api", func(group *ghttp.RouterGroup) {
					//不需要登录的路由组绑定
					// 中间件
					group.Middleware(
						service.Middleware().CORS,            // 跨域资源
						service.Middleware().Ctx,             // 上下文
						service.Middleware().ResponseHandler, //
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
							log.Println(err)
						}
						group.ALLMap(g.Map{
							"/admin/info": controller.Admin.Info,
						})
						group.Bind(
							controller.File,     // 本地文件
							controller.Upload,   // 云平台
							controller.Category, // 商品分类管理
						)
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
