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
	IComment interface {
		AddComment(ctx context.Context, in model.CommentAddInput) (res *model.CommentAddOutput, err error)
		// 兼容处理：优先根据收藏id删除，收藏id为0；再根据对象id和type删除
		DeleteComment(ctx context.Context, in model.CommentDeleteInput) (res *model.CommentDeleteOutput, err error)
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.CommentGetListInput) (out *model.CommentGetListOutput, err error)
	}
)

var (
	localComment IComment
)

func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

func RegisterComment(i IComment) {
	localComment = i
}
