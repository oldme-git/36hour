package hall

import (
	"context"

	"libManager/internal/dao"
	"libManager/internal/model/do"
	"libManager/internal/model/entity"
	"libManager/internal/service"
)

type sHall struct {
}

func init() {
	service.RegisterHall(New())
}

func New() *sHall {
	return &sHall{}
}

func (s *sHall) Create(ctx context.Context, hall *entity.Hall) (id int, err error) {
	res, err := dao.Hall.Ctx(ctx).Data(do.Hall{
		LibId:    hall.LibId,
		FloorId:  hall.FloorId,
		HallName: hall.HallName,
	}).Insert()
	if err != nil {
		return 0, err
	}
	id64, _ := res.LastInsertId()
	return int(id64), nil
}

func (s *sHall) GetOne(ctx context.Context, id int) (hall *entity.Hall, err error) {
	hall = new(entity.Hall)
	err = dao.Hall.Ctx(ctx).Where("id", id).Scan(hall)
	if err != nil {
		return nil, err
	}
	return hall, nil
}

func (s *sHall) GetList(ctx context.Context, page, pageSize int) (halls []*entity.Hall, err error) {
	halls = make([]*entity.Hall, pageSize)
	err = dao.Hall.Ctx(ctx).Page(page, pageSize).Scan(&halls)
	if err != nil {
		return nil, err
	}
	return halls, nil
}

func (s *sHall) Update(ctx context.Context, hall *entity.Hall) (err error) {
	_, err = dao.Hall.Ctx(ctx).Data(do.Hall{
		LibId:    hall.LibId,
		FloorId:  hall.FloorId,
		HallName: hall.HallName,
	}).Where("id", hall.Id).Update()
	return err
}

func (s *sHall) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.Hall.Ctx(ctx).Where("id", id).Delete()
	return err
}