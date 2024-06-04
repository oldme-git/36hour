package account

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	auth "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"

	"github.com/oldme-git/36hour/app/gateway/api/account/admin"
)

func (c *ControllerAdmin) Info(ctx context.Context, req *admin.InfoReq) (res *admin.InfoRes, err error) {
	var (
		conn   = svc_disc.UserClient(ctx)
		client = auth.NewAuthClient(conn)
	)

	info, err := client.GetUserInfo(ctx, &auth.GetUserInfoReq{
		Token: g.RequestFromCtx(ctx).Request.Header.Get("Authorization"),
	})
	if err != nil {
		return nil, err
	}

	return &admin.InfoRes{
		Id:       int(info.User.Id),
		Username: info.User.Username,
		Phone:    info.User.Phone,
	}, nil
}
