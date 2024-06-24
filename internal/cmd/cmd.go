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
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				log.Println(err)
			}
			gfFrontendToken, err := StartFrontendGToken()
			if err != nil {
				log.Println(err)
			}
			// main URL
			s.Group("/", func(group *ghttp.RouterGroup) {
				// URL: /backend/
				// backend
				s.Group("/backend", func(group *ghttp.RouterGroup) {
					//不需要登录的路由组绑定
					// 中间件
					group.Middleware(
						service.Middleware().CORS,            // 跨域资源
						service.Middleware().Ctx,             // 上下文
						service.Middleware().ResponseHandler, //
					)
					group.Bind(
						controller.Admin.Create, // 管理员

						controller.Login, // 登录

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
							controller.Admin.Update, // 管理员
							controller.Admin.Delete, // 管理员
							controller.Admin.List,   // 管理员
							controller.Permission,   // 权限
							controller.Role,         //管理角色
							controller.Data,         // 数据

							controller.Rotation, // 轮播图
							controller.Position, // 手工位
							controller.File,     // 本地文件
							controller.Upload,   // 云平台

							controller.Category,     //商品分类管理
							controller.Coupon,       //商品优惠券管理
							controller.UserCoupon,   //用户商品优惠券管理
							controller.Goods,        // 商品管理
							controller.GoodsOptions, //商品规格管理
							controller.Article,      //文章管理
						)
					})
				})
				// frontend
				s.Group("/frontend", func(group *ghttp.RouterGroup) {
					// 中间件
					group.Middleware(
						service.Middleware().CORS,            // 跨域资源
						service.Middleware().Ctx,             // 上下文
						service.Middleware().ResponseHandler, //
					)
					group.Bind(
						controller.User.Register, //用户注册
					)
					//需要登录鉴权的路由组
					group.Group("/user", func(group *ghttp.RouterGroup) {
						err := gfFrontendToken.Middleware(ctx, group)
						if err != nil {
							return
						}
						//需要登录鉴权的接口放到这里
						group.Bind(
							controller.User.Info, //当前登录用户的信息
						)
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
