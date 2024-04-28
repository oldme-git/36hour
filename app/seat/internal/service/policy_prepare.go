// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/oldme-git/36hour/app/seat/internal/model"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
)

type (
	IPolicyPrepare interface {
		Create(ctx context.Context, policyPrepare *entity.PolicyPrepare) (id int, err error)
		GetOne(ctx context.Context, id int) (policyPrepare *entity.PolicyPrepare, err error)
		GetList(ctx context.Context, condition *model.PolicyPrepareSearchCondition) (policyPrepares []*entity.PolicyPrepare, err error)
		// GetTotal 获取PolicyPrepare总数
		GetTotal(ctx context.Context, condition *model.PolicyPrepareSearchCondition) (total int, err error)
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
