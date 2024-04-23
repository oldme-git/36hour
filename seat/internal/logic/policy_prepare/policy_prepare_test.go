package policy_prepare

import (
	"database/sql"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"seat/internal/dao"
	"seat/internal/model"
	"seat/internal/model/entity"
	"seat/internal/model/policy"
	"seat/internal/service"
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
		id, err := service.PolicyPrepare().Create(ctx, policyPrepareIn)
		t.AssertNil(err)

		// GetList
		condition := &model.PolicyPrepareSearchCondition{
			Page:     1,
			PageSize: 1,
			Name:     "Prepare",
		}
		policyPrepares, err := service.PolicyPrepare().GetList(ctx, condition)
		t.AssertNil(err)
		t.Assert(len(policyPrepares), 1)

		// GetOne
		policyPrepare, err = service.PolicyPrepare().GetOne(ctx, id)
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

		err = service.PolicyPrepare().Update(ctx, policyPrepareUptIn)
		t.AssertNil(err)
		policyPrepare, err = service.PolicyPrepare().GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(policyPrepare.Name, policyPrepareUptIn.Name)
		t.Assert(policyPrepare.Info, policyPrepareUptIn.Info)

		// Delete
		err = service.PolicyPrepare().Delete(ctx, id)
		t.AssertNil(err)
		_, err = service.PolicyPrepare().GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
	})
}
