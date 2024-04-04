// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Hall is the golang structure of table hall for DAO operations like Where/Data.
type Hall struct {
	g.Meta    `orm:"table:hall, do:true"`
	Id        interface{} //
	LibId     interface{} //
	FloorId   interface{} //
	HallName  interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
