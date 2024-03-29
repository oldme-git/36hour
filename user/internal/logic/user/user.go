package user

import (
	"context"

	"user/internal/dao"
	"user/internal/model"
	"user/internal/model/do"
	"user/internal/service"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Cre(ctx context.Context, in *model.User) (id int, err error) {
	_, err = dao.UserMain.Ctx(ctx).Data(do.UserMain{
		Id:       nil,
		Username: in.Username,
		Password: in.Password,
		Phone:    in.Phone,
	}).Insert()
	if err != nil {
		return 0, err
	}
	return 1, nil
}
