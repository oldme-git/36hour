package account

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	auth "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"
)

// GetToken 获取用户token
func GetToken(ctx context.Context) string {
	return g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
}

// GetInfo 获取用户信息
func GetInfo(ctx context.Context, token string) (*auth.GetUserInfoRes, error) {
	var (
		conn   = svc_disc.UserClientConn(ctx)
		client = auth.NewAuthClient(conn)
	)

	ctx, cancel := context.WithTimeout(ctx, svc_disc.Timeout)
	defer cancel()

	return client.GetUserInfo(ctx, &auth.GetUserInfoReq{
		Token: token,
	})
}
