package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/internal/model"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// Coupon 角色管理
var Coupon = cCoupon{}

type cCoupon struct{}

func (a *cCoupon) Create(ctx context.Context, req *backend.CouponCreateReq) (res *backend.CouponCreateRes, err error) {
	data := model.CouponCreateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Coupon().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	res = &backend.CouponCreateRes{
		Id: out.CouponId,
	}
	return
}

func (a *cCoupon) Update(ctx context.Context, req *backend.CouponUpdateReq) (res *backend.CouponUpdateRes, err error) {
	data := model.CouponUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}

	err = service.Coupon().Update(ctx, data)
	if err != nil {
		return
	}
	res = &backend.CouponUpdateRes{
		Id: int(req.Id),
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
