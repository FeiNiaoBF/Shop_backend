package backend

import "github.com/gogf/gf/v2/frame/g"

type UserCouponBase struct {
	UserId   uint  `json:"user_id" v:"required#用户id必填" dc:"用户id"`
	CouponId uint  `json:"coupon_id" v:"required#优惠券id必填" dc:"可用的商品分类"`
	Status   uint8 `json:"status" dc:"状态"`
}

type UserCouponCreateReq struct {
	g.Meta `path:"/user/coupon/add" tags:"用户优惠券" method:"post" summary:"添加用户优惠券"`
	UserCouponBase
}

type UserCouponCreateRes struct {
	Id int `json:"id"`
}

type UserCouponUpdateReq struct {
	g.Meta `path:"/user/coupon/update" tags:"用户优惠券" method:"post" summary:"更新优惠券种类接口"`
	Id     int `json:"id" v:"min:1#请选择需要更新的优惠券" dc:"优惠券id"`
	UserCouponBase
}

type UserCouponUpdateRes struct {
	Id int `json:"id"`
}

type UserCouponDeleteReq struct {
	g.Meta `path:"/user/coupon/delete" tags:"用户优惠券" method:"delete" summary:"删除优惠券种类接口"`
	Id     int `json:"id" v:"min:1#请选择需要删除的优惠券" dc:"优惠券分类id"`
}

type UserCouponDeleteRes struct{}

type UserCouponListReq struct {
	g.Meta `path:"/user/coupon/list" tags:"用户优惠券" method:"get" summary:"列表优惠券种类接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type UserCouponListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type UserCouponAllListReq struct {
	g.Meta `path:"/user/coupon/list/all" tags:"用户优惠券" method:"get" summary:"列表全部优惠券种类接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
}

type UserCouponAllListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Total int         `json:"total" description:"数据总数"`
}
