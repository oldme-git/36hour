// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LayoutDao is the data access object for table layout.
type LayoutDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns LayoutColumns // columns contains all the column names of Table for convenient usage.
}

// LayoutColumns defines and stores column names for table layout.
type LayoutColumns struct {
	Id         string //
	LocationId string //
	PolicyCId  string // 公共策略id，优先使用
	PolicyLId  string // 属于自己的策略id，最后使用
	LayoutName string //
	Map        string // 布局信息
	Status     string // 是否正常启用，0启用 1不启用
	Sort       string // 排序，越小越靠前
	Seats      string // 座位总数
	CreatedAt  string //
	UpdatedAt  string //
}

// layoutColumns holds the columns for table layout.
var layoutColumns = LayoutColumns{
	Id:         "id",
	LocationId: "location_id",
	PolicyCId:  "policy_c_id",
	PolicyLId:  "policy_l_id",
	LayoutName: "layout_name",
	Map:        "map",
	Status:     "status",
	Sort:       "sort",
	Seats:      "seats",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewLayoutDao creates and returns a new DAO object for table data access.
func NewLayoutDao() *LayoutDao {
	return &LayoutDao{
		group:   "default",
		table:   "layout",
		columns: layoutColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LayoutDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LayoutDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LayoutDao) Columns() LayoutColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LayoutDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LayoutDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LayoutDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
