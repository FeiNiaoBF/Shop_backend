package coupon

import (
	"context"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/model/entity"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sCoupon struct{}

func init() {
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{}
}

// Create 新添
func (s *sCoupon) Create(ctx context.Context, in model.CouponCreateInput) (out model.CouponCreateOutput, err error) {
	insertID, err := dao.CouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return
	}
	out.CouponId = int(insertID)
	return out, nil
}

// Delete 删除
func (s *sCoupon) Delete(ctx context.Context, id uint) error {
	return dao.CouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.CouponInfo.Ctx(ctx).Where(g.Map{
			dao.CouponInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改
func (s *sCoupon) Update(ctx context.Context, in model.CouponUpdateInput) (err error) {
	_, err = dao.CouponInfo.
		Ctx(ctx).Data(in).
		FieldsEx(dao.CouponInfo.Columns().Id).
		Where(dao.CouponInfo.Columns().Id, in.Id).
		Update()
	return
}

// GetList 查询内容列表
func (s *sCoupon) GetList(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error) {
	m := dao.CouponInfo.Ctx(ctx)
	out = &model.CouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.CouponInfo.Columns().Price)
	// 执行查询
	var list []*entity.CouponInfo
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

func (s *sCoupon) GetAllList(ctx context.Context, in model.CouponGetALLListInput) (out *model.CouponGetAllListOutput, err error) {
	m := dao.CouponInfo.Ctx(ctx)
	out = &model.CouponGetAllListOutput{}

	// 分配查询
	listModel := m
	// 排序方式
	listModel = listModel.OrderDesc(dao.CouponInfo.Columns().Price)
	// 执行查询
	var list []*entity.CouponInfo
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
