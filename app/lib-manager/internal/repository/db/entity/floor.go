// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Floor is the golang structure for table floor.
type Floor struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	LibId     int         `json:"libId"     orm:"lib_id"     description:""`
	FloorName string      `json:"floorName" orm:"floor_name" description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
