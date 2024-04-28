package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/oldme-git/36hour/app/user/internal/consts"
	"github.com/oldme-git/36hour/app/user/internal/model"
	"github.com/oldme-git/36hour/app/user/internal/service"
	"github.com/oldme-git/36hour/utility"
)

type sAuth struct {
}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

var jwtKey = consts.JwtKey

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
		return "", utility.Err.New(1001)
	}
	userClaims := &UserClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	}
	tokenRaw := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	token, err = tokenRaw.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return token, err
}

func (s *sAuth) Logout(ctx context.Context, id model.Id) (err error) {
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

func (s *sAuth) GetUserInfo(ctx context.Context, token string) (user *model.User, err error) {
	tokenClaims, _ := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	claims, ok := tokenClaims.Claims.(*UserClaims)
	if !(ok && tokenClaims.Valid) {
		err = utility.Err.New(1002)
		return
	}
	user = new(model.User)
	user, err = service.User().GetOne(ctx, claims.Id)
	if err != nil {
		return nil, err
	}
	return
}
