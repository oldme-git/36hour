package policy_layout

import (
	"context"

	"seat/internal/dao"
	"seat/internal/model/do"
	"seat/internal/model/entity"
	"seat/internal/model/policy"
	"seat/internal/service"
)

func init() {
	service.RegisterPolicyLayout(New())
}

type sPolicyLayout struct {
}

func New() *sPolicyLayout {
	return &sPolicyLayout{}
}

func (s *sPolicyLayout) Create(ctx context.Context, policyLayout *entity.PolicyLayout) (id int, err error) {
	// 数据验证
	if err := s.hookValid(ctx, policyLayout); err != nil {
		return 0, err
	}

	res, err := dao.PolicyLayout.Ctx(ctx).Data(do.PolicyLayout{
		Info: policyLayout.Info,
	}).Insert()
	if err != nil {
		return 0, err
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func (s *sPolicyLayout) GetOne(ctx context.Context, id int) (policyLayout *entity.PolicyLayout, err error) {
	policyLayout = new(entity.PolicyLayout)
	err = dao.PolicyLayout.Ctx(ctx).Where("id", id).Scan(policyLayout)
	if err != nil {
		return nil, err
	}
	return policyLayout, nil
}

func (s *sPolicyLayout) Update(ctx context.Context, policyLayout *entity.PolicyLayout) (err error) {
	// 数据验证
	if err := s.hookValid(ctx, policyLayout); err != nil {
		return err
	}

	_, err = dao.PolicyLayout.Ctx(ctx).Data(do.PolicyLayout{
		Info: policyLayout.Info,
	}).Where("id", policyLayout.Id).Update()
	return
}

func (s *sPolicyLayout) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.PolicyLayout.Ctx(ctx).Where("id", id).Delete()
	return
}

// hookValid 入库前的数据验证钩子
func (s *sPolicyLayout) hookValid(ctx context.Context, policyLayout *entity.PolicyLayout) (err error) {
	// 格式化 json 数据
	if policyLayout.Info != "" {
		policyLayout.Info, err = policy.Format(policyLayout.Info)
		if err != nil {
			return
		}
	}
	return
}
