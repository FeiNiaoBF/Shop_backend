// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsInfo is the golang structure for table goods_info.
type GoodsInfo struct {
	Id               int         `json:"id"                  orm:"id"                 description:""`
	PicUrl           string      `json:"pic_url"             orm:"pic_url"            description:"图片"`
	Name             string      `json:"name"                orm:"name"               description:"商品名称"`
	Price            int         `json:"price"               orm:"price"              description:"价格 单位分"`
	Level1CategoryId int         `json:"level_1_category_id" orm:"level1_category_id" description:"1级分类id"`
	Level2CategoryId int         `json:"level_2_category_id" orm:"level2_category_id" description:"2级分类id"`
	Level3CategoryId int         `json:"level_3_category_id" orm:"level3_category_id" description:"3级分类id"`
	Brand            string      `json:"brand"               orm:"brand"              description:"品牌"`
	Stock            int         `json:"stock"               orm:"stock"              description:"库存"`
	Sale             int         `json:"sale"                orm:"sale"               description:"销量"`
	Tags             string      `json:"tags"                orm:"tags"               description:"标签"`
	DetailInfo       string      `json:"detail_info"         orm:"detail_info"        description:"商品详情"`
	CreatedAt        *gtime.Time `json:"created_at"          orm:"created_at"         description:""`
	UpdatedAt        *gtime.Time `json:"updated_at"          orm:"updated_at"         description:""`
	DeletedAt        *gtime.Time `json:"deleted_at"          orm:"deleted_at"         description:""`
}
