// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CouponInfoDao is the data access object for table coupon_info.
type CouponInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CouponInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CouponInfoColumns defines and stores column names for table coupon_info.
type CouponInfoColumns struct {
	Id         string //
	Name       string //
	Price      string // 优惠前面值 单位分
	GoodsIds   string // 关联使用的goods_ids  逗号分隔
	CategoryId string // 关联使用的分类id
	CreatedAt  string //
	UpdatedAt  string //
	DeletedAt  string //
}

// couponInfoColumns holds the columns for table coupon_info.
var couponInfoColumns = CouponInfoColumns{
	Id:         "id",
	Name:       "name",
	Price:      "price",
	GoodsIds:   "goods_ids",
	CategoryId: "category_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewCouponInfoDao creates and returns a new DAO object for table data access.
func NewCouponInfoDao() *CouponInfoDao {
	return &CouponInfoDao{
		group:   "default",
		table:   "coupon_info",
		columns: couponInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CouponInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CouponInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CouponInfoDao) Columns() CouponInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CouponInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CouponInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CouponInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
