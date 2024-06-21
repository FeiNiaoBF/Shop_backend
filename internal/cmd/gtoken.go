package cmd

import (
	"context"
	"goBack/api/backend"
	"goBack/api/frontend"
	"goBack/internal/consts"
	"goBack/internal/dao"
	"goBack/internal/model/entity"
	"goBack/utility"
	"goBack/utility/response"
	"strconv"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// StartBackendGToken 管理后台相关gtoken
func StartBackendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		CacheMode:        consts.CacheMod,
		ServerName:       consts.ServerName,
		LoginPath:        "/login",
		LoginBeforeFunc:  loginBeforeFunc,
		LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/user/logout",
		AuthPaths:        g.SliceStr{"/admin"},
		AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径
		AuthAfterFunc:    authAfterFunc,
		MultiLogin:       consts.BackendMultiLogin,
	}
	err = gfAdminToken.Start()
	return
}

func loginBeforeFunc(r *ghttp.Request) (string, any) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrNullPassword))
		r.ExitAll()
	}

	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrNullUser))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrPassword))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GTokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		g.Dump("adminId:", adminId)
		//根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &backend.LoginRes{
			Type:        consts.TokenType,
			Token:       respData.GetString("token"),
			ExpireIn:    consts.TokenExpireIn,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//账号被冻结拉黑
	if adminInfo.DeletedAt != nil {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}

// StartFrontendGToken 管理前台相关gtoken
// for 前台项目
func StartFrontendGToken() (gfFrontendToken *gtoken.GfToken, err error) {
	gfFrontendToken = &gtoken.GfToken{
		CacheMode:       consts.CacheMod,
		ServerName:      consts.ServerName,
		LoginPath:       "/login",
		LoginBeforeFunc: loginFrontendFunc,
		LoginAfterFunc:  loginAfterFrontendFunc,
		LogoutPath:      "/user/logout",
		//AuthPaths:        g.SliceStr{"/admin"},
		//AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径
		AuthAfterFunc: authAfterFunc,
		MultiLogin:    consts.FrontendMultiLogin,
	}
	//todo 去掉全局校验，只用cmd中的路由组校验
	return
}

func loginFrontendFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrNullPassword))
		r.ExitAll()
	}

	//验证账号密码是否正确
	userInfo := entity.UserInfo{}
	err := dao.UserInfo.Ctx(ctx).Where("name", name).Scan(&userInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrNullUser))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, userInfo.UserSalt) != userInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrPassword))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GTokenUserPrefix + strconv.Itoa(userInfo.Id), userInfo
}

func loginAfterFrontendFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		userId := gstr.StrEx(userKey, consts.GTokenUserPrefix)
		g.Dump("userId:", userId)
		//根据id获得登录用户其他信息
		userInfo := entity.UserInfo{}
		err := dao.UserInfo.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
		if err != nil {
			return
		}
		data := &frontend.LoginRes{
			Type:     consts.TokenType,
			Token:    respData.GetString("token"),
			ExpireIn: consts.TokenExpireIn,
			Name:     userInfo.Name,
			Avatar:   userInfo.Avatar,
			Sex:      uint8(userInfo.Sex),
			Sign:     userInfo.Sign,
			Status:   uint8(userInfo.Status),
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}
