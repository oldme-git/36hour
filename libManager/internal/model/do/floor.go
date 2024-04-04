// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Floor is the golang structure of table floor for DAO operations like Where/Data.
type Floor struct {
	g.Meta    `orm:"table:floor, do:true"`
	Id        interface{} //
	LibId     interface{} //
	FloorName interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
