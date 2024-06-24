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
	IGoods interface {
		// Create 新添
		Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) (err error)
		// Update 修改
		Update(ctx context.Context, in model.GoodsUpdateInput) error
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error)
		GetAllList(ctx context.Context, in model.GoodsGetALLListInput) (out *model.GoodsGetAllListOutput, err error)
		// 商品详情
		Detail(ctx context.Context, in model.GoodsDetailInput) (out model.GoodsDetailOutput, err error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
