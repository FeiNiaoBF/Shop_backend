package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goBack/api/frontend"
	"goBack/internal/model"
	"goBack/internal/service"
)

var Collection = cCollection{}

type cCollection struct{}

func (a *cCollection) Create(ctx context.Context, req *frontend.AddCollectionReq) (res *frontend.AddCollectionRes, err error) {
	data := model.CollectionAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().AddCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	res = &frontend.AddCollectionRes{
		Id: out.Id,
	}
	return
}

func (a *cCollection) Delete(ctx context.Context, req *frontend.AddCollectionReq) (res *frontend.AddCollectionRes, err error) {
	data := model.CollectionDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return
	}
	collection, err := service.Collection().DeleteCollection(ctx, data)
	if err != nil {
		return
	}
	res = &frontend.AddCollectionRes{
		Id: collection.Id,
	}
	return
}

func (a *cCollection) GetList(ctx context.Context, req *frontend.ListCollectionReq) (res *frontend.ListCollectionRes, err error) {
	list, err := service.Collection().GetList(ctx, model.CollectionGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return
	}
	res = &frontend.ListCollectionRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}
	return
}
