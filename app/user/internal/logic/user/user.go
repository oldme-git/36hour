package user

import (
	"context"

	"github.com/oldme-git/36hour/app/user/internal/dao"
	"github.com/oldme-git/36hour/app/user/internal/logic/snowflake"
	"github.com/oldme-git/36hour/app/user/internal/model"
	"github.com/oldme-git/36hour/app/user/internal/model/do"
	"github.com/oldme-git/36hour/utility"
)

// newId 使用雪花获取一个新的用户id
func newId() model.Id {
	id := snowflake.Get()
	return model.Id(id.Int64())
}

func Create(ctx context.Context, user *model.User) (id model.Id, err error) {
	if user.Id == 0 {
		user.Id = newId()
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
		return 0, utility.Err.NewSys(err)
	}
	return user.Id, nil
}

func GetList(ctx context.Context, page int, pageSize int) (users []*model.User, err error) {
	users = make([]*model.User, 0)
	err = dao.UserMain.Ctx(ctx).Page(page, pageSize).Scan(&users)
	if err != nil {
		return nil, utility.Err.NewSys(err)
	}
	return users, nil
}

func GetOne(ctx context.Context, id model.Id) (user *model.User, err error) {
	user = new(model.User)
	err = dao.UserMain.Ctx(ctx).Where("id", id).Scan(user)
	if err != nil {
		return nil, utility.Err.NewSys(err)
	}

	return
}

func GetOneByUsername(ctx context.Context, username string) (user *model.User, err error) {
	user = new(model.User)
	data, err := dao.UserMain.Ctx(ctx).Where("username", username).One()
	if err != nil {
		return nil, utility.Err.NewSys(err)
	}
	err = data.Struct(user)
	if err != nil {
		return nil, utility.Err.New(1001)
	}
	return user, nil
}

// Update 不会修改密码
func Update(ctx context.Context, user *model.User) (err error) {
	_, err = dao.UserMain.Ctx(ctx).Data(do.UserMain{
		Username: user.Username,
		Phone:    user.Phone,
	}).Where("id", user.Id).Update()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return err
}

func Delete(ctx context.Context, id model.Id) (err error) {
	_, err = dao.UserMain.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return utility.Err.NewSys(err)
	}
	return err
}

// CheckPassword 输入明文和密码，判断密码是否正确
func CheckPassword(ctx context.Context, plainPwd, hashedPwd string) (bool, error) {
	encrypt, err := EncryptPwd(plainPwd)
	if err != nil {
		return false, utility.Err.NewSys(err)
	}
	if encrypt != hashedPwd {
		return false, nil
	}
	return true, nil
}
