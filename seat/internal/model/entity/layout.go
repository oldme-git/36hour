// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Layout is the golang structure for table layout.
type Layout struct {
	Id         int         `json:"id"         ` //
	LocationId int         `json:"locationId" ` //
	PolicyCId  int         `json:"policyCId"  ` // 公共策略id，优先使用
	PolicyLId  int         `json:"policyLId"  ` // 属于自己的策略id，最后使用
	LayoutName string      `json:"layoutName" ` //
	Map        string      `json:"map"        ` // 布局信息
	Status     int         `json:"status"     ` // 是否正常启用，0启用 1不启用
	Sort       int         `json:"sort"       ` // 排序，越小越靠前
	Seats      int         `json:"seats"      ` // 座位总数
	CreatedAt  *gtime.Time `json:"createdAt"  ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` //
}
