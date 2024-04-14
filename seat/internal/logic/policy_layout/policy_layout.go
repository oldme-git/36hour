package policy_layout

import (
	"context"

	"seat/internal/dao"
	"seat/internal/model/do"
	"seat/internal/model/entity"
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
	_, err = dao.PolicyLayout.Ctx(ctx).Data(do.PolicyLayout{
		Info: policyLayout.Info,
	}).Where("id", policyLayout.Id).Update()
	return
}

func (s *sPolicyLayout) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.PolicyLayout.Ctx(ctx).Where("id", id).Delete()
	return
}
