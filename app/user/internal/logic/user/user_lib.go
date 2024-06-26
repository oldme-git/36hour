package user

import (
	"context"

	"github.com/oldme-git/36hour/app/user/internal/dao"
	"github.com/oldme-git/36hour/app/user/internal/model"
	"github.com/oldme-git/36hour/app/user/internal/model/do"
	"github.com/oldme-git/36hour/app/user/internal/model/entity"
	"github.com/oldme-git/36hour/utility"
)

// BindLib 用户绑定图书馆
func BindLib(ctx context.Context, userId, libId model.Id) (err error) {
	_, err = dao.UserLib.Ctx(ctx).
		Where("user_id", userId).
		Data(&do.UserLib{
			UserId: userId,
			LibId:  libId,
		}).OnConflict(dao.UserLib.Columns().UserId).Save()
	if err != nil {
		err = utility.Err.NewSys(err)
	}
	return err
}

// GetUserLib 获取用户绑定的图书馆 Id
func GetUserLib(ctx context.Context, userId model.Id) (libId model.Id, err error) {
	userLib := new(entity.UserLib)
	err = dao.UserLib.Ctx(ctx).Where("user_id", userId).Scan(userLib)
	if err != nil {
		return 0, utility.Err.NewSys(err)
	}
	return model.Id(userLib.LibId), nil
}
