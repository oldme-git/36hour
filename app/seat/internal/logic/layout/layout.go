package layout

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/36hour/app/seat/internal/dao"
	"github.com/oldme-git/36hour/app/seat/internal/model"
	"github.com/oldme-git/36hour/app/seat/internal/model/do"
	"github.com/oldme-git/36hour/app/seat/internal/model/entity"
	"github.com/oldme-git/36hour/app/seat/internal/service"
	"github.com/oldme-git/36hour/utility"
)

func init() {
	service.RegisterLayout(New())
}

type sLayout struct {
}

func New() *sLayout {
	return &sLayout{}
}

func (s *sLayout) Create(ctx context.Context, layout *entity.Layout) (id int, err error) {
	// 数据验证
	if err := s.hookValid(ctx, layout); err != nil {
		return 0, err
	}
	res, err := dao.Layout.Ctx(ctx).Data(do.Layout{
		LocationId: layout.LocationId,
		PolicyCId:  layout.PolicyCId,
		PolicyLId:  layout.PolicyLId,
		LayoutName: layout.LayoutName,
		Map:        layout.Map,
		Status:     layout.Status,
		Sort:       layout.Sort,
		Seats:      layout.Seats,
	}).Insert()
	if err != nil {
		return 0, err
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func (s *sLayout) GetOne(ctx context.Context, id int) (layout *entity.Layout, err error) {
	layout = new(entity.Layout)
	err = dao.Layout.Ctx(ctx).Where("id", id).Scan(layout)
	if err != nil {
		return nil, err
	}

	return layout, nil
}

func (s *sLayout) GetList(ctx context.Context, condition *model.LayoutSearchCondition) (layouts []*entity.Layout, err error) {
	if condition.Page <= 0 {
		condition.Page = 1
	}
	if condition.PageSize <= 0 {
		condition.PageSize = 20
	}
	layouts = make([]*entity.Layout, condition.PageSize)
	db := dao.Layout.Ctx(ctx)
	if condition.LayoutName != "" {
		db = db.WhereLike("layout_name", "%"+condition.LayoutName+"%")
	}
	if condition.Status > 0 {
		db = db.Where("status", condition.Status)
	}
	if condition.SeatsMin > 0 && condition.SeatsMax > 0 {
		db = db.WhereBetween("github.com/oldme-git/36hour/app/seats", condition.SeatsMin, condition.SeatsMax)
	}
	err = db.Page(condition.Page, condition.PageSize).Scan(&layouts)
	if err != nil {
		return nil, err
	}
	return layouts, nil
}

// GetTotal 获取 Layout 总数
func (s *sLayout) GetTotal(ctx context.Context, condition *model.LayoutSearchCondition) (total int, err error) {
	db := dao.Layout.Ctx(ctx)
	if condition.LayoutName != "" {
		db = db.WhereLike("layout_name", "%"+condition.LayoutName+"%")
	}
	if condition.Status > 0 {
		db = db.Where("status", condition.Status)
	}
	if condition.SeatsMin > 0 && condition.SeatsMax > 0 {
		db = db.WhereBetween("github.com/oldme-git/36hour/app/seats", condition.SeatsMin, condition.SeatsMax)
	}
	total, err = db.Count()
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *sLayout) Update(ctx context.Context, layout *entity.Layout) (err error) {
	// 数据验证
	if err := s.hookValid(ctx, layout); err != nil {
		return err
	}
	_, err = dao.Layout.Ctx(ctx).Data(do.Layout{
		LocationId: layout.LocationId,
		PolicyCId:  layout.PolicyCId,
		PolicyLId:  layout.PolicyLId,
		LayoutName: layout.LayoutName,
		Map:        layout.Map,
		Status:     layout.Status,
		Sort:       layout.Sort,
		Seats:      layout.Seats,
	}).Where("id", layout.Id).Update()
	return err
}

func (s *sLayout) Delete(ctx context.Context, id int) (err error) {
	// 查询私有策略id
	policyLId, err := dao.Layout.Ctx(ctx).Where("id", id).Value("policy_l_id")
	if err != nil {
		return err
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var err error

		_, err = dao.Layout.Ctx(ctx).Where("id", id).Delete()
		if err != nil {
			return err
		}

		if policyLId.Int() == 0 {
			return nil
		}
		// 删除私有策略
		_, err = dao.PolicyLayout.Ctx(ctx).Where("id", policyLId).Delete()
		if err != nil {
			return err
		}

		return nil
	})
}

// hookValid 入库前的数据验证钩子
func (s *sLayout) hookValid(ctx context.Context, layout *entity.Layout) error {
	// 判断 map 是否为合法的 json
	if !gjson.Valid(layout.Map) {
		return utility.Err.New(3001)
	}

	// 公共策略id和私有策略id不能同时为空
	if layout.PolicyCId == 0 && layout.PolicyLId == 0 {
		return utility.Err.New(3002)
	}

	// 计算座位数并注入
	seats, err := s.CalculateSeatsByJson(ctx, layout.Map)
	if err != nil {
		return err
	}
	layout.Seats = seats

	return nil
}
