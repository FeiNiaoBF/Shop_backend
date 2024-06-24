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
	IPraise interface {
		AddPraise(ctx context.Context, in model.PraiseAddInput) (res *model.PraiseAddOutput, err error)
		// 兼容处理：优先根据收藏id删除，收藏id为0；再根据对象id和type删除
		DeletePraise(ctx context.Context, in model.PraiseDeleteInput) (res *model.PraiseDeleteOutput, err error)
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.PraiseGetListInput) (out *model.PraiseGetListOutput, err error)
	}
)

var (
	localPraise IPraise
)

func Praise() IPraise {
	if localPraise == nil {
		panic("implement not found for interface IPraise, forgot register?")
	}
	return localPraise
}

func RegisterPraise(i IPraise) {
	localPraise = i
}
