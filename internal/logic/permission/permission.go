package permission

import (
	"context"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/model/entity"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	insertID, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return
	}
	out.PermissionId = int(insertID)
	return out, nil
}

// Delete 删除
func (s *sPermission) Delete(ctx context.Context, id uint) error {
	return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
			dao.PermissionInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 修改
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	_, err := dao.PermissionInfo.Ctx(ctx).Data(in).OmitEmpty().
		Where(dao.PermissionInfo.Columns().Id, in.Id).
		Update()
	if err != nil {
		return err
	}

	return nil
}

// GetList 查询内容列表
func (s *sPermission) GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	m := dao.PermissionInfo.Ctx(ctx)
	out = &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.PermissionInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 计数
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	//不指定item的键名用：Scan
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
