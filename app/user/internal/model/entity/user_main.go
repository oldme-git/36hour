// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMain is the golang structure for table user_main.
type UserMain struct {
	Id        int64       `json:"id"        orm:"id"         ` //
	Username  string      `json:"username"  orm:"username"   ` //
	Password  string      `json:"password"  orm:"password"   ` //
	Phone     string      `json:"phone"     orm:"phone"      ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
}
