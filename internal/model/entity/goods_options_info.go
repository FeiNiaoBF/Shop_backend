// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsOptionsInfo is the golang structure for table goods_options_info.
type GoodsOptionsInfo struct {
	Id        int         `json:"id"         orm:"id"         description:""`
	GoodsId   int         `json:"goods_id"   orm:"goods_id"   description:"商品id"`
	PicUrl    string      `json:"pic_url"    orm:"pic_url"    description:"图片"`
	Name      string      `json:"name"       orm:"name"       description:"商品名称"`
	Price     int         `json:"price"      orm:"price"      description:"价格 单位分"`
	Stock     int         `json:"stock"      orm:"stock"      description:"库存"`
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:""`
}
