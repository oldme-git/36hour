package lib

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
	service.RegisterLib(New())
}

type sLib struct {
}

func New() *sLib {
	return &sLib{}
}

func (s *sLib) Create(ctx context.Context, lib *entity.Lib) (id int, err error) {
	res, err := dao.Lib.Ctx(ctx).Data(do.Lib{
		LibName: lib.LibName,
		Address: lib.Address,
		Active:  lib.Active,
	}).Insert()
	if err != nil {
		return 0, err
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func (s *sLib) GetOne(ctx context.Context, id int) (lib *entity.Lib, err error) {
	lib = new(entity.Lib)
	err = dao.Lib.Ctx(ctx).Where("id", id).Scan(lib)
	if err != nil {
		return nil, err
	}
	return lib, nil
}

func (s *sLib) GetList(ctx context.Context, condition *dao.LibSearchCondition) (libs []*entity.Lib, err error) {
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
		return nil, err
	}
	return libs, nil
}

func (s *sLib) Update(ctx context.Context, lib *entity.Lib) (err error) {
	_, err = dao.Lib.Ctx(ctx).Data(do.Lib{
		LibName: lib.LibName,
		Address: lib.Address,
		Active:  lib.Active,
	}).Where("id", lib.Id).Update()
	return err
}

func (s *sLib) Delete(ctx context.Context, id int) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var err error

		_, err = tx.Ctx(ctx).Model(dao.Lib.Table()).Where("id", id).Delete()
		if err != nil {
			return err
		}

		_, err = tx.Ctx(ctx).Model(dao.Floor.Table()).Where("lib_id", id).Delete()
		if err != nil {
			return err
		}

		_, err = tx.Ctx(ctx).Model(dao.Hall.Table()).Where("lib_id", id).Delete()
		if err != nil {
			return err
		}

		return nil
	})
}
