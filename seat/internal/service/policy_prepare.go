// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"seat/internal/dao"
	"seat/internal/model/entity"
)

type (
	IPolicyPrepare interface {
		Create(ctx context.Context, policyPrepare *entity.PolicyPrepare) (id int, err error)
		GetOne(ctx context.Context, id int) (policyPrepare *entity.PolicyPrepare, err error)
		GetList(ctx context.Context, condition *dao.PolicyPrepareSearchCondition) (policyPrepares []*entity.PolicyPrepare, err error)
		Update(ctx context.Context, policyPrepare *entity.PolicyPrepare) (err error)
		Delete(ctx context.Context, id int) (err error)
	}
)

var (
	localPolicyPrepare IPolicyPrepare
)

func PolicyPrepare() IPolicyPrepare {
	if localPolicyPrepare == nil {
		panic("implement not found for interface IPolicyPrepare, forgot register?")
	}
	return localPolicyPrepare
}

func RegisterPolicyPrepare(i IPolicyPrepare) {
	localPolicyPrepare = i
}
