// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HallDao is the data access object for table hall.
type HallDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns HallColumns // columns contains all the column names of Table for convenient usage.
}

// HallColumns defines and stores column names for table hall.
type HallColumns struct {
	Id        string //
	LibId     string //
	FloorId   string //
	HallName  string //
	CreatedAt string //
	UpdatedAt string //
}

// hallColumns holds the columns for table hall.
var hallColumns = HallColumns{
	Id:        "id",
	LibId:     "lib_id",
	FloorId:   "floor_id",
	HallName:  "hall_name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewHallDao creates and returns a new DAO object for table data access.
func NewHallDao() *HallDao {
	return &HallDao{
		group:   "default",
		table:   "hall",
		columns: hallColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HallDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HallDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HallDao) Columns() HallColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HallDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HallDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HallDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
