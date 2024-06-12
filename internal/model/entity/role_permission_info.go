// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermissionInfo is the golang structure for table role_permission_info.
type RolePermissionInfo struct {
	Id           int         `json:"id"            orm:"id"            description:""`
	RoleId       int         `json:"role_id"       orm:"role_id"       description:"角色id"`
	PermissionId int         `json:"permission_id" orm:"permission_id" description:"权限id"`
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updated_at"    orm:"updated_at"    description:""`
}
