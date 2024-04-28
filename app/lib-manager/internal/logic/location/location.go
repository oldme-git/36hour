package location

import (
	"context"

	"github.com/oldme-git/36hour/app/lib-manager/internal/dao"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/do"
	"github.com/oldme-git/36hour/app/lib-manager/internal/model/entity"
	"github.com/oldme-git/36hour/app/lib-manager/internal/service"
)

type sLocation struct {
}

func init() {
	service.RegisterLocation(New())
}

func New() *sLocation {
	return &sLocation{}
}

func (s *sLocation) Create(ctx context.Context, location *entity.Location) (id int, err error) {
	// 数据验证
	if err := s.hookValid(ctx, location); err != nil {
		return 0, err
	}
	res, err := dao.Location.Ctx(ctx).Data(do.Location{
		LibId:        location.LibId,
		FloorId:      location.FloorId,
		LocationName: location.LocationName,
	}).Insert()
	if err != nil {
		return 0, err
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func (s *sLocation) GetOne(ctx context.Context, id int) (location *entity.Location, err error) {
	location = new(entity.Location)
	err = dao.Location.Ctx(ctx).Where("id", id).Scan(location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (s *sLocation) GetList(ctx context.Context, condition *dao.LocationSearchCondition) (locations []*entity.Location, err error) {
	if condition.Page <= 0 {
		condition.Page = 1
	}
	if condition.PageSize <= 0 {
		condition.PageSize = 20
	}
	locations = make([]*entity.Location, condition.PageSize)
	db := dao.Location.Ctx(ctx)
	if condition.LibId > 0 {
		db = db.Where("lib_id", condition.LibId)
	}
	if condition.FloorId > 0 {
		db = db.Where("floor_id", condition.FloorId)
	}

	err = db.Page(condition.Page, condition.PageSize).Scan(&locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (s *sLocation) Update(ctx context.Context, location *entity.Location) (err error) {
	// 数据验证
	if err := s.hookValid(ctx, location); err != nil {
		return err
	}
	_, err = dao.Location.Ctx(ctx).Data(do.Location{
		FloorId:      location.FloorId,
		LocationName: location.LocationName,
	}).Where("id", location.Id).Update()
	return err
}

func (s *sLocation) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.Location.Ctx(ctx).Where("id", id).Delete()
	return err
}

// hookValid 入库前的数据验证钩子
func (s *sLocation) hookValid(ctx context.Context, location *entity.Location) error {
	// 判断图书馆是否存在
	if err := service.Lib().Exist(ctx, location.LibId); err != nil {
		return err
	}
	// 判断楼层是否存在
	if err := service.Floor().Exist(ctx, location.FloorId); err != nil {
		return err
	}
	return nil
}
