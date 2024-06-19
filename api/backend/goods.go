package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsBase struct {
	Name             string `json:"name" v:"required#名称必填" sm:"商品名称"`
	PicURL           string `json:"pic_url" v:"required#商品图片必填" sm:"商品图片链接"`
	Price            int    `json:"price" v:"required#商品金额必填" sm:"商品金额"`
	Level1CategoryId uint   `json:"level1_category_id" sm:"1级分类id"`
	Level2CategoryId uint   `json:"level2_category_id" sm:"2级分类id"`
	Level3CategoryId uint   `json:"level3_category_id" sm:"3级分类id"`
	Brand            string `json:"brand" sm:"品牌" v:"max-length:30#品牌名称最大30个字"`
	Stock            int    `json:"stock" sm:"库存"`
	Sale             int    `json:"sale" sm:"销量"`
	Tags             string `json:"tags" sm:"标签"`
	DetailInfo       string `json:"detail_info" sm:"商品详情"`
}

type GoodsCreateReq struct {
	g.Meta `path:"/goods/add" tags:"商品分类" method:"post" sm:"添加商品种类接口"`
	GoodsBase
}

type GoodsCreateRes struct {
	Id uint `json:"id"`
}

type GoodsUpdateReq struct {
	g.Meta `path:"/goods/update" tags:"商品分类" method:"post" sm:"更新商品种类接口"`
	Id     uint `json:"id" v:"min:1#请选择需要更新的商品" sm:"商品id"`
	GoodsBase
}

type GoodsUpdateRes struct {
	Id uint `json:"id"`
	// UpdatedAt string `json:"updated_at"` // 修改时间
}

type GoodsDeleteReq struct {
	g.Meta `path:"/goods/delete" tags:"商品分类" method:"delete" sm:"删除商品种类接口"`
	Id     int `json:"id" v:"min:1#请选择需要删除的商品" dc:"商品分类id"`
}

type GoodsDeleteRes struct {
	// DeletedAt string `json:"deleted_at"` // 删除时间
}

type GoodsListReq struct {
	g.Meta `path:"/goods/list" tags:"商品分类" method:"get" sm:"列表商品种类接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type GoodsListRes struct {
	List  interface{} `json:"list" sm:"列表"`
	Page  int         `json:"page" sm:"分页码"`
	Size  int         `json:"size" sm:"分页数量"`
	Total int         `json:"total" sm:"数据总数"`
}

type GoodsAllListReq struct {
	g.Meta `path:"/goods/list/all" tags:"商品分类" method:"get" summary:"列表全部商品种类接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
}

type GoodsAllListRes struct {
	List  interface{} `json:"list" sm:"列表"`
	Total int         `json:"total" sm:"数据总数"`
}
