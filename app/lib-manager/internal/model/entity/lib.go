// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Lib is the golang structure for table lib.
type Lib struct {
	Id        int         `json:"id"        orm:"id"         ` //
	LibName   string      `json:"libName"   orm:"lib_name"   ` //
	Address   string      `json:"address"   orm:"address"    ` // 地址
	Active    bool        `json:"active"    orm:"active"     ` // 是否正在使用
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
}
