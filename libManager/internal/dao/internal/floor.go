// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FloorDao is the data access object for table floor.
type FloorDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns FloorColumns // columns contains all the column names of Table for convenient usage.
}

// FloorColumns defines and stores column names for table floor.
type FloorColumns struct {
	Id        string //
	LibId     string //
	FloorName string //
	CreatedAt string //
	UpdatedAt string //
}

// floorColumns holds the columns for table floor.
var floorColumns = FloorColumns{
	Id:        "id",
	LibId:     "lib_id",
	FloorName: "floor_name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewFloorDao creates and returns a new DAO object for table data access.
func NewFloorDao() *FloorDao {
	return &FloorDao{
		group:   "default",
		table:   "floor",
		columns: floorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FloorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FloorDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FloorDao) Columns() FloorColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FloorDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FloorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FloorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
