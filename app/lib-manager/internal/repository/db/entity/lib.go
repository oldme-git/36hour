// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Lib is the golang structure for table lib.
type Lib struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	LibName   string      `json:"libName"   orm:"lib_name"   description:""`
	Address   string      `json:"address"   orm:"address"    description:"地址"`
	Active    bool        `json:"active"    orm:"active"     description:"是否正在使用"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
