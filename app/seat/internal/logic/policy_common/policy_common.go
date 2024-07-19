package policy_common

import (
	"context"

	"github.com/oldme-git/36hour/app/seat/internal/dao"
	"github.com/oldme-git/36hour/app/seat/internal/imodel"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/policy"
	"github.com/oldme-git/36hour/app/seat/internal/model/do"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
	"github.com/oldme-git/36hour/utility"
)

func Create(ctx context.Context, policyCommon *entity.PolicyCommon) (id int, err error) {
	// 数据验证
	if err := hookValid(ctx, policyCommon); err != nil {
		return 0, err
	}

	res, err := dao.PolicyCommon.Ctx(ctx).Data(do.PolicyCommon{
		Name: policyCommon.Name,
		Info: policyCommon.Info,
	}).Insert()
	if err != nil {
		return 0, utility.Err.NewSys(err)
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func GetOne(ctx context.Context, id int) (policyCommon *entity.PolicyCommon, err error) {
	policyCommon = new(entity.PolicyCommon)
	err = dao.PolicyCommon.Ctx(ctx).Where("id", id).Scan(policyCommon)
	if err != nil {
		return nil, utility.Err.NewSys(err)
	}
	return policyCommon, nil
}

func GetList(ctx context.Context, condition *imodel.PolicyCommonSearchCondition) (policyCommons []*entity.PolicyCommon, err error) {
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
		return nil, utility.Err.NewSys(err)
	}
	return policyCommons, err
}

// GetTotal 获取PolicyPrepare总数
func GetTotal(ctx context.Context, condition *imodel.PolicyCommonSearchCondition) (total int, err error) {
	db := dao.PolicyCommon.Ctx(ctx)
	if condition.Name != "" {
		db = db.WhereLike("name", "%"+condition.Name+"%")
	}
	total, err = db.Count()
	if err != nil {
		return 0, utility.Err.NewSys(err)
	}
	return
}

func Update(ctx context.Context, policyCommon *entity.PolicyCommon) (err error) {
	// 数据验证
	if err := hookValid(ctx, policyCommon); err != nil {
		return err
	}

	_, err = dao.PolicyCommon.Ctx(ctx).Data(do.PolicyCommon{
		Name: policyCommon.Name,
		Info: policyCommon.Info,
	}).Where("id", policyCommon.Id).Update()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return
}

func Delete(ctx context.Context, id int) (err error) {
	_, err = dao.PolicyCommon.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return err
}

// hookValid 入库前的数据验证钩子
func hookValid(ctx context.Context, policyCommon *entity.PolicyCommon) (err error) {
	// 格式化 json 数据
	if policyCommon.Info != "" {
		policyCommon.Info, err = policy.Format(policyCommon.Info)
		if err != nil {
			return
		}
	}
	return
}
