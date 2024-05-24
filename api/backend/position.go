package backend

import "github.com/gogf/gf/v2/frame/g"

type PositionReq struct {
	g.Meta    `path:"/position/add" tags:"Position" method:"post" summary:"You first position api"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	GoodsId   uint   `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"`  //mysql三范式
	Sort      int    `json:"sort"    dc:"排序"`
}

type PositionRes struct {
	PositionId int `json:"position_id"`
}

// Delete
type PositionDeleteReq struct {
	g.Meta `path:"/Position/delete" method:"delete" tags:"手工位图" summary:"删除手工位图接口"`
	Id     uint `v:"min:1#请选择需要删除的手工位图" dc:"手工位图id"`
}

type PositionDeleteRes struct {
}

// Update
type PositionUpdateReq struct {
	g.Meta    `path:"/Position/update" method:"post" tags:"手工位图" summary:"修改手工位图接口"`
	Id        uint   `json:"id"      v:"min:1#请选择需要修改的手工位图" dc:"手工位图Id"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	GoodsName string `json:"goods_name" v:"required#商品名称不能为空" dc:"商品名称"` //冗余设计
	GoodsId   uint   `json:"goods_id"  v:"required#商品Id不能为空" dc:"商品ID"`  //mysql三范式
	Sort      int    `json:"sort"    dc:"排序"`
}

type PositionUpdateRes struct {
	Id uint `json:"id"`
}

// Get
// 分页查找
type PositionGetListCommonReq struct {
	g.Meta `path:"/Position/list" method:"get" tags:"手工位图" summary:"手工位图列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type PositionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
