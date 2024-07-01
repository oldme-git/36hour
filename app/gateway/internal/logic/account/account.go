package account

import (
	"context"

	auth "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"
)

// GetInfo 获取用户信息
func GetInfo(ctx context.Context, token string) (*auth.GetUserInfoRes, error) {
	var (
		conn   = svc_disc.UserClientConn(ctx)
		client = auth.NewAuthClient(conn)
	)

	return client.GetUserInfo(ctx, &auth.GetUserInfoReq{
		Token: token,
	})
}
