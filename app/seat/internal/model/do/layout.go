// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Layout is the golang structure of table layout for DAO operations like Where/Data.
type Layout struct {
	g.Meta     `orm:"table:layout, do:true"`
	Id         interface{} //
	LocationId interface{} //
	PolicyCId  interface{} // 公共策略id，优先使用
	PolicyLId  interface{} // 属于自己的策略id，最后使用
	LayoutName interface{} //
	Map        interface{} // 布局信息
	Status     interface{} // 是否正常启用，1启用 2不启用
	Sort       interface{} // 排序，越小越靠前
	Seats      interface{} // 座位总数
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
