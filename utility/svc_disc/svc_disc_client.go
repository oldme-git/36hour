package svc_disc

import (
	"context"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

// Client 获取一个 grpc 客户端
func Client(ctx context.Context, service string, opts ...grpc.DialOption) *grpc.ClientConn {
	// 获取 grpc 服务地址，如果是开发环境，则使用本地地址
	serviceNameOrAddress := getAddress(ctx, service)
	opts = append(opts,
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
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

// UserClient 获取用户服务的客户端
func UserClient(ctx context.Context) *grpc.ClientConn {
	return Client(ctx, "user")
}

// LibManagerClient 获取图书馆管理服务的客户端
func LibManagerClient(ctx context.Context) *grpc.ClientConn {
	return Client(ctx, "libManager")
}

// SeatClient 获取座位服务的客户端
func SeatClient(ctx context.Context) *grpc.ClientConn {
	return Client(ctx, "seat")
}
