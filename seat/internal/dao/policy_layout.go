// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"seat/internal/dao/internal"
)

// internalPolicyLayoutDao is internal type for wrapping internal DAO implements.
type internalPolicyLayoutDao = *internal.PolicyLayoutDao

// policyLayoutDao is the data access object for table policy_layout.
// You can define custom methods on it to extend its functionality as you wish.
type policyLayoutDao struct {
	internalPolicyLayoutDao
}

var (
	// PolicyLayout is globally public accessible object for table policy_layout operations.
	PolicyLayout = policyLayoutDao{
		internal.NewPolicyLayoutDao(),
	}
)

// Fill with you ideas below.

const (
	LayoutStatusAll = iota
	LayoutStatusEnable
	LayoutStatusDisable
)

type PolicyLayoutSearchCondition struct {
	Page       int
	PageSize   int
	LayoutName string
	Status     int
	// 座位数量区间
	SeatsMin int
	SeatsMax int
}
