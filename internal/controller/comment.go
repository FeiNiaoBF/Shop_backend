package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goBack/api/frontend"
	"goBack/internal/model"
	"goBack/internal/service"
)

var Comment = cComment{}

type cComment struct{}

func (a *cComment) Add(ctx context.Context, req *frontend.CommentAddReq) (res *frontend.CommentAddRes, err error) {
	data := model.CommentAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Comment().AddComment(ctx, data)
	if err != nil {
		return nil, err
	}
	res = &frontend.CommentAddRes{
		Id: out.Id,
	}
	return
}

func (a *cComment) Delete(ctx context.Context, req *frontend.CommentDeleteReq) (res *frontend.CommentDeleteRes, err error) {
	data := model.CommentDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return
	}
	Comment, err := service.Comment().DeleteComment(ctx, data)
	if err != nil {
		return
	}
	res = &frontend.CommentDeleteRes{
		Id: Comment.Id,
	}
	return
}

func (a *cComment) GetList(ctx context.Context, req *frontend.ListCommentReq) (res *frontend.ListCommentRes, err error) {
	list, err := service.Comment().GetList(ctx, model.CommentGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return
	}
	res = &frontend.ListCommentRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}
	return
}
