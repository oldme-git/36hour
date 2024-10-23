// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Location is the golang structure for table location.
type Location struct {
	Id           int         `json:"id"           orm:"id"            ` //
	LibId        int         `json:"libId"        orm:"lib_id"        ` //
	FloorId      int         `json:"floorId"      orm:"floor_id"      ` //
	LocationName string      `json:"locationName" orm:"location_name" ` //
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    ` //
}
