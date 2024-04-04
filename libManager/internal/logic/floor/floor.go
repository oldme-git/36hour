package hall

import (
	"context"

	"libManager/internal/service"
)

func init() {
	service.RegisterHall(New())
}

type sFloor struct {
}

func New() *sFloor {
	return &sFloor{}
}

func (s *sFloor) Create(ctx context.Context) (err error) {
	return nil
}

func (s *sFloor) GetOne(ctx context.Context) (err error) {
	return nil
}

func (s *sFloor) GetList(ctx context.Context) (err error) {
	return nil
}

func (s *sFloor) Update(ctx context.Context) (err error) {
	return nil
}

func (s *sFloor) Delete(ctx context.Context) (err error) {
	return nil
}
