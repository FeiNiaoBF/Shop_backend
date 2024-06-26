package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goBack/internal/consts"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/model/do"
	"goBack/internal/service"
	"goBack/utility"

	"github.com/gogf/gf/v2/util/grand"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// 注册
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	//处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	//插入数据返回id
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

// 修改密码
func (*sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	//	验证密保问题
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))

	err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(&userInfo)
	if err != nil {
		return
	}
	// TODO：单纯的比较可能需要调整
	if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer {
		g.Dump("userInfo.SecretAnswer:", userInfo.SecretAnswer)
		g.Dump("in.SecretAnswer:", in.SecretAnswer)
		return out, errors.New(consts.ErrSecretAnswerMsg)
	}
	// 密保正确
	userSalt := grand.S(10)
	in.UserSalt = userSalt
	in.Password = utility.EncryptPassword(in.Password, userSalt)
	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	return model.UpdatePasswordOutput{Id: userId}, nil
}
