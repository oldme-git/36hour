// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserMainDao is the data access object for table user_main.
type UserMainDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserMainColumns // columns contains all the column names of Table for convenient usage.
}

// UserMainColumns defines and stores column names for table user_main.
type UserMainColumns struct {
	Id        string //
	Username  string //
	Password  string //
	Phone     string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// userMainColumns holds the columns for table user_main.
var userMainColumns = UserMainColumns{
	Id:        "id",
	Username:  "username",
	Password:  "password",
	Phone:     "phone",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewUserMainDao creates and returns a new DAO object for table data access.
func NewUserMainDao() *UserMainDao {
	return &UserMainDao{
		group:   "default",
		table:   "user_main",
		columns: userMainColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserMainDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserMainDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserMainDao) Columns() UserMainColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserMainDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserMainDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserMainDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
