package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/internal/consts"
	"goBack/internal/model"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// Article 角色管理
var Article = cArticle{}

type cArticle struct{}

func (a *cArticle) Create(ctx context.Context, req *backend.ArticleCreateReq) (res *backend.ArticleCreateRes, err error) {
	data := model.ArticleCreateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	out, err := service.Article().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	res = &backend.ArticleCreateRes{
		Id: int(out.Id),
	}
	return
}

func (a *cArticle) Update(ctx context.Context, req *backend.ArticleUpdateReq) (res *backend.ArticleUpdateRes, err error) {
	data := model.ArticleUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	err = service.Article().Update(ctx, data)
	if err != nil {
		return
	}
	res = &backend.ArticleUpdateRes{
		Id: int(req.Id),
	}
	return
}

func (a *cArticle) Delete(ctx context.Context, req *backend.ArticleDeleteReq) (res *backend.ArticleDeleteRes, err error) {
	err = service.Article().Delete(ctx, uint(req.Id))
	if err != nil {
		return
	}
	return
}

func (a *cArticle) GetList(ctx context.Context, req *backend.ArticleListReq) (res *backend.ArticleListRes, err error) {
	list, err := service.Article().GetList(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return
	}
	res = &backend.ArticleListRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}
	return
}

func (a *cArticle) GetAllList(ctx context.Context, req *backend.ArticleAllListReq) (res *backend.ArticleAllListRes, err error) {
	list, err := service.Article().GetList(ctx, model.ArticleGetListInput{
		Sort: req.Sort,
	})
	if err != nil {
		return
	}
	res = &backend.ArticleAllListRes{
		List:  list.List,
		Total: list.Total,
	}
	return
}
