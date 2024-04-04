package hall

import (
	"context"

	"libManager/internal/service"
)

func init() {
	service.RegisterHall(New())
}

type sLib struct {
}

func New() *sLib {
	return &sLib{}
}

func (s *sLib) Create(ctx context.Context) (err error) {
	return nil
}

func (s *sLib) GetOne(ctx context.Context) (err error) {
	return nil
}

func (s *sLib) GetList(ctx context.Context) (err error) {
	return nil
}

func (s *sLib) Update(ctx context.Context) (err error) {
	return nil
}

func (s *sLib) Delete(ctx context.Context) (err error) {
	return nil
}
