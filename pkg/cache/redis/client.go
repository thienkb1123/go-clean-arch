package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/thienkb1123/go-clean-arch/config"
)

type RedisClient struct {
	rdbClient *redis.Client
}

// Returns new redis client
func NewRedisClient(cfg *config.RedisClient) (*RedisClient, error) {
	redisHost := cfg.RedisAddr

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: cfg.MinIdleConns,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &RedisClient{
		rdbClient: client,
	}, nil
}

func (r *RedisClient) Get(ctx context.Context, key string) ([]byte, error) {
	return r.rdbClient.Get(ctx, key).Bytes()
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.rdbClient.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Del(ctx context.Context, keys ...string) error {
	return r.rdbClient.Del(ctx, keys...).Err()
}

func (r *RedisClient) Close() error {
	return r.rdbClient.Close()
}

func (r *RedisClient) Ping(ctx context.Context) error {
	return r.rdbClient.Ping(ctx).Err()
}
