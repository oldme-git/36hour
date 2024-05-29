package layout

import (
	"context"

	"github.com/oldme-git/36hour/app/seat/internal/logic/policy_common"
	"github.com/oldme-git/36hour/app/seat/internal/logic/policy_layout"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
)

// GetRuntimePolicy 获取运行中的策略
// 如果有公共策略，获取公共策略信息
// 如果没有，则返回自己的策略信息
func GetRuntimePolicy(ctx context.Context, layout *entity.Layout) (policyInfo string, err error) {
	if layout.PolicyCId != 0 {
		policyCommon, err := policy_common.GetOne(ctx, layout.PolicyCId)
		if err != nil {
			return "", err
		}
		policyInfo = policyCommon.Info
	} else {
		policyLayout, err := policy_layout.GetOne(ctx, layout.PolicyLId)
		if err != nil {
			return "", err
		}
		policyInfo = policyLayout.Info
	}
	return

}
