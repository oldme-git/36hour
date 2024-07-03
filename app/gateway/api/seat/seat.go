// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package seat

import (
	"context"

	"github.com/oldme-git/36hour/app/gateway/api/seat/v1"
)

type ISeatV1 interface {
	GetLayouts(ctx context.Context, req *v1.GetLayoutsReq) (res *v1.GetLayoutsRes, err error)
}
