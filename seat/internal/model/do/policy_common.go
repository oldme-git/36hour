// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PolicyCommon is the golang structure of table policy_common for DAO operations like Where/Data.
type PolicyCommon struct {
	g.Meta `orm:"table:policy_common, do:true"`
	Id     interface{} //
	Name   interface{} //
	Info   interface{} // 策略信息
}
