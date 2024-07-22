package cache

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ExpireAtEndOfDay 设置缓存过期时间为当天的23:59:59
func ExpireAtEndOfDay(ctx context.Context, key string) (err error) {
	var (
		redis    = g.Redis()
		expireAt = gtime.Now().EndOfDay().Time
	)
	_, err = redis.ExpireAt(ctx, key, expireAt)
	if err != nil {
		return err
	}
	return
}
