// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure for table admin_info.
type AdminInfo struct {
	Id        int         `json:"id"         orm:"id"         description:""`
	Name      string      `json:"name"       orm:"name"       description:"用户名"`
	Password  string      `json:"password"   orm:"password"   description:"密码"`
	RoleIds   string      `json:"role_ids"   orm:"role_ids"   description:"角色ids"`
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`
	UserSalt  string      `json:"user_salt"  orm:"user_salt"  description:"加密盐"`
	IsAdmin   int         `json:"is_admin"   orm:"is_admin"   description:"是否超级管理员"`
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:""`
}
