package account

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/api/account/v1"
	"github.com/oldme-git/36hour/app/gateway/internal/logic/account"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	token := account.GetToken(ctx)
	info, err := account.GetInfo(ctx, token)
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
