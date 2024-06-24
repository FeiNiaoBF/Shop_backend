package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type PraiseBase struct {
	UserId   uint  `json:"user_id"    dc:"用户id"`
	ObjectId uint  `json:"object_id"  dc:"对象id"`
	Type     uint8 `json:"type"       dc:"收藏类型：1商品 2文章"`
}

type PraiseAddInput struct {
	PraiseBase
}

type PraiseAddOutput struct {
	Id uint
}

type PraiseDeleteInput struct {
	Id uint
	PraiseBase
}

type PraiseDeleteOutput struct {
	Id uint
}

// PraiseGetALLListInput 获取内容列表
type PraiseGetALLListInput struct {
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// PraiseGetAllListOutput 查询列表结果
type PraiseGetAllListOutput struct {
	List  []PraiseListItem `json:"list" description:"列表"`
	Total int              `json:"total" description:"数据总数"`
}

// PraiseGetListInput 获取内容列表
type PraiseGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Type uint8
}

// PraiseGetListOutput 查询列表结果
type PraiseGetListOutput struct {
	List  []PraiseListItem `json:"list" description:"列表"`
	Page  int              `json:"page" description:"分页码"`
	Size  int              `json:"size" description:"分页数量"`
	Total int              `json:"total" description:"数据总数"`
}

// PraiseListItem 单一对象
type PraiseListItem struct {
	Id int `json:"id"`
	PraiseBase
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article" orm:"with:id=object_id"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
