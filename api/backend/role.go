package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/role/add" method:"post" summary:"添加角色角色接口" tags:"权限角色"`
	Name   string `json:"name" v:"required#角色职位不能为空" dc:"职位"`
	Desc   string `json:"desc" dc:"职位描述"`
}

type RoleRes struct {
	RoleId int `json:"roleId"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"delete" summary:"删除角色接口" tags:"权限角色"`
	Id     uint `v:"min:1#请选择需要删除的角色" dc:"角色"`
}
type RoleDeleteRes struct{}

type RoleUpdateReq struct {
	g.Meta `path:"/role/update" method:"post" summary:"更新角色接口" tags:"权限角色"`
	Id     uint   `json:"id" v:"min:1#请选择需要修改的角色" dc:"角色Id"`
	Name   string `json:"name" v:"required#角色职位不能为空" dc:"职位"`
	Desc   string `json:"desc" dc:"职位描述"`
}
type RoleUpdateRes struct {
	Id uint `json:"id"`
}
type RoleGetListCommonReq struct {
	g.Meta `path:"/role/list" method:"get" summary:"列表角色接口" tags:"权限角色"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type RoleGetInfoReq struct {
	g.Meta `path:"/role/info" method:"get"`
}

type AddPermissionReq struct {
	g.Meta       `path:"/role/add/permission" method:"post" tags:"权限角色" summary:"角色添加权限接口"`
	RoleId       uint `json:"role_id" desc:"角色id"`
	PermissionId uint `json:"permission_id" desc:"权限id"`
}

type AddPermissionRes struct {
	Id uint `json:"id"`
}

type DeletePermissionReq struct {
	g.Meta       `path:"/role/delete/permission" method:"delete" tags:"权限角色" summary:"角色删除权限接口"`
	RoleId       uint `json:"role_id" desc:"角色id"`
	PermissionId uint `json:"permission_id" desc:"权限id"`
}

type DeletePermissionRes struct {
}
