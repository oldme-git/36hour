package lib

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/do"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
	"github.com/oldme-git/36hour/utility"
)

func Create(ctx context.Context, lib *entity.Lib) (id int, err error) {
	res, err := dao.Lib.Ctx(ctx).Data(do.Lib{
		LibName: lib.LibName,
		Address: lib.Address,
		Active:  lib.Active,
	}).Insert()
	if err != nil {
		return 0, utility.Err.NewSys(err)
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func GetOne(ctx context.Context, id int) (lib *entity.Lib, err error) {
	lib = new(entity.Lib)
	err = dao.Lib.Ctx(ctx).Where("id", id).Scan(lib)
	if err != nil {
		return nil, utility.Err.NewSys(err)
	}
	return lib, nil
}

func GetList(ctx context.Context, condition *dao.LibSearchCondition) (libs []*entity.Lib, err error) {
	if condition.Page <= 0 {
		condition.Page = 1
	}
	if condition.PageSize <= 0 {
		condition.PageSize = 20
	}
	libs = make([]*entity.Lib, condition.PageSize)
	db := dao.Lib.Ctx(ctx)
	if condition.LibName != "" {
		db = db.WhereLike("lib_name", "%"+condition.LibName+"%")
	}
	if condition.Address != "" {
		db = db.WhereLike("address", "%"+condition.Address+"%")
	}
	if condition.Active {
		db = db.Where("active", true)
	} else {
		db = db.Where("active", false)
	}

	err = db.Page(condition.Page, condition.PageSize).Scan(&libs)
	if err != nil {
		return nil, utility.Err.NewSys(err)
	}
	return libs, nil
}

func Update(ctx context.Context, lib *entity.Lib) (err error) {
	_, err = dao.Lib.Ctx(ctx).Data(do.Lib{
		LibName: lib.LibName,
		Address: lib.Address,
		Active:  lib.Active,
	}).Where("id", lib.Id).Update()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return err
}

func Delete(ctx context.Context, id int) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var err error

		_, err = tx.Ctx(ctx).Model(dao.Lib.Table()).Where("id", id).Delete()
		if err != nil {
			return utility.Err.NewSys(err)
		}

		_, err = tx.Ctx(ctx).Model(dao.Floor.Table()).Where("lib_id", id).Delete()
		if err != nil {
			return utility.Err.NewSys(err)
		}

		_, err = tx.Ctx(ctx).Model(dao.Location.Table()).Where("lib_id", id).Delete()
		if err != nil {
			return utility.Err.NewSys(err)
		}

		return nil
	})
}

func Exist(ctx context.Context, id int) error {
	count, err := dao.Lib.Ctx(ctx).Where("id", id).Count()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	if count == 0 {
		return utility.Err.New(2001)
	}
	return nil
}
