// 生成测试数据的文件

package test

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gutil"
	"seat/internal/dao"
	"seat/internal/model/entity"
	"seat/internal/model/layout"
	"seat/internal/model/policy"
	"seat/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "seat/internal/logic/layout"
	_ "seat/internal/logic/policy_common"
	_ "seat/internal/logic/policy_layout"
	_ "seat/internal/logic/policy_prepare"
)

// 准备一个PolicyCommon
func TestPolicyCommon(t *testing.T) {
	var (
		ctx = gctx.New()
		p   = &policy.Policy{
			Black: policy.Black{
				BlackOpen: false,
			},
			Leave: policy.Leave{
				LeaveOpen: false,
			},
			Mark: policy.Mark{
				MarkOpen: false,
			},
			Renew: policy.Renew{
				RenewOpen: false,
			},
			Reserve: policy.Reserve{
				ReserveCloseMinute:  0,
				ReserveMinute:       0,
				ReserveNoticeMinute: 10,
			},
			Time: policy.Time{
				PolicyTimeConf: map[policy.Week]policy.TimeConf{
					policy.Sun: {
						OpenTime:  "08:00",
						CloseTime: "18:00",
					},
					policy.Mon: {
						OpenTime:  "08:00",
						CloseTime: "18:00",
					},
					policy.Tue: {
						OpenTime:  "08:00",
						CloseTime: "18:00",
					},
					policy.Wed: {
						OpenTime:  "08:00",
						CloseTime: "18:00",
					},
					policy.Thu: {
						OpenTime:  "08:00",
						CloseTime: "18:00",
					},
					policy.Fri: {
						OpenTime:  "08:00",
						CloseTime: "18:00",
					},
					policy.Sat: {
						OpenTime:  "08:00",
						CloseTime: "18:00",
					},
				},
				ReservePrepareMinute: 30,
				ReserveKeepMinute:    60,
			},
		}
		pStr, _      = p.String()
		policyCommon = &entity.PolicyCommon{
			Name: "policyC1",
			Info: pStr,
		}
	)
	_, err := service.PolicyCommon().Create(ctx, policyCommon)
	if err != nil {
		panic(err)
	}
}

// 准备一个Layout
func TestLayout(t *testing.T) {
	var (
		ctx = gctx.New()
		lay = &entity.Layout{
			LocationId: 1,
			PolicyCId:  1,
			LayoutName: "layout1",
			Status:     layout.StatusEnable,
			Sort:       1,
		}
	)

	// 获取commonId
	idValue, err := dao.PolicyCommon.Ctx(ctx).Where("name", "policyC1").Value("id")
	if err != nil {
		panic(err)
	}
	lay.PolicyCId = idValue.Int()

	// 获取map
	m := getLayoutMap(100)
	lay.Map, _ = service.Layout().LayoutCellsToJson(ctx, m)

	_, err = service.Layout().Create(ctx, lay)
	if err != nil {
		panic(err)
	}
}

// 准备一个LayoutMap
func getLayoutMap(num int) []layout.Cell {
	var (
		m  []layout.Cell
		no = 1
	)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			m = append(m, layout.Cell{
				X:    i,
				Y:    j,
				No:   no,
				Type: layout.CellSeat,
			})
			no++
		}
	}
	return m
}

func TestLayoutMap(t *testing.T) {
	var m = getLayoutMap(5)
	gutil.Dump(m)
}
