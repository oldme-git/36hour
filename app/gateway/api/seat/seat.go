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
	GetLayout(ctx context.Context, req *v1.GetLayoutReq) (res *v1.GetLayoutRes, err error)
	GetLayoutRuntime(ctx context.Context, req *v1.GetLayoutRuntimeReq) (res *v1.GetLayoutRuntimeRes, err error)
	InitLayout(ctx context.Context, req *v1.InitLayoutReq) (res *v1.InitLayoutRes, err error)
	Reserve(ctx context.Context, req *v1.ReserveReq) (res *v1.ReserveRes, err error)
}
