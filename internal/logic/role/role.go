package role

import (
	"context"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	insertID, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return
	}
	out.RoleId = int(insertID)
	return out, nil
}
