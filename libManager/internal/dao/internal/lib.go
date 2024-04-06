// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LibDao is the data access object for table lib.
type LibDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns LibColumns // columns contains all the column names of Table for convenient usage.
}

// LibColumns defines and stores column names for table lib.
type LibColumns struct {
	Id        string //
	LibName   string //
	Address   string // 地址
	Active    string // 是否正在使用
	CreatedAt string //
	UpdatedAt string //
}

// libColumns holds the columns for table lib.
var libColumns = LibColumns{
	Id:        "id",
	LibName:   "lib_name",
	Address:   "address",
	Active:    "active",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewLibDao creates and returns a new DAO object for table data access.
func NewLibDao() *LibDao {
	return &LibDao{
		group:   "default",
		table:   "lib",
		columns: libColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LibDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LibDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LibDao) Columns() LibColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LibDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LibDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LibDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
