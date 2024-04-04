// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Hall is the golang structure for table hall.
type Hall struct {
	Id        int         `json:"id"        ` //
	LibId     int         `json:"libId"     ` //
	FloorId   int         `json:"floorId"   ` //
	HallName  string      `json:"hallName"  ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
