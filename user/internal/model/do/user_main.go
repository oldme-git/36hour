// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMain is the golang structure of table user_main for DAO operations like Where/Data.
type UserMain struct {
	g.Meta    `orm:"table:user_main, do:true"`
	Id        interface{} //
	Username  interface{} //
	Password  interface{} //
	Phone     interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
