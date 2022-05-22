package clientplugin

import (
	"context"
	"time"

	"github.com/flare-rpc/flarego/client"
	"github.com/flare-rpc/flarego/server"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

var _ client.PreCallPlugin = (*RedisRateLimitingPlugin)(nil)

// RedisRateLimitingPlugin can limit requests per unit time
type RedisRateLimitingPlugin struct {
	addrs   []string
	limiter redis_rate.Limiter
	limit   redis_rate.Limit
}

// NewRedisRateLimitingPlugin creates a new RateLimitingPlugin
func NewRedisRateLimitingPlugin(addrs []string, rate int, burst int, period time.Duration) *RedisRateLimitingPlugin {
	limit := redis_rate.Limit{
		Rate:   rate,
		Burst:  burst,
		Period: period,
	}
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
	})

	limiter := redis_rate.NewLimiter(rdb)

	return &RedisRateLimitingPlugin{
		addrs:   addrs,
		limiter: *limiter,
		limit:   limit,
	}
}

// PreReadRequest can limit request processing.
func (plugin *RedisRateLimitingPlugin) PreCall(ctx context.Context, servicePath, serviceMethod string, args interface{}) error {
	res, err := plugin.limiter.Allow(ctx, servicePath+"/"+serviceMethod, plugin.limit)
	if err != nil {
		return err
	}

	if res.Allowed > 0 {
		return nil
	}
	return server.ErrReqReachLimit
}
