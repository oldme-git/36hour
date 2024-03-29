// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMain is the golang structure for table user_main.
type UserMain struct {
	Id        int64       `json:"id"        ` //
	Username  string      `json:"username"  ` //
	Password  string      `json:"password"  ` //
	Phone     string      `json:"phone"     ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
}
