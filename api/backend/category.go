package backend

import "github.com/gogf/gf/v2/frame/g"

type CategoryBase struct {
	ParentId uint   `json:"parent_id" dc:"种类所属大类"`
	Name     string `json:"name" v:"required#种类名称不能为空" dc:"种类名称"`
	PicURL   string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Level    uint   `json:"level" dc:"等级（默认为一级）"`
	Sort     int    `json:"sort" dc:"排序"`
}

type CategoryCreateReq struct {
	g.Meta `path:"/category/add" tags:"商品分类" method:"post" summary:"添加商品种类接口"`
	CategoryBase
}

type CategoryCreateRes struct {
	Id int `json:"id"`
}

type CategoryUpdateReq struct {
	g.Meta `path:"/category/update" tags:"商品分类" method:"post" summary:"更新商品种类接口"`
	Id     int `json:"id" v:"min:1#请选择需要删除的商品分类" dc:"商品分类id"`
	CategoryBase
}

type CategoryUpdateRes struct {
	Id int `json:"id"`
}

type CategoryDeleteReq struct {
	g.Meta `path:"/category/delete" tags:"商品分类" method:"delete" summary:"删除商品种类接口"`
	Id     int `json:"id" v:"min:1#请选择需要删除的商品分类" dc:"商品分类id"`
}

type CategoryDeleteRes struct{}

type CategoryListReq struct {
	g.Meta `path:"/category/list" tags:"商品分类" method:"get" summary:"列表商品种类接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type CategoryListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CategoryAllListReq struct {
	g.Meta `path:"/category/list/all" tags:"商品分类" method:"get" summary:"列表全部商品种类接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
}

type CategoryAllListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Total int         `json:"total" description:"数据总数"`
}
