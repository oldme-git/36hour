// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Location is the golang structure for table location.
type Location struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	LibId        int         `json:"libId"        orm:"lib_id"        description:""`
	FloorId      int         `json:"floorId"      orm:"floor_id"      description:""`
	LocationName string      `json:"locationName" orm:"location_name" description:""`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`
}
