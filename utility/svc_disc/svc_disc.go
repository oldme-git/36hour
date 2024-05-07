package svc_disc

import (
	"context"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc"
)

// Register 注册服务发现，使用 etcd
func Register(address string, option ...etcd.Option) {
	grpcx.Resolver.Register(etcd.New(address, option...))
}

// RegisterWithConf 使用 etcd.yaml 配置文件注册服务发现
func RegisterWithConf(ctx context.Context) {
	address, err := g.Cfg("etcd").Get(ctx, "address")
	if err != nil {
		panic(err)
	}
	Register(address.String())
}

// Client 获取一个 grpc 客户端
func Client(serviceNameOrAddress string, opts ...grpc.DialOption) *grpc.ClientConn {
	return grpcx.Client.MustNewGrpcClientConn(serviceNameOrAddress, opts...)
}
