package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/internal/model"
	"goBack/internal/service"
)

// UserCoupon 角色管理
var UserCoupon = cUserCoupon{}

type cUserCoupon struct{}

func (a *cUserCoupon) Create(ctx context.Context, req *backend.UserCouponCreateReq) (res *backend.UserCouponCreateRes, err error) {
	out, err := service.UserCoupon().Create(ctx, model.UserCouponCreateInput{
		UserCouponBase: model.UserCouponBase{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
		},
	})
	if err != nil {
		return
	}
	res = &backend.UserCouponCreateRes{
		Id: out.UserCouponId,
	}
	return
}

func (a *cUserCoupon) Update(ctx context.Context, req *backend.UserCouponUpdateReq) (res *backend.UserCouponUpdateRes, err error) {
	err = service.UserCoupon().Update(ctx, model.UserCouponUpdateInput{
		Id: uint(req.Id),
		UserCouponBase: model.UserCouponBase{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
		},
	})
	if err != nil {
		return
	}
	res = &backend.UserCouponUpdateRes{
		Id: req.Id,
	}
	return
}

func (a *cUserCoupon) Delete(ctx context.Context, req *backend.UserCouponDeleteReq) (res *backend.UserCouponDeleteRes, err error) {
	err = service.UserCoupon().Delete(ctx, uint(req.Id))
	if err != nil {
		return
	}
	return
}

func (a *cUserCoupon) GetList(ctx context.Context, req *backend.UserCouponListReq) (res *backend.UserCouponListRes, err error) {
	list, err := service.UserCoupon().GetList(ctx, model.UserCouponGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return
	}
	res = &backend.UserCouponListRes{
		List:  list.List,
		Page:  list.Page,
		Size:  list.Size,
		Total: list.Total,
	}
	return
}

func (a *cUserCoupon) GetAllList(ctx context.Context, req *backend.UserCouponAllListReq) (res *backend.UserCouponAllListRes, err error) {
	list, err := service.UserCoupon().GetList(ctx, model.UserCouponGetListInput{})
	if err != nil {
		return
	}
	res = &backend.UserCouponAllListRes{
		List:  list.List,
		Total: list.Total,
	}
	return
}
