package backend

import (
	"goBack/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求接口" tags:"登录"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

type LoginDoRes struct {
	Info   interface{} `json:"info"`
	Token  string      `json:"token"`
	Expire time.Time   `json:"expire"`
}

type LoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expire_in"`
	IsAdmin     int                     `json:"is_admin"`    //是否超管
	RoleIds     string                  `json:"role_ids"`    //角色
	Permissions []entity.PermissionInfo `json:"permissions"` //权限列表
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post" summary:"刷新登录接口" tags:"登录"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"登录" summary:"退出接口"`
}

type LogoutRes struct {
}
