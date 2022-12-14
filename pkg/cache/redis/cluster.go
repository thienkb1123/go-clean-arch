package redis

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/thienkb1123/go-clean-arch/config"
)

type RedisCluster struct {
	rdbCluster *redis.ClusterClient
}

// Returns new redis cluster client
func NewRedisClusterClient(cfg *config.RedisCluster) (*RedisCluster, error) {
	addrs := strings.Split(cfg.Addrs, ",")
	cluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:         addrs,
		DialTimeout:   time.Duration(cfg.DialTimeout) * time.Second,
		ReadTimeout:   time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:  time.Duration(cfg.WriteTimeout) * time.Second,
		PoolSize:      cfg.PoolSize,
		PoolTimeout:   time.Duration(cfg.PoolTimeout) * time.Second,
		RouteRandomly: true,
	})

	err := cluster.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &RedisCluster{
		rdbCluster: cluster,
	}, nil
}

func (r *RedisCluster) Get(ctx context.Context, key string) ([]byte, error) {
	return r.rdbCluster.Get(ctx, key).Bytes()
}

func (r *RedisCluster) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.rdbCluster.Set(ctx, key, value, expiration).Err()
}

func (r *RedisCluster) Del(ctx context.Context, keys ...string) error {
	return r.rdbCluster.Del(ctx, keys...).Err()
}

func (r *RedisCluster) Close() error {
	return r.rdbCluster.Close()
}

func (r *RedisCluster) Ping(ctx context.Context) error {
	return r.rdbCluster.Ping(ctx).Err()
}
