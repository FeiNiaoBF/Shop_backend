package model

// RoleCreateUpdateBase 创建/修改内容基类
type RoleCreateUpdateBase struct {
	Name string
	Desc string
}

// RoleCreateInput 创建内容
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// RoleCreateOutput 创建内容返回结果
type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}

// RoleUpdateInput 修改内容
type RoleUpdateInput struct {
	RoleCreateUpdateBase
	Id uint
}

// RoleGetListInput 获取内容列表
type RoleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}
