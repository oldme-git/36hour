package policy_layout

import (
	"context"

	"github.com/oldme-git/36hour/app/seat/internal/dao"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/policy"
	"github.com/oldme-git/36hour/app/seat/internal/model/do"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
	"github.com/oldme-git/36hour/utility"
)

func Create(ctx context.Context, policyLayout *entity.PolicyLayout) (id int, err error) {
	// 数据验证
	if err := hookValid(ctx, policyLayout); err != nil {
		return 0, err
	}

	res, err := dao.PolicyLayout.Ctx(ctx).Data(do.PolicyLayout{
		Info: policyLayout.Info,
	}).Insert()
	if err != nil {
		return 0, utility.Err.NewSys(err)
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func GetOne(ctx context.Context, id int) (policyLayout *entity.PolicyLayout, err error) {
	policyLayout = new(entity.PolicyLayout)
	err = dao.PolicyLayout.Ctx(ctx).Where("id", id).Scan(policyLayout)
	if err != nil {
		return nil, utility.Err.NewSys(err)
	}
	return policyLayout, nil
}

func Update(ctx context.Context, policyLayout *entity.PolicyLayout) (err error) {
	// 数据验证
	if err := hookValid(ctx, policyLayout); err != nil {
		return err
	}

	_, err = dao.PolicyLayout.Ctx(ctx).Data(do.PolicyLayout{
		Info: policyLayout.Info,
	}).Where("id", policyLayout.Id).Update()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return
}

func Delete(ctx context.Context, id int) (err error) {
	_, err = dao.PolicyLayout.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return
}

// hookValid 入库前的数据验证钩子
func hookValid(ctx context.Context, policyLayout *entity.PolicyLayout) (err error) {
	// 格式化 json 数据
	if policyLayout.Info != "" {
		policyLayout.Info, err = policy.Format(policyLayout.Info)
		if err != nil {
			return
		}
	}
	return
}
