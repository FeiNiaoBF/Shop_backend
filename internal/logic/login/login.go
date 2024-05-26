package login

import (
	"context"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/model/entity"
	"goBack/internal/service"
	"goBack/utility"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gutil"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	// 验证账号密码是否正确
	// 模型(结构体)与数据集合
	adminInfo := entity.AdminInfo{}
	// 用dao实现CRUD
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	gutil.Dump("加密后密码：", utility.EncryptPassword(in.Password, adminInfo.UserSalt))
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("账号或者密码不正确")
	}
	// 设置 session ID
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线 for session
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}

// 注销
func (s *sLogin) Logout(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}
