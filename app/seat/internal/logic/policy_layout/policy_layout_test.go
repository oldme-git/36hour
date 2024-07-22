package policy_layout_test

import (
	"database/sql"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/oldme-git/36hour/app/seat/internal/imodel/policy"
	"github.com/oldme-git/36hour/app/seat/internal/logic/policy_layout"
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
			policyLayout     = new(entity.PolicyLayout)
			policyLayoutIn   = &entity.PolicyLayout{
				Info: policyDataStr,
			}
		)

		// Create
		id, err := policy_layout.Create(ctx, policyLayoutIn)
		t.AssertNil(err)

		// GetOne
		policyLayout, err = policy_layout.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(policyLayout.Id, id)
		t.Assert(policyLayout.Info, policyLayoutIn.Info)

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
			policyLayoutUptIn   = &entity.PolicyLayout{
				Id:   id,
				Info: policyDataUptStr,
			}
		)

		err = policy_layout.Update(ctx, policyLayoutUptIn)
		t.AssertNil(err)
		policyLayout, err = policy_layout.GetOne(ctx, id)
		t.AssertNil(err)
		t.Assert(policyLayout.Info, policyLayoutUptIn.Info)

		// Delete
		err = policy_layout.Delete(ctx, id)
		t.AssertNil(err)
		_, err = policy_layout.GetOne(ctx, id)
		t.Assert(err, sql.ErrNoRows)
	})
}
