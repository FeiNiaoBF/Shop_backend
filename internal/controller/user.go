package controller

import (
	"context"
	"goBack/api/frontend"
	"goBack/internal/consts"
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

func (c *cUser) Info(ctx context.Context, req *frontend.InfoReq) (res *frontend.InfoRes, err error) {
	//res = &frontend.InfoRes{}
	res.Id = gconv.Uint(ctx.Value(consts.CtxUserId))
	res.Name = gconv.String(ctx.Value(consts.CtxUserName))
	res.Avatar = gconv.String(ctx.Value(consts.CtxUserAvatar))
	res.Sex = gconv.Uint8(ctx.Value(consts.CtxUserSex))
	res.Sign = gconv.String(ctx.Value(consts.CtxUserSign))
	res.Status = gconv.Uint8(ctx.Value(consts.CtxUserStatus))
	return res, nil
}
