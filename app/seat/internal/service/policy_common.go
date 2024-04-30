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
	IPolicyCommon interface {
		Create(ctx context.Context, policyCommon *entity.PolicyCommon) (id int, err error)
		GetOne(ctx context.Context, id int) (policyCommon *entity.PolicyCommon, err error)
		GetList(ctx context.Context, condition *model.PolicyCommonSearchCondition) (policyCommons []*entity.PolicyCommon, err error)
		// GetTotal 获取PolicyPrepare总数
		GetTotal(ctx context.Context, condition *model.PolicyCommonSearchCondition) (total int, err error)
		Update(ctx context.Context, policyCommon *entity.PolicyCommon) (err error)
		Delete(ctx context.Context, id int) (err error)
	}
)

var (
	localPolicyCommon IPolicyCommon
)

func PolicyCommon() IPolicyCommon {
	if localPolicyCommon == nil {
		panic("implement not found for interface IPolicyCommon, forgot register?")
	}
	return localPolicyCommon
}

func RegisterPolicyCommon(i IPolicyCommon) {
	localPolicyCommon = i
}