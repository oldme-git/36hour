// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/oldme-git/36hour/app/seat/internal/model/policy"
)

type (
	IPolicy interface {
		// New 从 json 字符串创建一个 Policy 对象
		New(ctx context.Context, str string) (p *policy.Policy, err error)
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
