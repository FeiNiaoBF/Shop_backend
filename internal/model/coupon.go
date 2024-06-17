package model

import "github.com/gogf/gf/v2/os/gtime"

// CouponBase 创建/修改内容基类
type CouponBase struct {
	Name       string `json:"name" v:"required#名称必填" dc:"名称"`
	Price      int    `json:"price" v:"required#优惠券金额必填" dc:"优惠券金额"`
	GoodsIds   string `json:"goods_ids" dc:"可用的商品id 多个够好分隔"`
	CategoryId uint   `json:"category_id"  dc:"可用的商品分类"`
}

// CouponCreateInput 创建内容
type CouponCreateInput struct {
	CouponBase
}

// CouponCreateOutput 创建内容返回结果
type CouponCreateOutput struct {
	CouponId int `json:"coupon_id"`
}

// CouponUpdateInput 修改内容
type CouponUpdateInput struct {
	Id uint
	CouponBase
}

// CouponUpdateOutput 修改内容
type CouponUpdateOutput struct {
	Id        uint
	UpdatedAt string `json:"updated_at"` // 修改时间
}

// CouponGetALLListInput 获取内容列表
type CouponGetALLListInput struct {
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CouponGetAllListOutput 查询列表结果
type CouponGetAllListOutput struct {
	List  []CouponListItem `json:"list" description:"列表"`
	Total int              `json:"total" description:"数据总数"`
}

// CouponGetListInput 获取内容列表
type CouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CouponGetListOutput 查询列表结果
type CouponGetListOutput struct {
	List  []CouponListItem `json:"list" description:"列表"`
	Page  int              `json:"page" description:"分页码"`
	Size  int              `json:"size" description:"分页数量"`
	Total int              `json:"total" description:"数据总数"`
}

// CouponSearchInput 通过 Search 的关系对象结构
type CouponSearchInput struct {
	Key      string // 关键字
	Type     string // 内容模型
	CouponId uint   // 栏目ID
	CouponGetListInput
}

// CouponSearchOutput 返回关系对象结构
type CouponSearchOutput struct {
	Stats map[string]int `json:"stats"` // 搜索统计
	CouponGetListOutput
}

// CouponListItem 单一对象
type CouponListItem struct {
	Id uint `json:"id"`
	CouponBase
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
