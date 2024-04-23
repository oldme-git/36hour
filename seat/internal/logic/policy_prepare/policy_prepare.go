package policy_prepare

import (
	"context"

	"seat/internal/dao"
	"seat/internal/model"
	"seat/internal/model/do"
	"seat/internal/model/entity"
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

func (s *sPolicyPrepare) Create(ctx context.Context, policyPrepare *entity.PolicyPrepare) (id int, err error) {
	res, err := dao.PolicyPrepare.Ctx(ctx).Data(do.PolicyPrepare{
		Name: policyPrepare.Name,
		Info: policyPrepare.Info,
	}).Insert()
	if err != nil {
		return 0, err
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func (s *sPolicyPrepare) GetOne(ctx context.Context, id int) (policyPrepare *entity.PolicyPrepare, err error) {
	policyPrepare = new(entity.PolicyPrepare)
	err = dao.PolicyPrepare.Ctx(ctx).Where("id", id).Scan(policyPrepare)
	if err != nil {
		return nil, err
	}
	return policyPrepare, nil
}

func (s *sPolicyPrepare) GetList(ctx context.Context, condition *model.PolicyPrepareSearchCondition) (policyPrepares []*entity.PolicyPrepare, err error) {
	if condition.Page <= 0 {
		condition.Page = 1
	}
	if condition.PageSize <= 0 {
		condition.PageSize = 20
	}
	policyPrepares = make([]*entity.PolicyPrepare, condition.PageSize)
	db := dao.PolicyPrepare.Ctx(ctx)
	if condition.Name != "" {
		db = db.WhereLike("name", "%"+condition.Name+"%")
	}
	err = db.Page(condition.Page, condition.PageSize).Scan(&policyPrepares)
	if err != nil {
		return nil, err
	}
	return policyPrepares, err
}

// GetTotal 获取PolicyPrepare总数
func (s *sPolicyPrepare) GetTotal(ctx context.Context, condition *model.PolicyPrepareSearchCondition) (total int, err error) {
	db := dao.PolicyPrepare.Ctx(ctx)
	if condition.Name != "" {
		db = db.WhereLike("name", "%"+condition.Name+"%")
	}
	total, err = db.Count()
	return
}

func (s *sPolicyPrepare) Update(ctx context.Context, policyPrepare *entity.PolicyPrepare) (err error) {
	_, err = dao.PolicyPrepare.Ctx(ctx).Data(do.PolicyPrepare{
		Name: policyPrepare.Name,
		Info: policyPrepare.Info,
	}).Where("id", policyPrepare.Id).Update()
	return
}

func (s *sPolicyPrepare) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.PolicyPrepare.Ctx(ctx).Where("id", id).Delete()
	return err
}
