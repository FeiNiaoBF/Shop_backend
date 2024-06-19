package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/internal/model"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// Goods 角色管理
var Goods = cGoods{}

type cGoods struct{}

func (a *cGoods) Create(ctx context.Context, req *backend.GoodsCreateReq) (res *backend.GoodsCreateRes, err error) {
	data := model.GoodsCreateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Goods().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsCreateRes{Id: uint(out.GoodsId)}, nil
}

func (a *cGoods) Delete(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {
	err = service.Goods().Delete(ctx, uint(req.Id))
	return
}

func (a *cGoods) Update(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {
	data := model.GoodsUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.Goods().Update(ctx, data)
	return &backend.GoodsUpdateRes{Id: req.Id}, nil
}

func (a *cGoods) GetList(ctx context.Context, req *backend.GoodsListReq) (res *backend.GoodsListRes, err error) {
	list, err := service.Goods().GetList(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return
	}
	res = &backend.GoodsListRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}
	return
}
