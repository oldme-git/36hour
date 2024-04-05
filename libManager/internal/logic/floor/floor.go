package floor

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"libManager/internal/dao"
	"libManager/internal/model/do"
	"libManager/internal/model/entity"
	"libManager/internal/service"
)

func init() {
	service.RegisterFloor(New())
}

type sFloor struct {
}

func New() *sFloor {
	return &sFloor{}
}

func (s *sFloor) Create(ctx context.Context, floor *entity.Floor) (id int, err error) {
	res, err := dao.Floor.Ctx(ctx).Data(do.Floor{
		LibId:     floor.LibId,
		FloorName: floor.FloorName,
	}).Insert()
	if err != nil {
		return 0, err
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func (s *sFloor) GetOne(ctx context.Context, id int) (floor *entity.Floor, err error) {
	floor = new(entity.Floor)
	err = dao.Floor.Ctx(ctx).Where("id", id).Scan(floor)
	if err != nil {
		return nil, err
	}
	return floor, nil
}

func (s *sFloor) GetList(ctx context.Context, condition *dao.FloorSearchCondition) (floors []*entity.Floor, err error) {
	if condition.Page <= 0 {
		condition.Page = 1
	}
	if condition.PageSize <= 0 {
		condition.PageSize = 20
	}
	floors = make([]*entity.Floor, condition.PageSize)
	db := dao.Floor.Ctx(ctx)
	if condition.LibId > 0 {
		db = db.Where("lib_id", condition.LibId)
	}

	err = db.Page(condition.Page, condition.PageSize).Scan(&floors)
	if err != nil {
		return nil, err
	}
	return floors, nil
}

func (s *sFloor) Update(ctx context.Context, floor *entity.Floor) (err error) {
	_, err = dao.Floor.Ctx(ctx).Data(do.Floor{
		LibId:     floor.LibId,
		FloorName: floor.FloorName,
	}).Where("id", floor.Id).Update()
	return err
}

func (s *sFloor) Delete(ctx context.Context, id int) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var err error

		_, err = tx.Ctx(ctx).Model(dao.Floor.Table()).Where("id", id).Delete()
		if err != nil {
			return err
		}

		_, err = tx.Ctx(ctx).Model(dao.Hall.Table()).Where("floor_id", id).Delete()
		if err != nil {
			return err
		}

		return nil
	})
}
