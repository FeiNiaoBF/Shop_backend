// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goBack/internal/dao/internal"
)

// internalCollectionInfoDao is internal type for wrapping internal DAO implements.
type internalCollectionInfoDao = *internal.CollectionInfoDao

// collectionInfoDao is the data access object for table collection_info.
// You can define custom methods on it to extend its functionality as you wish.
type collectionInfoDao struct {
	internalCollectionInfoDao
}

var (
	// CollectionInfo is globally public accessible object for table collection_info operations.
	CollectionInfo = collectionInfoDao{
		internal.NewCollectionInfoDao(),
	}
)

// Fill with you ideas below.
