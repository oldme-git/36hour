// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package seat

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/api/seat/v1"
)

type ISeatV1 interface {
	GetLayoutList(ctx context.Context, req *v1.GetLayoutListReq) (res *v1.GetLayoutListRes, err error)
}
