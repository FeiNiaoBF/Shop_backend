package goods

import (
	"context"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sGoods struct{}

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

// Create 新添
func (s *sGoods) Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error) {
	insertID, err := dao.GoodsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return
	}
	out.GoodsId = int(insertID)
	return out, nil
}

// Delete 删除
func (s *sGoods) Delete(ctx context.Context, id uint) (err error) {
	_, err = dao.GoodsInfo.Ctx(ctx).Where(g.Map{
		dao.GoodsInfo.Columns().Id: id,
	}).Delete()
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sGoods) Update(ctx context.Context, in model.GoodsUpdateInput) error {
	_, err := dao.GoodsInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.GoodsInfo.Columns().Id).
		Where(dao.GoodsInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询内容列表
func (s *sGoods) GetList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error) {
	m := dao.GoodsInfo.Ctx(ctx)
	out = &model.GoodsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	out.List = make([]model.GoodsListItem, 0, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (s *sGoods) GetAllList(ctx context.Context, in model.GoodsGetALLListInput) (out *model.GoodsGetAllListOutput, err error) {
	m := dao.GoodsInfo.Ctx(ctx)
	out = &model.GoodsGetAllListOutput{}

	// 分配查询
	listModel := m
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	out.List = make([]model.GoodsListItem, 0, out.Total)
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// 商品详情
func (*sGoods) Detail(ctx context.Context, in model.GoodsDetailInput) (out model.GoodsDetailOutput, err error) {
	err = dao.GoodsInfo.Ctx(ctx).WithAll().WherePri(in.Id).Scan(&out)
	if err != nil {
		return model.GoodsDetailOutput{}, err
	}
	return
}
