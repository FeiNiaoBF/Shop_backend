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
	IAdmin interface {
		Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// Update 修改
		Update(ctx context.Context, in model.AdminUpdateInput) error
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error)
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
