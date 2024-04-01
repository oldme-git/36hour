package model

import "github.com/gogf/gf/v2/os/gtime"

type Id int64

type User struct {
	Id        Id          `json:"id"`
	Username  string      `json:"username"`
	Password  string      `json:"password"`
	Phone     string      `json:"phone"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}
