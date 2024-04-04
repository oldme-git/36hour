// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"libManager/internal/dao/internal"
)

// internalLibDao is internal type for wrapping internal DAO implements.
type internalLibDao = *internal.LibDao

// libDao is the data access object for table lib.
// You can define custom methods on it to extend its functionality as you wish.
type libDao struct {
	internalLibDao
}

var (
	// Lib is globally public accessible object for table lib operations.
	Lib = libDao{
		internal.NewLibDao(),
	}
)

// Fill with you ideas below.
