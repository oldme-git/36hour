package policy_common_test

import (
	"database/sql"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/seat/internal/dao"
	"github.com/oldme-git/36hour/app/seat/internal/imodel"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/policy"
	"github.com/oldme-git/36hour/app/seat/internal/logic/policy_common"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
)

func TestCRUD(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx        = gctx.New()
			policyData = policy.Policy{
				Black: policy.Black{
					BlackOpen: true,
				},
			}
			policyDataStr, _ = policyData.String()
			policyCommon     = new(entity.PolicyCommon)
			policyCommonIn   = &entity.PolicyCommon{
				Name: "policyCommonTest",
				Info: policyDataStr,
			}
		)

		dao.PolicyCommon.Ctx(ctx).Where("name", "policyCommonTest").Delete("id > 0")

		// Create
		id, err := policy_common.Create(ctx, policyCommonIn)
		t.AssertNil(err)

		// GetList
		condition := &imodel.PolicyCommonSearchCondition{
			Page:     1,
			PageSize: 1,
			Name:     "Common",
		}
		policyCommons, err := policy_common.GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(policyCommons), 1)

		// GetOne
		policyCommon, err = policy_common.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(policyCommon.Id, id)
		t.Assert(policyCommon.Name, policyCommonIn.Name)
		t.Assert(policyCommon.Info, policyCommonIn.Info)

		// Update
		var (
			policyDataUpt = policy.Policy{
				Black: policy.Black{
					BlackOpen: false,
				},
				Leave: policy.Leave{
					LeaveOpen: true,
				},
			}
			policyDataUptStr, _ = policyDataUpt.String()
			policyCommonUptIn   = &entity.PolicyCommon{
				Id:   id,
				Name: "policyCommonTestUpt",
				Info: policyDataUptStr,
			}
		)

		err = policy_common.Update(ctx, policyCommonUptIn)
		t.AssertNil(err)
		policyCommon, err = policy_common.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(policyCommon.Name, policyCommonUptIn.Name)
		t.Assert(policyCommon.Info, policyCommonUptIn.Info)

		// Delete
		err = policy_common.Delete(ctx, id)
		t.AssertNil(err)
		_, err = policy_common.GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
	})
}
