// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SeatLog is the golang structure for table seat_log.
type SeatLog struct {
	Id         int `json:"id"         orm:"id"          ` //
	LocationId int `json:"locationId" orm:"location_id" ` //
	LayoutId   int `json:"layoutId"   orm:"layout_id"   ` //
	No         int `json:"no"         orm:"no"          ` //
	Type       int `json:"type"       orm:"type"        ` // 1预约 2签到 3退坐
	Uid        int `json:"uid"        orm:"uid"         ` //
}
