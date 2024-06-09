package model

import "github.com/gogf/gf/v2/os/gtime"

// AdminCreateUpdateBase 创建/修改内容基类
type AdminCreateUpdateBase struct {
	Name     string
	Password string
	RoleIds  string
	UserSalt string
	IsAdmin  int
}

// AdminCreateInput 创建内容
type AdminCreateInput struct {
	AdminCreateUpdateBase
}

// AdminCreateOutput 创建内容返回结果
type AdminCreateOutput struct {
	AdminId int `json:"admin_id"`
}

// AdminUpdateInput 修改内容
type AdminUpdateInput struct {
	AdminCreateUpdateBase
	Id uint
}

// AdminGetListInput 获取内容列表
type AdminGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// AdminGetListOutput 查询列表结果
type AdminGetListOutput struct {
	List  []AdminGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

type AdminGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	Name      string      `json:"name"`
	RoleIds   string      `json:"role_ids"`
	IsAdmin   uint        `json:"is_admin"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
	DeleteAt  *gtime.Time `json:"delete_at"`  // 删除时间
}

type AdminSearchOutputItem struct {
	AdminGetListOutputItem
}
