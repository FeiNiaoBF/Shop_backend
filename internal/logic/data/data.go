package data

import (
	"context"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/service"
	"goBack/utility"

	"github.com/gogf/gf/v2/os/gtime"
)

type sData struct{}

func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) GetData(ctx context.Context) (out *model.DataHeadOutput, err error) {
	out = &model.DataHeadOutput{
		TodayOrderCount: getTodayOrder(ctx),
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(50),
	}
	return
}

func getTodayOrder(ctx context.Context) (count int) {
	todayStart := gtime.Now().StartOfDay()
	todayEnd := gtime.Now().EndOfDay()
	createdCol := dao.OrderInfo.Columns().CreatedAt
	count, err := dao.OrderInfo.Ctx(ctx).WhereBetween(createdCol, todayStart, todayEnd).
		Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return
}
