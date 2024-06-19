package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/internal/model"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// GoodsOptions 角色管理
var GoodsOptions = cGoodsOptions{}

type cGoodsOptions struct{}

func (a *cGoodsOptions) Create(ctx context.Context, req *backend.GoodsOptionsCreateReq) (res *backend.GoodsOptionsCreateRes, err error) {
	data := model.GoodsOptionsCreateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.GoodsOptions().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsOptionsCreateRes{Id: uint(out.GoodsOptionsId)}, nil
}

func (a *cGoodsOptions) Delete(ctx context.Context, req *backend.GoodsOptionsDeleteReq) (res *backend.GoodsOptionsDeleteRes, err error) {
	err = service.GoodsOptions().Delete(ctx, uint(req.Id))
	return
}

func (a *cGoodsOptions) Update(ctx context.Context, req *backend.GoodsOptionsUpdateReq) (res *backend.GoodsOptionsUpdateRes, err error) {
	data := model.GoodsOptionsUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.GoodsOptions().Update(ctx, data)
	return &backend.GoodsOptionsUpdateRes{Id: req.Id}, nil
}

func (a *cGoodsOptions) GetList(ctx context.Context, req *backend.GoodsOptionsListReq) (res *backend.GoodsOptionsListRes, err error) {
	list, err := service.GoodsOptions().GetList(ctx, model.GoodsOptionsGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return
	}
	res = &backend.GoodsOptionsListRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}
	return
}
