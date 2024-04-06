// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Lib is the golang structure for table lib.
type Lib struct {
	Id        int         `json:"id"        ` //
	LibName   string      `json:"libName"   ` //
	Address   string      `json:"address"   ` // 地址
	Active    bool        `json:"active"    ` // 是否正在使用
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
