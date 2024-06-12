// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure for table file_info.
type FileInfo struct {
	Id        int         `json:"id"         orm:"id"         description:""`
	Name      string      `json:"name"       orm:"name"       description:"图片名称"`
	Src       string      `json:"src"        orm:"src"        description:"本地文件存储路径"`
	Url       string      `json:"url"        orm:"url"        description:"URL地址"`
	UserId    int         `json:"user_id"    orm:"user_id"    description:"用户id"`
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`
}
