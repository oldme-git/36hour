package floor

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/logic/lib"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/do"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
	"github.com/oldme-git/36hour/utility"
)

func Create(ctx context.Context, floor *entity.Floor) (id int, err error) {
	// 数据验证
	if err := hookValid(ctx, floor); err != nil {
		return 0, err
	}
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

func GetOne(ctx context.Context, id int) (floor *entity.Floor, err error) {
	floor = new(entity.Floor)
	err = dao.Floor.Ctx(ctx).Where("id", id).Scan(floor)
	if err != nil {
		return nil, err
	}
	return floor, nil
}

func GetList(ctx context.Context, condition *dao.FloorSearchCondition) (floors []*entity.Floor, err error) {
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

func Update(ctx context.Context, floor *entity.Floor) (err error) {
	// 数据验证
	if err := hookValid(ctx, floor); err != nil {
		return err
	}
	_, err = dao.Floor.Ctx(ctx).Data(do.Floor{
		FloorName: floor.FloorName,
	}).Where("id", floor.Id).Update()
	return err
}

func Delete(ctx context.Context, id int) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var err error

		_, err = tx.Ctx(ctx).Model(dao.Floor.Table()).Where("id", id).Delete()
		if err != nil {
			return err
		}

		_, err = tx.Ctx(ctx).Model(dao.Location.Table()).Where("floor_id", id).Delete()
		if err != nil {
			return err
		}

		return nil
	})
}

func Exist(ctx context.Context, id int) error {
	count, err := dao.Floor.Ctx(ctx).Where("id", id).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return utility.Err.New(2001)
	}
	return nil
}

// hookValid 入库前的数据验证钩子
func hookValid(ctx context.Context, floor *entity.Floor) error {
	// 判断图书馆是否存在
	if err := lib.Exist(ctx, floor.LibId); err != nil {
		return err
	}
	return nil
}
