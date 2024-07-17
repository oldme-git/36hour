package login

import (
	"context"

	auth "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"

	"github.com/oldme-git/36hour/app/gateway/api/login/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	var (
		conn   = svc_disc.UserClientConn(ctx)
		client = auth.NewAuthClient(conn)
	)

	ctx, cancel := context.WithTimeout(ctx, svc_disc.Timeout)
	defer cancel()

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
