package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/internal/model"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// Category 角色管理
var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Create(ctx context.Context, req *backend.CategoryCreateReq) (res *backend.CategoryCreateRes, err error) {
	data := model.CategoryCreateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Category().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	res = &backend.CategoryCreateRes{
		Id: out.CategoryId,
	}
	return
}

func (a *cCategory) Update(ctx context.Context, req *backend.CategoryUpdateReq) (res *backend.CategoryUpdateRes, err error) {
	data := model.CategoryUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}

	err = service.Category().Update(ctx, data)
	if err != nil {
		return
	}
	res = &backend.CategoryUpdateRes{
		Id: req.Id,
	}
	return
}

func (a *cCategory) Delete(ctx context.Context, req *backend.CategoryDeleteReq) (res *backend.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, uint(req.Id))
	if err != nil {
		return
	}
	return
}

func (a *cCategory) GetList(ctx context.Context, req *backend.CategoryListReq) (res *backend.CategoryListRes, err error) {
	list, err := service.Category().GetList(ctx, model.CategoryGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return
	}
	res = &backend.CategoryListRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}
	return
}

func (a *cCategory) GetAllList(ctx context.Context, req *backend.CategoryAllListReq) (res *backend.CategoryAllListRes, err error) {
	list, err := service.Category().GetList(ctx, model.CategoryGetListInput{
		Sort: req.Sort,
	})
	if err != nil {
		return
	}
	res = &backend.CategoryAllListRes{
		List:  list.List,
		Total: list.Total,
	}
	return
}
