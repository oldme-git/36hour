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

// newId 使用雪花获取一个新的用户id
func (s *sUser) newId() model.Id {
	id := service.Snowflake().Get()
	return model.Id(id.Int64())
}

func (s *sUser) Cre(ctx context.Context, user *model.User) (id model.Id, err error) {
	id = s.newId()
	_, err = dao.UserMain.Ctx(ctx).Data(do.UserMain{
		Id:       id,
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
	}).Insert()
	if err != nil {
		return 0, err
	}
	return
}
