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
	IGoodsOptions interface {
		// Create 新添
		Create(ctx context.Context, in model.GoodsOptionsCreateInput) (out model.GoodsOptionsCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) (err error)
		// Update 修改
		Update(ctx context.Context, in model.GoodsOptionsUpdateInput) error
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error)
		GetAllList(ctx context.Context, in model.GoodsOptionsGetALLListInput) (out *model.GoodsOptionsGetAllListOutput, err error)
	}
)

var (
	localGoodsOptions IGoodsOptions
)

func GoodsOptions() IGoodsOptions {
	if localGoodsOptions == nil {
		panic("implement not found for interface IGoodsOptions, forgot register?")
	}
	return localGoodsOptions
}

func RegisterGoodsOptions(i IGoodsOptions) {
	localGoodsOptions = i
}
