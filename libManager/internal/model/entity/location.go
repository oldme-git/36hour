// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Location is the golang structure for table location.
type Location struct {
	Id           int         `json:"id"           ` //
	LibId        int         `json:"libId"        ` //
	FloorId      int         `json:"floorId"      ` //
	LocationName string      `json:"locationName" ` //
	CreatedAt    *gtime.Time `json:"createdAt"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` //
}
