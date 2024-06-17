package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/internal/model"
	"goBack/internal/service"
)

// Coupon 角色管理
var Coupon = cCoupon{}

type cCoupon struct{}

func (a *cCoupon) Create(ctx context.Context, req *backend.CouponCreateReq) (res *backend.CouponCreateRes, err error) {
	out, err := service.Coupon().Create(ctx, model.CouponCreateInput{
		CouponBase: model.CouponBase{
			Name:       req.Name,
			Price:      req.Price,
			GoodsIds:   req.GoodsIds,
			CategoryId: req.CategoryId,
		},
	})
	if err != nil {
		return
	}
	res = &backend.CouponCreateRes{
		Id: out.CouponId,
	}
	return
}

func (a *cCoupon) Update(ctx context.Context, req *backend.CouponUpdateReq) (res *backend.CouponUpdateRes, err error) {
	err = service.Coupon().Update(ctx, model.CouponUpdateInput{
		Id: uint(req.Id),
		CouponBase: model.CouponBase{
			Name:       req.Name,
			Price:      req.Price,
			GoodsIds:   req.GoodsIds,
			CategoryId: req.CategoryId,
		},
	})
	if err != nil {
		return
	}
	res = &backend.CouponUpdateRes{
		Id: req.Id,
	}
	return
}

func (a *cCoupon) Delete(ctx context.Context, req *backend.CouponDeleteReq) (res *backend.CouponDeleteRes, err error) {
	err = service.Coupon().Delete(ctx, uint(req.Id))
	if err != nil {
		return
	}
	return
}

func (a *cCoupon) GetList(ctx context.Context, req *backend.CouponListReq) (res *backend.CouponListRes, err error) {
	list, err := service.Coupon().GetList(ctx, model.CouponGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return
	}
	res = &backend.CouponListRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}
	return
}

func (a *cCoupon) GetAllList(ctx context.Context, req *backend.CouponAllListReq) (res *backend.CouponAllListRes, err error) {
	list, err := service.Coupon().GetList(ctx, model.CouponGetListInput{
		Sort: req.Sort,
	})
	if err != nil {
		return
	}
	res = &backend.CouponAllListRes{
		List:  list.List,
		Total: list.Total,
	}
	return
}