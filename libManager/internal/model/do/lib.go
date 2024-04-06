// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Lib is the golang structure of table lib for DAO operations like Where/Data.
type Lib struct {
	g.Meta    `orm:"table:lib, do:true"`
	Id        interface{} //
	LibName   interface{} //
	Address   interface{} // 地址
	Active    interface{} // 是否正在使用
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
