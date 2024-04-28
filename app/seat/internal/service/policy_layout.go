// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
)

type (
	IPolicyLayout interface {
		Create(ctx context.Context, policyLayout *entity.PolicyLayout) (id int, err error)
		GetOne(ctx context.Context, id int) (policyLayout *entity.PolicyLayout, err error)
		Update(ctx context.Context, policyLayout *entity.PolicyLayout) (err error)
		Delete(ctx context.Context, id int) (err error)
	}
)

var (
	localPolicyLayout IPolicyLayout
)

func PolicyLayout() IPolicyLayout {
	if localPolicyLayout == nil {
		panic("implement not found for interface IPolicyLayout, forgot register?")
	}
	return localPolicyLayout
}

func RegisterPolicyLayout(i IPolicyLayout) {
	localPolicyLayout = i
}
