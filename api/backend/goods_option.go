package backend

import "github.com/gogf/gf/v2/frame/g"

type GoodsOptionsBase struct {
	GoodsId uint   `json:"goods_id" sm:"主商品id"`
	PicUrl  string `json:"pic_url"  sm:"图片"`
	Name    string `json:"name"     sm:"商品规格名称" v:"required#商品规格名称必传"`
	Price   int    `json:"price"    sm:"价格 单位分" v:"required#商品规格价格必传"`
	Stock   int    `json:"stock"    sm:"库存"`
}

type GoodsOptionsCreateReq struct {
	g.Meta `path:"/goods/options/add" tags:"商品分类" method:"post" sm:"添加商品规格接口"`
	GoodsOptionsBase
}

type GoodsOptionsCreateRes struct {
	Id uint `json:"id"`
}

type GoodsOptionsUpdateReq struct {
	g.Meta `path:"/goods/options/update" tags:"商品分类" method:"post" sm:"更新商品规格接口"`
	Id     uint `json:"id" v:"min:1#请选择需要更新的商品" sm:"商品id"`
	GoodsOptionsBase
}

type GoodsOptionsUpdateRes struct {
	Id uint `json:"id"`
	// UpdatedAt string `json:"updated_at"` // 修改时间
}

type GoodsOptionsDeleteReq struct {
	g.Meta `path:"/goods/options/delete" tags:"商品分类" method:"delete" sm:"删除商品规格接口"`
	Id     int `json:"id" v:"min:1#请选择需要删除的商品" dc:"商品分类id"`
}

type GoodsOptionsDeleteRes struct {
	// DeletedAt string `json:"deleted_at"` // 删除时间
}

type GoodsOptionsListReq struct {
	g.Meta `path:"/goods/options/list" tags:"商品分类" method:"get" sm:"列表商品规格接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type GoodsOptionsListRes struct {
	List  interface{} `json:"list" sm:"列表"`
	Page  int         `json:"page" sm:"分页码"`
	Size  int         `json:"size" sm:"分页数量"`
	Total int         `json:"total" sm:"数据总数"`
}

type GoodsOptionsAllListReq struct {
	g.Meta `path:"/goods/options/list/all" tags:"商品分类" method:"get" sm:"列表全部商品规格接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
}

type GoodsOptionsAllListRes struct {
	List  interface{} `json:"list" sm:"列表"`
	Total int         `json:"total" sm:"数据总数"`
}
