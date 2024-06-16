package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/internal/model"
	"goBack/internal/service"
)

// Category 角色管理
var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Create(ctx context.Context, req *backend.CategoryCreateReq) (res *backend.CategoryCreateRes, err error) {
	out, err := service.Category().Create(ctx, model.CategoryCreateInput{
		CategoryBase: model.CategoryBase{
			Name:     req.Name,
			ParentId: req.ParentId,
			PicURL:   req.PicURL,
			Level:    req.Level,
			Sort:     req.Sort,
		},
	})
	if err != nil {
		return
	}
	res = &backend.CategoryCreateRes{
		Id: out.CategoryId,
	}
	return
}

func (a *cCategory) Update(ctx context.Context, req *backend.CategoryUpdateReq) (res *backend.CategoryUpdateRes, err error) {
	err = service.Category().Update(ctx, model.CategoryUpdateInput{
		Id: uint(req.Id),
		CategoryBase: model.CategoryBase{
			Name:     req.Name,
			ParentId: req.ParentId,
			PicURL:   req.PicURL,
			Level:    req.Level,
			Sort:     req.Sort,
		},
	})
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
