// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Floor is the golang structure for table floor.
type Floor struct {
	Id        int         `json:"id"        ` //
	LibId     int         `json:"libId"     ` //
	FloorName string      `json:"floorName" ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
