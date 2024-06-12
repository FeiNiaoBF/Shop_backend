// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RefundInfo is the golang structure for table refund_info.
type RefundInfo struct {
	Id        int         `json:"id"         orm:"id"         description:"售后退款表"`
	Number    string      `json:"number"     orm:"number"     description:"售后订单号"`
	OrderId   int         `json:"order_id"   orm:"order_id"   description:"订单id"`
	GoodsId   int         `json:"goods_id"   orm:"goods_id"   description:"要售后的商品id"`
	Reason    string      `json:"reason"     orm:"reason"     description:"退款原因"`
	Status    int         `json:"status"     orm:"status"     description:"状态 1待处理 2同意退款 3拒绝退款"`
	UserId    int         `json:"user_id"    orm:"user_id"    description:"用户id"`
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:""`
}
