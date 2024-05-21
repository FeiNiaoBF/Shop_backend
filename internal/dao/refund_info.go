// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goBack/internal/dao/internal"
)

// internalRefundInfoDao is internal type for wrapping internal DAO implements.
type internalRefundInfoDao = *internal.RefundInfoDao

// refundInfoDao is the data access object for table refund_info.
// You can define custom methods on it to extend its functionality as you wish.
type refundInfoDao struct {
	internalRefundInfoDao
}

var (
	// RefundInfo is globally public accessible object for table refund_info operations.
	RefundInfo = refundInfoDao{
		internal.NewRefundInfoDao(),
	}
)

// Fill with you ideas below.
