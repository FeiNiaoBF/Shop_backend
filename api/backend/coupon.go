package backend

import "github.com/gogf/gf/v2/frame/g"

type CouponBase struct {
	Name       string `json:"name" v:"required#名称必填" dc:"名称"`
	Price      int    `json:"price" v:"required#优惠券金额" dc:"优惠券金额"`
	GoodsIds   string `json:"goods_ids" dc:"可用的商品id 多个够好分隔"`
	CategoryId uint   `json:"category_id"  dc:"可用的商品分类"`
}

type CouponCreateReq struct {
	g.Meta `path:"/coupon/add" tags:"优惠券分类" method:"post" summary:"添加优惠券种类接口"`
	CouponBase
}

type CouponCreateRes struct {
	Id int `json:"id"`
}

type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update" tags:"优惠券分类" method:"post" summary:"更新优惠券种类接口"`
	Id     uint `json:"id" v:"min:1#请选择需要更新的优惠券" dc:"优惠券id"`
	CouponBase
}

type CouponUpdateRes struct {
	Id int `json:"id"`
}

type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete" tags:"优惠券分类" method:"delete" summary:"删除优惠券种类接口"`
	Id     int `json:"id" v:"min:1#请选择需要删除的优惠券" dc:"优惠券分类id"`
}

type CouponDeleteRes struct{}

type CouponListReq struct {
	g.Meta `path:"/coupon/list" tags:"优惠券分类" method:"get" summary:"列表优惠券种类接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type CouponListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CouponAllListReq struct {
	g.Meta `path:"/coupon/list/all" tags:"优惠券分类" method:"get" summary:"列表全部优惠券种类接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
}

type CouponAllListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Total int         `json:"total" description:"数据总数"`
}
