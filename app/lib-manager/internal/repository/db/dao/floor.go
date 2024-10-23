// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/oldme-git/36hour/app/lib-manager/internal/repository/db/dao/internal"
)

// internalFloorDao is internal type for wrapping internal DAO implements.
type internalFloorDao = *internal.FloorDao

// floorDao is the data access object for table floor.
// You can define custom methods on it to extend its functionality as you wish.
type floorDao struct {
	internalFloorDao
}

var (
	// Floor is globally public accessible object for table floor operations.
	Floor = floorDao{
		internal.NewFloorDao(),
	}
)

// Fill with you ideas below.
