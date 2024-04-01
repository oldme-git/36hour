package user

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
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

func (s *sUser) Create(ctx context.Context, user *model.User) (id model.Id, err error) {
	if user.Id == 0 {
		user.Id = s.newId()
	}

	// 密码加密
	user.Password, _ = gmd5.EncryptString(user.Password)

	_, err = dao.UserMain.Ctx(ctx).Data(do.UserMain{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
	}).Insert()
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (s *sUser) GetList(ctx context.Context, page int, pageSize int) (users []*model.User, err error) {
	users = make([]*model.User, 0)
	err = dao.UserMain.Ctx(ctx).Page(page, pageSize).Scan(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *sUser) GetOne(ctx context.Context, id model.Id) (user *model.User, err error) {
	user = new(model.User)
	err = dao.UserMain.Ctx(ctx).Where("id", id).Scan(user)
	if err != nil {
		return nil, err
	}

	return
}

func (s *sUser) Update(ctx context.Context, user *model.User) (err error) {
	// 密码加密
	user.Password, _ = gmd5.EncryptString(user.Password)

	_, err = dao.UserMain.Ctx(ctx).Data(do.UserMain{
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
	}).Where("id", user.Id).Update()
	return err
}

func (s *sUser) Delete(ctx context.Context, id model.Id) (err error) {
	_, err = dao.UserMain.Ctx(ctx).Where("id", id).Delete()
	return err
}
