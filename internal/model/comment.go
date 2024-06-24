package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"goBack/internal/model/do"
)

type CommentBase struct {
	UserId   uint   `json:"user_id"    dc:"用户id"`
	ObjectId uint   `json:"object_id"  dc:"对象id"`
	Type     uint8  `json:"type"      dc:"收藏类型：1商品 2文章"`
	ParentId uint   `json:"parent_id" dc:"父级评论id"`
	Content  string `json:"content" dc:"具体评论信息"`
}

type CommentAddInput struct {
	CommentBase
}

type CommentAddOutput struct {
	Id uint
}

type CommentDeleteInput struct {
	Id uint
	CommentBase
}

type CommentDeleteOutput struct {
	Id uint
}

// CommentGetALLListInput 获取内容列表
type CommentGetALLListInput struct {
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CommentGetAllListOutput 查询列表结果
type CommentGetAllListOutput struct {
	List  []CommentListItem `json:"list" description:"列表"`
	Total int               `json:"total" description:"数据总数"`
}

// CommentGetListInput 获取内容列表
type CommentGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Type uint8
}

// CommentGetListOutput 查询列表结果
type CommentGetListOutput struct {
	List  []CommentListItem `json:"list" description:"列表"`
	Page  int               `json:"page" description:"分页码"`
	Size  int               `json:"size" description:"分页数量"`
	Total int               `json:"total" description:"数据总数"`
}

// CommentListItem 单一对象
type CommentListItem struct {
	Id int `json:"id"`
	CommentBase
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article" orm:"with:id=object_id"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type CommentInfoBase struct {
	do.CommentInfo
	User UserInfoBase `json:"user" orm:"with:id=user_id"`
}
