// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Layout is the golang structure for table layout.
type Layout struct {
	Id         int         `json:"id"         orm:"id"          ` //
	LocationId int         `json:"locationId" orm:"location_id" ` //
	PolicyCId  int         `json:"policyCId"  orm:"policy_c_id" ` // 公共策略id，优先使用
	PolicyLId  int         `json:"policyLId"  orm:"policy_l_id" ` // 属于自己的策略id，最后使用
	LayoutName string      `json:"layoutName" orm:"layout_name" ` //
	Map        string      `json:"map"        orm:"map"         ` // 布局信息
	Status     int         `json:"status"     orm:"status"      ` // 是否正常启用，1启用 2不启用
	Sort       int         `json:"sort"       orm:"sort"        ` // 排序，越小越靠前
	Seats      int         `json:"seats"      orm:"seats"       ` // 座位总数
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  ` //
}
