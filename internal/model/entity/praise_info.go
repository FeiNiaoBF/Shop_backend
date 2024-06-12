// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PraiseInfo is the golang structure for table praise_info.
type PraiseInfo struct {
	Id        int         `json:"id"         orm:"id"         description:"点赞表"`
	UserId    int         `json:"user_id"    orm:"user_id"    description:""`
	Type      int         `json:"type"       orm:"type"       description:"点赞类型 1商品 2文章"`
	ObjectId  int         `json:"object_id"  orm:"object_id"  description:"点赞对象id 方便后期扩展"`
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`
}
