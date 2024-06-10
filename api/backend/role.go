package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/role/add" method:"post" desc:"添加管理角色" tags:"role"`
	Name   string `json:"name" v:"required#管理职位不能为空" dc:"职位"`
	Desc   string `json:"desc" dc:"职位描述"`
}

type RoleRes struct {
	RoleId int `json:"roleId"`
}
