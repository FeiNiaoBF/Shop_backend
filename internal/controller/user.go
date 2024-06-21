package controller

import (
	"context"
	"goBack/api/frontend"
	"goBack/internal/model"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

var User = cUser{}

type cUser struct{}

func (a *cUser) Register(ctx context.Context, req *frontend.RegisterReq) (res *frontend.RegisterRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return
	}
	out, err := service.User().Register(ctx, data)
	if err != nil {
		return
	}
	res = &frontend.RegisterRes{
		Id: out.Id,
	}
	return
}
