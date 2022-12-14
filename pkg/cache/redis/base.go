package redis

import (
	"context"
	"time"

	"github.com/thienkb1123/go-clean-arch/config"
)

const (
	redisClusterMode    = "cluster"
	redisStandaloneMode = "standalone"
)

type (
	Client interface {
		Get(ctx context.Context, key string) ([]byte, error)
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
		Del(ctx context.Context, keys ...string) error
		Close() error
		Ping(ctx context.Context) error
	}
)

func NewClient(cfg *config.RedisConfig) (Client, error) {
	if cfg.Mode == redisStandaloneMode {
		return NewRedisClient(&cfg.Standalone)
	}

	return NewRedisClusterClient(&cfg.Cluster)
}
