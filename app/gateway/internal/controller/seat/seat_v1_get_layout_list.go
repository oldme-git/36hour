package seat

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/oldme-git/36hour/app/gateway/api/seat/v1"
)

func (c *ControllerV1) GetLayoutList(ctx context.Context, req *v1.GetLayoutListReq) (res *v1.GetLayoutListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
