package frontend

import "github.com/gogf/gf/v2/frame/g"

type CollectionAddReq struct {
	g.Meta   `path:"/collection/add" method:"post" tags:"收藏" sm:"添加收藏接口"`
	UserId   uint  `json:"user_id"    dc:"用户id"`
	ObjectId uint  `json:"object_id" v:"required#收藏对象id必填" dc:"对象id"`
	Type     uint8 `json:"type" v:"in:1,2" dc:"收藏类型：1商品 2文章" ` //数据校验 范围约束
}

type CollectionAddRes struct {
	Id uint `json:"id"`
}

type CollectionDeleteReq struct {
	g.Meta   `path:"/collection/delete" method:"post" tags:"收藏" sm:"移除收藏接口"`
	Id       uint  `json:"id"`
	Type     uint8 `json:"type"`
	ObjectId uint  `json:"object_id"`
}

type CollectionDeleteRes struct {
	Id uint `json:"id"`
}

type ListCollectionReq struct {
	g.Meta `path:"/collection/list" method:"post" tags:"收藏" sm:"收藏列表"`
	CommonPaginationReq
}

type ListCollectionRes struct {
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}

type ListCollectionItem struct {
	Id       int         `json:"id"        description:""`
	UserId   int         `json:"user_id"    description:"用户id"`
	ObjectId int         `json:"object_id"  description:"对象id"`
	Type     int         `json:"type"      description:"收藏类型：1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
