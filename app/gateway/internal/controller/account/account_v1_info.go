package account

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	auth "github.com/oldme-git/36hour/app/user/api/auth/v1"
	"github.com/oldme-git/36hour/utility/svc_disc"

	"github.com/oldme-git/36hour/app/gateway/api/account/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	var (
		conn   = svc_disc.UserClientConn(ctx)
		client = auth.NewAuthClient(conn)
	)

	info, err := client.GetUserInfo(ctx, &auth.GetUserInfoReq{
		Token: g.RequestFromCtx(ctx).Request.Header.Get("Authorization"),
	})
	if err != nil {
		return nil, err
	}

	return &v1.InfoRes{
		Id:       int(info.User.GetId()),
		Username: info.User.GetUsername(),
		Phone:    info.User.GetPhone(),
		Lib: &v1.Lib{
			LibId:   int(info.Lib.GetLibId()),
			LibName: info.Lib.GetLibName(),
		},
	}, nil
}
