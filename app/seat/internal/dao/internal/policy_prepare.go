// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PolicyPrepareDao is the data access object for table policy_prepare.
type PolicyPrepareDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns PolicyPrepareColumns // columns contains all the column names of Table for convenient usage.
}

// PolicyPrepareColumns defines and stores column names for table policy_prepare.
type PolicyPrepareColumns struct {
	Id   string //
	Name string //
	Info string // 策略信息
}

// policyPrepareColumns holds the columns for table policy_prepare.
var policyPrepareColumns = PolicyPrepareColumns{
	Id:   "id",
	Name: "name",
	Info: "info",
}

// NewPolicyPrepareDao creates and returns a new DAO object for table data access.
func NewPolicyPrepareDao() *PolicyPrepareDao {
	return &PolicyPrepareDao{
		group:   "default",
		table:   "policy_prepare",
		columns: policyPrepareColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PolicyPrepareDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PolicyPrepareDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PolicyPrepareDao) Columns() PolicyPrepareColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PolicyPrepareDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PolicyPrepareDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PolicyPrepareDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
