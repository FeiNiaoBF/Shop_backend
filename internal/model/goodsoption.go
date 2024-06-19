package model

import (
	"goBack/internal/model/entity"
)

// GoodsOptionsBase 创建/修改内容基类
type GoodsOptionsBase struct {
	GoodsId uint   `json:"goods_id" sm:"主商品id"`
	PicUrl  string `json:"pic_url"  sm:"图片"`
	Name    string `json:"name"     sm:"商品规格名称" v:"required#商品规格名称必传"`
	Price   int    `json:"price"    sm:"价格 单位分" v:"required#商品规格价格必传"`
	Stock   int    `json:"stock"    sm:"库存"`
}

// GoodsOptionsCreateInput 创建内容
type GoodsOptionsCreateInput struct {
	GoodsOptionsBase
}

// GoodsOptionsCreateOutput 创建内容返回结果
type GoodsOptionsCreateOutput struct {
	GoodsOptionsId int `json:"goods_id"`
}

// GoodsOptionsUpdateInput 修改内容
type GoodsOptionsUpdateInput struct {
	Id uint `json:"id"`
	GoodsOptionsBase
}

// GoodsOptionsUpdateOutput 修改内容
type GoodsOptionsUpdateOutput struct {
	Id uint
	// UpdatedAt string `json:"updated_at"` // 修改时间
}

// GoodsOptionsDeleteInput 删除内容
type GoodsOptionsDeleteInput struct {
	Id uint `json:"id"`
}

// GoodsOptionsDeleteOutput 删除内容
type GoodsOptionsDeleteOutput struct {
	Id        uint   `json:"id"`
	DeletedAt string `json:"deleted_at"` // 修改时间
}

// GoodsOptionsGetALLListInput 获取内容列表
type GoodsOptionsGetALLListInput struct {
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GoodsOptionsGetAllListOutput 查询列表结果
type GoodsOptionsGetAllListOutput struct {
	List  []GoodsOptionsListItem `json:"list" description:"列表"`
	Total int                    `json:"total" description:"数据总数"`
}

// GoodsOptionsGetListInput 获取内容列表
type GoodsOptionsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GoodsOptionsGetListOutput 查询列表结果
type GoodsOptionsGetListOutput struct {
	List  []GoodsOptionsListItem `json:"list" description:"列表"`
	Page  int                    `json:"page" description:"分页码"`
	Size  int                    `json:"size" description:"分页数量"`
	Total int                    `json:"total" description:"数据总数"`
}

// GoodsOptionsListItem 单一对象
type GoodsOptionsListItem struct {
	entity.GoodsOptionsInfo
}
