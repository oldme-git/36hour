package svc_disc

import (
	"context"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

var Timeout = 5 * time.Second

// ClientConn 获取一个 grpc 客户端
func ClientConn(ctx context.Context, service string, opts ...grpc.DialOption) *grpc.ClientConn {
	// 获取 grpc 服务地址，如果是开发环境，则使用本地地址
	serviceNameOrAddress := getAddress(ctx, service)
	return grpcx.Client.MustNewGrpcClientConn(serviceNameOrAddress, opts...)
}

func getAddress(ctx context.Context, service string) string {
	dev, err := cache.Get(ctx, "dev")
	if err != nil {
		return service
	}
	if dev.Bool() {
		devConfRaw, err := cache.Get(ctx, "devConf")
		if err != nil {
			return service
		}
		devConf := devConfRaw.Map()
		return devConf[service].(string)
	}
	return service
}

// UserClientConn 获取用户服务的客户端连接
func UserClientConn(ctx context.Context) *grpc.ClientConn {
	return ClientConn(ctx, "user")
}

// LibManagerClientConn 获取图书馆管理服务的连接
func LibManagerClientConn(ctx context.Context) *grpc.ClientConn {
	return ClientConn(ctx, "libManager")
}

// SeatClientConn 获取座位服务的连接
func SeatClientConn(ctx context.Context) *grpc.ClientConn {
	return ClientConn(ctx, "seat_handle")
}
