// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goBack/internal/model"
)

type (
	IUserCoupon interface {
		// Create 新添
		Create(ctx context.Context, in model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// Update 修改
		Update(ctx context.Context, in model.UserCouponUpdateInput) (err error)
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error)
		GetAllList(ctx context.Context, in model.UserCouponGetALLListInput) (out *model.UserCouponGetAllListOutput, err error)
	}
)

var (
	localUserCoupon IUserCoupon
)

func UserCoupon() IUserCoupon {
	if localUserCoupon == nil {
		panic("implement not found for interface IUserCoupon, forgot register?")
	}
	return localUserCoupon
}

func RegisterUserCoupon(i IUserCoupon) {
	localUserCoupon = i
}
