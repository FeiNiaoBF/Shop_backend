// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConsigneeInfo is the golang structure for table consignee_info.
type ConsigneeInfo struct {
	Id        int         `json:"id"         orm:"id"         description:"收货地址表"`
	UserId    int         `json:"user_id"    orm:"user_id"    description:""`
	IsDefault int         `json:"is_default" orm:"is_default" description:"默认地址1  非默认0"`
	Name      string      `json:"name"       orm:"name"       description:""`
	Phone     string      `json:"phone"      orm:"phone"      description:""`
	Province  string      `json:"province"   orm:"province"   description:""`
	City      string      `json:"city"       orm:"city"       description:""`
	Town      string      `json:"town"       orm:"town"       description:"县区"`
	Street    string      `json:"street"     orm:"street"     description:"街道乡镇"`
	Detail    string      `json:"detail"     orm:"detail"     description:"地址详情"`
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:""`
}
