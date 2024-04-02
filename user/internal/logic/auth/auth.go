package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"user/internal/consts"
	"user/internal/model"
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

func (s *sAuth) Login(ctx context.Context) (token string, err error) {
	userClaims := &UserClaims{
		Id:       1,
		Username: "admin",
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
