// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryInfo is the golang structure for table category_info.
type CategoryInfo struct {
	Id        int         `json:"id"         orm:"id"         description:""`
	ParentId  int         `json:"parent_id"  orm:"parent_id"  description:"父级id"`
	Name      string      `json:"name"       orm:"name"       description:""`
	PicUrl    string      `json:"pic_url"    orm:"pic_url"    description:"icon"`
	Level     int         `json:"level"      orm:"level"      description:"等级 默认1级分类"`
	Sort      int         `json:"sort"       orm:"sort"       description:""`
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:""`
}
