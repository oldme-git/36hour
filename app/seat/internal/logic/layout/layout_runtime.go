package layout

import (
	"context"

	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
	"github.com/oldme-git/36hour/app/seat/internal/service"
)

// GetRuntimePolicy 获取运行中的策略
// 如果有公共策略，获取公共策略信息
// 如果没有，则返回自己的策略信息
func (s *sLayout) GetRuntimePolicy(ctx context.Context, layout *entity.Layout) (policyInfo string, err error) {
	if layout.PolicyCId != 0 {
		policyCommon, err := service.PolicyCommon().GetOne(ctx, layout.PolicyCId)
		if err != nil {
			return "", err
		}
		policyInfo = policyCommon.Info
	} else {
		policyLayout, err := service.PolicyLayout().GetOne(ctx, layout.PolicyLId)
		if err != nil {
			return "", err
		}
		policyInfo = policyLayout.Info
	}
	return

}
