package model

import "github.com/gogf/gf/v2/os/gtime"

// CategoryBase 创建/修改内容基类
type CategoryBase struct {
	ParentId uint   `json:"parent_id" dc:"种类所属大类Id"`
	Name     string `json:"name" dc:"种类名称"`
	PicURL   string `json:"pic_url" dc:"图片链接"`
	Level    uint   `json:"level" dc:"等级（默认为一级）"`
	Sort     int    `json:"sort" dc:"排序"`
}

// CategoryCreateInput 创建内容
type CategoryCreateInput struct {
	CategoryBase
}

// CategoryCreateOutput 创建内容返回结果
type CategoryCreateOutput struct {
	CategoryId int `json:"category_id"`
}

// CategoryUpdateInput 修改内容
type CategoryUpdateInput struct {
	Id uint
	CategoryBase
}

// CategoryUpdateOutput 修改内容
type CategoryUpdateOutput struct {
	Id        uint
	UpdatedAt string `json:"updated_at"` // 修改时间
}

// CategoryGetALLListInput 获取内容列表
type CategoryGetALLListInput struct {
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CategoryGetAllListOutput 查询列表结果
type CategoryGetAllListOutput struct {
	List  []CategoryListItem `json:"list" description:"列表"`
	Total int                `json:"total" description:"数据总数"`
}

// CategoryGetListInput 获取内容列表
type CategoryGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CategoryGetListOutput 查询列表结果
type CategoryGetListOutput struct {
	List  []CategoryListItem `json:"list" description:"列表"`
	Page  int                `json:"page" description:"分页码"`
	Size  int                `json:"size" description:"分页数量"`
	Total int                `json:"total" description:"数据总数"`
}

// CategorySearchInput 通过 Search 的关系对象结构
type CategorySearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	CategoryGetListInput
}

// CategorySearchOutput 返回关系对象结构
type CategorySearchOutput struct {
	Stats map[string]int `json:"stats"` // 搜索统计
	CategoryGetListOutput
}

// CategoryListItem 单一对象
type CategoryListItem struct {
	Id uint `json:"id"`
	CategoryBase
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
