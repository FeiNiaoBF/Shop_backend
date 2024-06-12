// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CouponInfo is the golang structure for table coupon_info.
type CouponInfo struct {
	Id         int         `json:"id"          orm:"id"          description:""`
	Name       string      `json:"name"        orm:"name"        description:""`
	Price      int         `json:"price"       orm:"price"       description:"优惠前面值 单位分"`
	GoodsIds   string      `json:"goods_ids"   orm:"goods_ids"   description:"关联使用的goods_ids  逗号分隔"`
	CategoryId int         `json:"category_id" orm:"category_id" description:"关联使用的分类id"`
	CreatedAt  *gtime.Time `json:"created_at"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updated_at"  orm:"updated_at"  description:""`
	DeletedAt  *gtime.Time `json:"deleted_at"  orm:"deleted_at"  description:""`
}
