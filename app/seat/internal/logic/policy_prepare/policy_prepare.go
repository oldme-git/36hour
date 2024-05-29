package policy_prepare

import (
	"context"

	"github.com/oldme-git/36hour/app/seat/internal/dao"
	"github.com/oldme-git/36hour/app/seat/internal/model"
	"github.com/oldme-git/36hour/app/seat/internal/model/do"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
	"github.com/oldme-git/36hour/app/seat/internal/model/policy"
	"github.com/oldme-git/36hour/utility"
)

func Create(ctx context.Context, policyPrepare *entity.PolicyPrepare) (id int, err error) {
	// 数据验证
	if err := hookValid(ctx, policyPrepare); err != nil {
		return 0, err
	}

	res, err := dao.PolicyPrepare.Ctx(ctx).Data(do.PolicyPrepare{
		Name: policyPrepare.Name,
		Info: policyPrepare.Info,
	}).Insert()
	if err != nil {
		return 0, utility.Err.NewSys(err)
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func GetOne(ctx context.Context, id int) (policyPrepare *entity.PolicyPrepare, err error) {
	policyPrepare = new(entity.PolicyPrepare)
	err = dao.PolicyPrepare.Ctx(ctx).Where("id", id).Scan(policyPrepare)
	if err != nil {
		return nil, utility.Err.NewSys(err)
	}
	return policyPrepare, nil
}

func GetList(ctx context.Context, condition *model.PolicyPrepareSearchCondition) (policyPrepares []*entity.PolicyPrepare, err error) {
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
		return nil, utility.Err.NewSys(err)
	}
	return policyPrepares, err
}

// GetTotal 获取PolicyPrepare总数
func GetTotal(ctx context.Context, condition *model.PolicyPrepareSearchCondition) (total int, err error) {
	db := dao.PolicyPrepare.Ctx(ctx)
	if condition.Name != "" {
		db = db.WhereLike("name", "%"+condition.Name+"%")
	}
	total, err = db.Count()
	if err != nil {
		return 0, utility.Err.NewSys(err)
	}
	return
}

func Update(ctx context.Context, policyPrepare *entity.PolicyPrepare) (err error) {
	// 数据验证
	if err := hookValid(ctx, policyPrepare); err != nil {
		return err
	}

	_, err = dao.PolicyPrepare.Ctx(ctx).Data(do.PolicyPrepare{
		Name: policyPrepare.Name,
		Info: policyPrepare.Info,
	}).Where("id", policyPrepare.Id).Update()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return
}

func Delete(ctx context.Context, id int) (err error) {
	_, err = dao.PolicyPrepare.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return err
}

// hookValid 入库前的数据验证钩子
func hookValid(ctx context.Context, policyPrepare *entity.PolicyPrepare) (err error) {
	// 格式化 json 数据
	if policyPrepare.Info != "" {
		policyPrepare.Info, err = policy.Format(policyPrepare.Info)
		if err != nil {
			return
		}
	}
	return
}
