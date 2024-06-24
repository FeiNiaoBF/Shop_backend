package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goBack/api/frontend"
	"goBack/internal/model"
	"goBack/internal/service"
)

var Praise = cPraise{}

type cPraise struct{}

func (a *cPraise) Add(ctx context.Context, req *frontend.PraiseAddReq) (res *frontend.PraiseAddRes, err error) {
	data := model.PraiseAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Praise().AddPraise(ctx, data)
	if err != nil {
		return nil, err
	}
	res = &frontend.PraiseAddRes{
		Id: out.Id,
	}
	return
}

func (a *cPraise) Delete(ctx context.Context, req *frontend.PraiseDeleteReq) (res *frontend.PraiseDeleteRes, err error) {
	data := model.PraiseDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return
	}
	Praise, err := service.Praise().DeletePraise(ctx, data)
	if err != nil {
		return
	}
	res = &frontend.PraiseDeleteRes{
		Id: Praise.Id,
	}
	return
}

func (a *cPraise) GetList(ctx context.Context, req *frontend.ListPraiseReq) (res *frontend.ListPraiseRes, err error) {
	list, err := service.Praise().GetList(ctx, model.PraiseGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return
	}
	res = &frontend.ListPraiseRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}
	return
}
