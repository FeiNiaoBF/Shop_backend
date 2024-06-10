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
	IPosition interface {
		Create(ctx context.Context, in model.PositionCreateInput) (out model.PositionCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// Update 修改
		Update(ctx context.Context, in model.PositionUpdateInput) error
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.PositionGetListInput) (out *model.PositionGetListOutput, err error)
	}
)

var (
	localPosition IPosition
)

func Position() IPosition {
	if localPosition == nil {
		panic("implement not found for interface IPosition, forgot register?")
	}
	return localPosition
}

func RegisterPosition(i IPosition) {
	localPosition = i
}
