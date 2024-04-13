package policy_prepare

import (
	"context"

	"seat/internal/model/entity"
	"seat/internal/model/policy"
	"seat/internal/service"
)

func init() {
	service.RegisterPolicyPrepare(New())
}

type sPolicyPrepare struct {
}

func New() *sPolicyPrepare {
	return &sPolicyPrepare{}
}

func (s *sPolicyPrepare) Create(ctx context.Context, name string, policy policy.Policy) (id int, err error) {
	return 0, nil
}

func (s *sPolicyPrepare) GetOne(ctx context.Context, id int) (policyPrepare *entity.PolicyPrepare, err error) {
	return nil, err
}

func (s *sPolicyPrepare) GetList(ctx context.Context) (libs []*entity.PolicyPrepare, err error) {
	return nil, err
}

func (s *sPolicyPrepare) Update(ctx context.Context, policyPrepare *entity.PolicyPrepare) (err error) {
	return nil
}

func (s *sPolicyPrepare) Delete(ctx context.Context, id int) (err error) {
	return nil
}
