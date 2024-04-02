package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"user/internal/consts"
	"user/internal/model"
	"user/internal/packed"
	"user/internal/service"
)

type sAuth struct {
}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

type UserClaims struct {
	Id       model.Id
	Username string
	jwt.RegisteredClaims
}

func (s *sAuth) Login(ctx context.Context, Username, Password string) (token string, err error) {
	// 1. 根据用户名查询用户信息
	user, err := service.User().GetOneByUsername(ctx, Username)
	if err != nil {
		return "", err
	}

	// 2. 验证密码
	ok, err := service.User().CheckPassword(ctx, Password, user.Password)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", packed.Err.New(1001)
	}
	userClaims := &UserClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	tokenRaw := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	token, err = tokenRaw.SignedString(consts.JwtKey)
	if err != nil {
		return "", err
	}
	return token, err
}

func (s *sAuth) Logout(ctx context.Context) (err error) {
	return err
}

func (s *sAuth) Register(ctx context.Context) (err error) {
	return err
}

func (s *sAuth) ChangePassword(ctx context.Context) (err error) {
	return err
}

func (s *sAuth) ResetPassword(ctx context.Context) (err error) {
	return err
}

func (s *sAuth) GetUserInfo(ctx context.Context) (err error) {
	return err
}
