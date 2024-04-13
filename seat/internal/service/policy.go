// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"seat/internal/model/policy"
)

type (
	IPolicy interface {
		New(ctx context.Context, p *policy.Policy)
	}
)

var (
	localPolicy IPolicy
)

func Policy() IPolicy {
	if localPolicy == nil {
		panic("implement not found for interface IPolicy, forgot register?")
	}
	return localPolicy
}

func RegisterPolicy(i IPolicy) {
	localPolicy = i
}
