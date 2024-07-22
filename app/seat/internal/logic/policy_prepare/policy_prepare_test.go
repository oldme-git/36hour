package policy_prepare_test

import (
	"database/sql"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/seat/internal/dao"
	"github.com/oldme-git/36hour/app/seat/internal/imodel"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/policy"
	"github.com/oldme-git/36hour/app/seat/internal/logic/policy_prepare"
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
			policyPrepare    = new(entity.PolicyPrepare)
			policyPrepareIn  = &entity.PolicyPrepare{
				Name: "policyPrepareTest",
				Info: policyDataStr,
			}
		)

		dao.PolicyPrepare.Ctx(ctx).Where("name", "policyPrepareTest").Delete("id > 0")

		// Create
		id, err := policy_prepare.Create(ctx, policyPrepareIn)
		t.AssertNil(err)

		// GetList
		condition := &imodel.PolicyPrepareSearchCondition{
			Page:     1,
			PageSize: 1,
			Name:     "Prepare",
		}
		policyPrepares, err := policy_prepare.GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(policyPrepares), 1)

		// GetOne
		policyPrepare, err = policy_prepare.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(policyPrepare.Id, id)
		t.Assert(policyPrepare.Name, policyPrepareIn.Name)
		t.Assert(policyPrepare.Info, policyPrepareIn.Info)

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
			policyPrepareUptIn  = &entity.PolicyPrepare{
				Id:   id,
				Name: "policyPrepareTestUpt",
				Info: policyDataUptStr,
			}
		)

		err = policy_prepare.Update(ctx, policyPrepareUptIn)
		t.AssertNil(err)
		policyPrepare, err = policy_prepare.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(policyPrepare.Name, policyPrepareUptIn.Name)
		t.Assert(policyPrepare.Info, policyPrepareUptIn.Info)

		// Delete
		err = policy_prepare.Delete(ctx, id)
		t.AssertNil(err)
		_, err = policy_prepare.GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
	})
}
