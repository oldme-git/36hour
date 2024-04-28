// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PolicyPrepare is the golang structure of table policy_prepare for DAO operations like Where/Data.
type PolicyPrepare struct {
	g.Meta `orm:"table:policy_prepare, do:true"`
	Id     interface{} //
	Name   interface{} //
	Info   interface{} // 策略信息
}
