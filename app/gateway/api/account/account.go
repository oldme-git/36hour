// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/api/account/admin"
)

type IAccountAdmin interface {
	Info(ctx context.Context, req *admin.InfoReq) (res *admin.InfoRes, err error)
}
