package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CollectionBase struct {
	UserId   uint  `json:"user_id"    description:"用户id"`
	ObjectId uint  `json:"object_id"  description:"对象id"`
	Type     uint8 `json:"type"      description:"收藏类型：1商品 2文章"`
}

type CollectionAddInput struct {
	CollectionBase
}

type CollectionAddOutput struct {
	Id uint
}

type CollectionDeleteInput struct {
	Id uint
	CollectionBase
}

type CollectionDeleteOutput struct {
	Id uint
}

// CollectionGetALLListInput 获取内容列表
type CollectionGetALLListInput struct {
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CollectionGetAllListOutput 查询列表结果
type CollectionGetAllListOutput struct {
	List  []CollectionListItem `json:"list" description:"列表"`
	Total int                  `json:"total" description:"数据总数"`
}

// CollectionGetListInput 获取内容列表
type CollectionGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CollectionGetListOutput 查询列表结果
type CollectionGetListOutput struct {
	List  []CollectionListItem `json:"list" description:"列表"`
	Page  int                  `json:"page" description:"分页码"`
	Size  int                  `json:"size" description:"分页数量"`
	Total int                  `json:"total" description:"数据总数"`
}

// CollectionListItem 单一对象
type CollectionListItem struct {
	Id int `json:"id"`
	CollectionBase
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article" orm:"with:id=object_id"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type GoodsItem struct {
	g.Meta `orm:"table:goods_info"`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	PicUrl string `json:"pic_url"`
	Price  int    `json:"price"`
}

type ArticleItem struct {
	g.Meta `orm:"table:article_info"`
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	PicUrl string `json:"pic_url"`
}
