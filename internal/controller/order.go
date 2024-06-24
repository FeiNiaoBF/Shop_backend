package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goBack/api/backend"
	"goBack/api/frontend"
	"goBack/internal/model"
	"goBack/internal/service"
)

var Order = cOrder{}

type cOrder struct{}

func (c *cOrder) Add(ctx context.Context, req *frontend.AddOrderReq) (res *frontend.AddOrderRes, err error) {
	orderAddInput := model.OrderAddInput{}
	if err = gconv.Scan(req, &orderAddInput); err != nil {
		return nil, err
	}

	addRes, err := service.Order().Add(ctx, orderAddInput)
	if err != nil {
		return nil, err
	}

	return &frontend.AddOrderRes{
		Id: addRes.Id,
	}, err
}

func (c *cOrder) List(ctx context.Context, req *backend.OrderListReq) (res *backend.OrderListRes, err error) {
	orderListInput := model.OrderListInput{}
	if err = gconv.Struct(req, &orderListInput); err != nil {
		return nil, err
	}

	orderListOutput, err := service.Order().GetList(ctx, orderListInput)
	if err != nil {
		return nil, err
	}

	return &backend.OrderListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  orderListOutput.List,
			Total: orderListOutput.Total,
			Page:  orderListOutput.Page,
			Size:  orderListOutput.Size,
		},
	}, err
}

func (c *cOrder) Detail(ctx context.Context, req *backend.OrderDetailReq) (res *backend.OrderDetailRes, err error) {
	detail, err := service.Order().Detail(ctx, model.OrderDetailInput{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	res = &backend.OrderDetailRes{}
	err = gconv.Struct(detail, res)

	return res, err
}
