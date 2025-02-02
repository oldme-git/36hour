// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Floor is the golang structure for table floor.
type Floor struct {
	Id        int         `json:"id"        orm:"id"         ` //
	LibId     int         `json:"libId"     orm:"lib_id"     ` //
	FloorName string      `json:"floorName" orm:"floor_name" ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
}
