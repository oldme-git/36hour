package policy_common

import (
	"context"

	"seat/internal/dao"
	"seat/internal/model"
	"seat/internal/model/do"
	"seat/internal/model/entity"
	"seat/internal/model/policy"
	"seat/internal/service"
)

func init() {
	service.RegisterPolicyCommon(New())
}

type sPolicyCommon struct {
}

func New() *sPolicyCommon {
	return &sPolicyCommon{}
}

func (s *sPolicyCommon) Create(ctx context.Context, policyCommon *entity.PolicyCommon) (id int, err error) {
	// 数据验证
	if err := s.hookValid(ctx, policyCommon); err != nil {
		return 0, err
	}

	res, err := dao.PolicyCommon.Ctx(ctx).Data(do.PolicyCommon{
		Name: policyCommon.Name,
		Info: policyCommon.Info,
	}).Insert()
	if err != nil {
		return 0, err
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func (s *sPolicyCommon) GetOne(ctx context.Context, id int) (policyCommon *entity.PolicyCommon, err error) {
	policyCommon = new(entity.PolicyCommon)
	err = dao.PolicyCommon.Ctx(ctx).Where("id", id).Scan(policyCommon)
	if err != nil {
		return nil, err
	}
	return policyCommon, nil
}

func (s *sPolicyCommon) GetList(ctx context.Context, condition *model.PolicyCommonSearchCondition) (policyCommons []*entity.PolicyCommon, err error) {
	if condition.Page <= 0 {
		condition.Page = 1
	}
	if condition.PageSize <= 0 {
		condition.PageSize = 20
	}
	policyCommons = make([]*entity.PolicyCommon, condition.PageSize)
	db := dao.PolicyCommon.Ctx(ctx)
	if condition.Name != "" {
		db = db.WhereLike("name", "%"+condition.Name+"%")
	}
	err = db.Page(condition.Page, condition.PageSize).Scan(&policyCommons)
	if err != nil {
		return nil, err
	}
	return policyCommons, err
}

// GetTotal 获取PolicyPrepare总数
func (s *sPolicyCommon) GetTotal(ctx context.Context, condition *model.PolicyCommonSearchCondition) (total int, err error) {
	db := dao.PolicyCommon.Ctx(ctx)
	if condition.Name != "" {
		db = db.WhereLike("name", "%"+condition.Name+"%")
	}
	total, err = db.Count()
	return
}

func (s *sPolicyCommon) Update(ctx context.Context, policyCommon *entity.PolicyCommon) (err error) {
	// 数据验证
	if err := s.hookValid(ctx, policyCommon); err != nil {
		return err
	}

	_, err = dao.PolicyCommon.Ctx(ctx).Data(do.PolicyCommon{
		Name: policyCommon.Name,
		Info: policyCommon.Info,
	}).Where("id", policyCommon.Id).Update()
	return
}

func (s *sPolicyCommon) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.PolicyCommon.Ctx(ctx).Where("id", id).Delete()
	return err
}

// hookValid 入库前的数据验证钩子
func (s *sPolicyCommon) hookValid(ctx context.Context, policyCommon *entity.PolicyCommon) (err error) {
	// 格式化 json 数据
	if policyCommon.Info != "" {
		policyCommon.Info, err = policy.Format(policyCommon.Info)
		if err != nil {
			return
		}
	}
	return
}
