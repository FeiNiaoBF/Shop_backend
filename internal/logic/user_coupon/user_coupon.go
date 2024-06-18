package user_coupon

import (
	"context"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/model/entity"
	"goBack/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sUserCoupon struct{}

func init() {
	service.RegisterUserCoupon(New())
}

func New() *sUserCoupon {
	return &sUserCoupon{}
}

// Create 新添
func (s *sUserCoupon) Create(ctx context.Context, in model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error) {
	insertID, err := dao.UserCouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return
	}
	out.UserCouponId = int(insertID)
	return out, nil
}

// Delete 删除
func (s *sUserCoupon) Delete(ctx context.Context, id uint) error {
	return dao.UserCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.UserCouponInfo.Ctx(ctx).Where(g.Map{
			dao.UserCouponInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改
func (s *sUserCoupon) Update(ctx context.Context, in model.UserCouponUpdateInput) (err error) {
	_, err = dao.UserCouponInfo.
		Ctx(ctx).Data(in).
		FieldsEx(dao.UserCouponInfo.Columns().Id).
		Where(dao.UserCouponInfo.Columns().Id, in.Id).
		Update()
	return
}

// GetList 查询内容列表
func (s *sUserCoupon) GetList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error) {
	m := dao.UserCouponInfo.Ctx(ctx)
	out = &model.UserCouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.UserCouponInfo
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

func (s *sUserCoupon) GetAllList(ctx context.Context, in model.UserCouponGetALLListInput) (out *model.UserCouponGetAllListOutput, err error) {
	m := dao.UserCouponInfo.Ctx(ctx)
	out = &model.UserCouponGetAllListOutput{}

	// 分配查询
	listModel := m

	// 执行查询
	var list []*entity.UserCouponInfo
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
