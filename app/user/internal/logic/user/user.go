package user

import (
	"context"

	"github.com/oldme-git/36hour/app/user/internal/dao"
	"github.com/oldme-git/36hour/app/user/internal/model"
	"github.com/oldme-git/36hour/app/user/internal/model/do"
	"github.com/oldme-git/36hour/app/user/internal/service"
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

	user.Password, err = EncryptPwd(user.Password)
	if err != nil {
		return 0, err
	}

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

func (s *sUser) GetOneByUsername(ctx context.Context, username string) (user *model.User, err error) {
	user = new(model.User)
	err = dao.UserMain.Ctx(ctx).Where("github.com/oldme-git/36hour/app/username", username).Scan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update 不会修改密码
func (s *sUser) Update(ctx context.Context, user *model.User) (err error) {
	_, err = dao.UserMain.Ctx(ctx).Data(do.UserMain{
		Username: user.Username,
		Phone:    user.Phone,
	}).Where("id", user.Id).Update()
	return err
}

func (s *sUser) Delete(ctx context.Context, id model.Id) (err error) {
	_, err = dao.UserMain.Ctx(ctx).Where("id", id).Delete()
	return err
}

// CheckPassword 输入明文和密码，判断密码是否正确
func (s *sUser) CheckPassword(ctx context.Context, plainPwd, hashedPwd string) (bool, error) {
	encrypt, err := EncryptPwd(plainPwd)
	if err != nil {
		return false, err
	}
	if encrypt != hashedPwd {
		return false, nil
	}
	return true, nil
}
