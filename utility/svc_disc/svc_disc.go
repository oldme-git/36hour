package svc_disc

import (
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

// Register 注册服务发现，使用 etcd
func Register(address string, option ...etcd.Option) {
	grpcx.Resolver.Register(etcd.New(address, option...))
}

// Client 获取一个 grpc 客户端
func Client(serviceNameOrAddress string, opts ...grpc.DialOption) *grpc.ClientConn {
	return grpcx.Client.MustNewGrpcClientConn(serviceNameOrAddress, opts...)
}
