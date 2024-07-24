// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SeatLog is the golang structure of table seat_log for DAO operations like Where/Data.
type SeatLog struct {
	g.Meta     `orm:"table:seat_log, do:true"`
	Id         interface{} //
	LocationId interface{} //
	LayoutId   interface{} //
	No         interface{} //
	Type       interface{} // 1预约 2签到 3退坐
	Uid        interface{} //
}
