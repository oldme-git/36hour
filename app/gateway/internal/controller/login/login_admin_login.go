package login

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/api/login/v1"
	auth "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"
)

func (c *ControllerAdmin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
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

	return &v1.LoginRes{
		Token: user.GetToken(),
	}, nil
}
