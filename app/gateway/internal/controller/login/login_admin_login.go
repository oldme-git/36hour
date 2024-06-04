package login

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/api/login/admin"
	auth "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"
)

func (c *ControllerAdmin) Login(ctx context.Context, req *admin.LoginReq) (res *admin.LoginRes, err error) {
	var (
		conn   = svc_disc.UserClient(ctx)
		client = auth.NewAuthClient(conn)
	)

	user, err := client.Login(ctx, &auth.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &admin.LoginRes{
		Token: user.GetToken(),
	}, nil
}
