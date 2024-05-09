package svc_disc

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/grpc"
)

// Client 获取一个 grpc 客户端
func Client(service string, opts ...grpc.DialOption) *grpc.ClientConn {
	// 获取 grpc 服务地址，如果是开发环境，则使用本地地址
	serviceNameOrAddress := getAddress(service)
	return grpcx.Client.MustNewGrpcClientConn(serviceNameOrAddress, opts...)
}

func getAddress(service string) string {
	dev, err := cache.Get(gctx.New(), "dev")
	if err != nil {
		return service
	}
	if dev.Bool() {
		devConfRaw, err := cache.Get(gctx.New(), "devConf")
		if err != nil {
			return service
		}
		devConf := devConfRaw.Map()
		return devConf[service].(string)
	}
	return service
}

// UserClient 获取用户服务的客户端
func UserClient() *grpc.ClientConn {
	return Client("user")
}

// LibManagerClient 获取图书馆管理服务的客户端
func LibManagerClient() *grpc.ClientConn {
	return Client("lib-manager")
}

// SeatClient 获取座位服务的客户端
func SeatClient() *grpc.ClientConn {
	return Client("seat")
}
