package svc_disc

import (
	"context"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var cache = gcache.New()

// Init 初始化服务发现
// 会初始化 etcd 服务发现
// 如果是 dev 环境，会读取 etcd-dev.yaml 配置到缓存，方便本地开发访问 grpc 服务
func Init(ctx context.Context) {
	confRaw, err := g.Cfg("etcd").Get(ctx, "etcd")
	if err != nil {
		panic(err)
	}

	var (
		conf    = confRaw.Map()
		address = conf["address"].(string)
		env     = conf["env"].(string)
	)

	// 注册 etcd 服务发现
	Register(address)

	// 如果是测试环境，保存 dev 配置到缓存
	if env == "dev" {
		err = cache.Set(ctx, "dev", true, 0)
		if err != nil {
			panic(err)
		}

		err = cache.Set(ctx, "devConf", conf["dev"], 0)
		if err != nil {
			panic(err)
		}
	}
}

// Register 注册服务发现，使用 etcd
func Register(address string, option ...etcd.Option) {
	grpcx.Resolver.Register(etcd.New(address, option...))
}
