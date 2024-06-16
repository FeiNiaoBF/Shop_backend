package category

import (
	"context"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/model/entity"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

// Create 新添
func (s *sCategory) Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error) {
	insertID, err := dao.CategoryInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return
	}
	out.CategoryId = int(insertID)
	return out, nil
}

// Delete 删除
func (s *sCategory) Delete(ctx context.Context, id uint) error {
	return dao.CategoryInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.CategoryInfo.Ctx(ctx).Where(g.Map{
			dao.CategoryInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改
func (s *sCategory) Update(ctx context.Context, in model.CategoryUpdateInput) (err error) {
	_, err = dao.CategoryInfo.
		Ctx(ctx).Data(in).
		FieldsEx(dao.CategoryInfo.Columns().Id).
		Where(dao.CategoryInfo.Columns().Id, in.Id).
		Update()
	return
}

// GetList 查询内容列表
func (s *sCategory) GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	m := dao.CategoryInfo.Ctx(ctx)
	out = &model.CategoryGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Sort)
	// 执行查询
	var list []*entity.CategoryInfo
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

func (s *sCategory) GetAllList(ctx context.Context, in model.CategoryGetALLListInput) (out *model.CategoryGetAllListOutput, err error) {
	m := dao.CategoryInfo.Ctx(ctx)
	out = &model.CategoryGetAllListOutput{}

	// 分配查询
	listModel := m
	// 排序方式
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Sort)
	// 执行查询
	var list []*entity.CategoryInfo
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
