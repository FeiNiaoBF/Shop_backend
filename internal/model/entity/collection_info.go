// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CollectionInfo is the golang structure for table collection_info.
type CollectionInfo struct {
	Id        int         `json:"id"         orm:"id"         description:""`
	UserId    int         `json:"user_id"    orm:"user_id"    description:"用户id"`
	ObjectId  int         `json:"object_id"  orm:"object_id"  description:"对象id"`
	Type      int         `json:"type"       orm:"type"       description:"收藏类型：1商品 2文章"`
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`
}
