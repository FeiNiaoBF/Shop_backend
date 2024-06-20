package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleBase struct {
	Title  string `json:"title"  dc:"文章标题" v:"required#名称必传"`
	Desc   string `json:"desc" dc:"文章概要"`
	PicUrl string `json:"pic_url"  dc:"图片"`
	// default value is 1
	IsAdmin uint   `d:"1"       dc:"1后台管理员发布 2前台用户发布"`
	Detail  string `json:"detail"  dc:"文章详情" v:"required#文章详情必填"`
	Praise  int    `json:"praise"  dc:"点赞数量"`
}

type ArticleCreateReq struct {
	g.Meta `path:"/article/add" method:"post" tags:"文章" summary:"添加文章接口"`
	ArticleBase
}

type ArticleCreateRes struct {
	Id int `json:"id"`
}

type ArticleUpdateReq struct {
	g.Meta `path:"/article/update" tags:"文章" method:"post" summary:"更新文章接口"`
	Id     uint `json:"id" v:"min:1#请选择需要更新的优惠券" dc:"优惠券id"`
	ArticleBase
}

type ArticleUpdateRes struct {
	Id int `json:"id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" tags:"文章" method:"delete" summary:"删除文章接口"`
	Id     int `json:"id" v:"min:1#请选择需要删除的优惠券" dc:"文章id"`
}

type ArticleDeleteRes struct{}

type ArticleListReq struct {
	g.Meta `path:"/article/list" tags:"文章" method:"get" summary:"列表文章接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type ArticleListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type ArticleAllListReq struct {
	g.Meta `path:"/article/list/all" tags:"文章" method:"get" summary:"列表全部文章接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
}

type ArticleAllListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Total int         `json:"total" description:"数据总数"`
}
