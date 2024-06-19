package model

import (
	"goBack/internal/model/entity"
)

// GoodsBase 创建/修改内容基类
type GoodsBase struct {
	Name             string `json:"name" v:"required#名称必填" sm:"商品名称"`
	PicURL           string `json:"pic_url" v:"required#商品图片必填" sm:"商品图片链接"`
	Price            int    `json:"price" v:"required#商品金额必填" sm:"商品金额"`
	Level1CategoryId int    `json:"level1_category_id" sm:"1级分类id"`
	Level2CategoryId int    `json:"level2_category_id" sm:"2级分类id"`
	Level3CategoryId int    `json:"level3_category_id" sm:"3级分类id"`
	Brand            string `json:"brand" sm:"品牌" v:"max-length:30#品牌名称最大30个字"`
	Stock            int    `json:"stock" sm:"库存"`
	Sale             int    `json:"sale" sm:"销量"`
	Tags             string `json:"tags" sm:"标签"`
	DetailInfo       string `json:"detail_info" sm:"商品详情"`
}

// GoodsCreateInput 创建内容
type GoodsCreateInput struct {
	GoodsBase
}

// GoodsCreateOutput 创建内容返回结果
type GoodsCreateOutput struct {
	GoodsId int `json:"goods_id"`
}

// GoodsUpdateInput 修改内容
type GoodsUpdateInput struct {
	Id uint `json:"id"`
	GoodsBase
}

// GoodsUpdateOutput 修改内容
type GoodsUpdateOutput struct {
	Id        uint
	UpdatedAt string `json:"updated_at"` // 修改时间
}

// GoodsDeleteInput 删除内容
type GoodsDeleteInput struct {
	Id uint `json:"id"`
}

// GoodsDeleteOutput 删除内容
type GoodsDeleteOutput struct {
	Id        uint   `json:"id"`
	DeletedAt string `json:"deleted_at"` // 修改时间
}

// GoodsGetALLListInput 获取内容列表
type GoodsGetALLListInput struct {
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GoodsGetAllListOutput 查询列表结果
type GoodsGetAllListOutput struct {
	List  []GoodsListItem `json:"list" description:"列表"`
	Total int             `json:"total" description:"数据总数"`
}

// GoodsGetListInput 获取内容列表
type GoodsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GoodsGetListOutput 查询列表结果
type GoodsGetListOutput struct {
	List  []GoodsListItem `json:"list" description:"列表"`
	Page  int             `json:"page" description:"分页码"`
	Size  int             `json:"size" description:"分页数量"`
	Total int             `json:"total" description:"数据总数"`
}

// GoodsListItem 单一对象
type GoodsListItem struct {
	entity.GoodsInfo
}
