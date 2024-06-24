package frontend

import "github.com/gogf/gf/v2/frame/g"

type PraiseAddReq struct {
	g.Meta   `path:"/praise/add" in:"post" method:"post" tags:"前台点赞" summary:"添加点赞接口"`
	UserId   uint  `json:"user_id"    dc:"用户id"`
	ObjectId uint  `json:"object_id" v:"required#点赞对象id必填" dc:"对象id"`
	Type     uint8 `json:"type" v:"in:1,2" dc:"点赞类型：1商品 2文章" ` //数据校验 范围约束
}

type PraiseAddRes struct {
	Id uint `json:"id"`
}

type PraiseDeleteReq struct {
	g.Meta   `path:"/praise/delete" method:"post" tags:"前台点赞" sm:"移除点赞接口"`
	Id       uint  `json:"id"`
	Type     uint8 `json:"type"`
	ObjectId uint  `json:"object_id"`
}

type PraiseDeleteRes struct {
	Id uint `json:"id"`
}

type ListPraiseReq struct {
	g.Meta `path:"/praise/list" method:"post" tags:"前台点赞" sm:"前台点赞列表"`
	CommonPaginationReq
}

type ListPraiseRes struct {
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}

type ListPraiseItem struct {
	Id       int         `json:"id"        description:""`
	UserId   int         `json:"user_id"    description:"用户id"`
	ObjectId int         `json:"object_id"  description:"对象id"`
	Type     int         `json:"type"      description:"收藏类型：1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
