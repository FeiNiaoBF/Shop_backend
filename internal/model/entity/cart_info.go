// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CartInfo is the golang structure for table cart_info.
type CartInfo struct {
	Id             int         `json:"id"               orm:"id"               description:"购物车表"`
	UserId         int         `json:"user_id"          orm:"user_id"          description:""`
	GoodsOptionsId int         `json:"goods_options_id" orm:"goods_options_id" description:"商品规格id"`
	Count          int         `json:"count"            orm:"count"            description:"商品数量"`
	CreatedAt      *gtime.Time `json:"created_at"       orm:"created_at"       description:""`
	UpdatedAt      *gtime.Time `json:"updated_at"       orm:"updated_at"       description:""`
	DeletedAt      *gtime.Time `json:"deleted_at"       orm:"deleted_at"       description:""`
}
