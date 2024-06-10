package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/internal/service"
)

var Data = cData{}

type cData struct {
}

func (c *cData) DataHead(ctx context.Context, req *backend.DataHeadReq) (res *backend.DataHeadRes, err error) {
	dataHead, err := service.Data().GetData(ctx)
	if err != nil {
		return nil, err
	}
	//log.Println(dataHead)
	return &backend.DataHeadRes{
		TodayOrderCount: dataHead.TodayOrderCount,
		DAU:             dataHead.DAU,
		ConversionRate:  dataHead.ConversionRate,
	}, nil
}
