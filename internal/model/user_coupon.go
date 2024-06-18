package model

import "github.com/gogf/gf/v2/os/gtime"

// UserCouponBase 创建/修改内容基类
type UserCouponBase struct {
	UserId   uint  `json:"user_id" v:"required#用户id必填" dc:"用户id"`
	CouponId uint  `json:"coupon_id" v:"required#优惠券id必填" dc:"可用的商品分类"`
	Status   uint8 `json:"status" dc:"状态"`
}

// UserCouponCreateInput 创建内容
type UserCouponCreateInput struct {
	UserCouponBase
}

// UserCouponCreateOutput 创建内容返回结果
type UserCouponCreateOutput struct {
	UserCouponId int `json:"user_coupon_id"`
}

// UserCouponUpdateInput 修改内容
type UserCouponUpdateInput struct {
	Id uint
	UserCouponBase
}

// UserCouponUpdateOutput 修改内容
type UserCouponUpdateOutput struct {
	Id        uint
	UpdatedAt string `json:"updated_at"` // 修改时间
}

// UserCouponGetALLListInput 获取内容列表
type UserCouponGetALLListInput struct {
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// UserCouponGetAllListOutput 查询列表结果
type UserCouponGetAllListOutput struct {
	List  []UserCouponListItem `json:"list" description:"列表"`
	Total int                  `json:"total" description:"数据总数"`
}

// UserCouponGetListInput 获取内容列表
type UserCouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// UserCouponGetListOutput 查询列表结果
type UserCouponGetListOutput struct {
	List  []UserCouponListItem `json:"list" description:"列表"`
	Page  int                  `json:"page" description:"分页码"`
	Size  int                  `json:"size" description:"分页数量"`
	Total int                  `json:"total" description:"数据总数"`
}

// UserCouponSearchInput 通过 Search 的关系对象结构
type UserCouponSearchInput struct {
	Key          string // 关键字
	Type         string // 内容模型
	UserCouponId uint   // 栏目ID
	UserCouponGetListInput
}

// UserCouponSearchOutput 返回关系对象结构
type UserCouponSearchOutput struct {
	Stats map[string]int `json:"stats"` // 搜索统计
	UserCouponGetListOutput
}

// UserCouponListItem 单一对象
type UserCouponListItem struct {
	Id uint `json:"id"`
	UserCouponBase
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
