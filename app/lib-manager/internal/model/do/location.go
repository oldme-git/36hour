// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Location is the golang structure of table location for DAO operations like Where/Data.
type Location struct {
	g.Meta       `orm:"table:location, do:true"`
	Id           interface{} //
	LibId        interface{} //
	FloorId      interface{} //
	LocationName interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
