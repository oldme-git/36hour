// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package login

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/api/login/admin"
)

type ILoginAdmin interface {
	Login(ctx context.Context, req *admin.LoginReq) (res *admin.LoginRes, err error)
}
