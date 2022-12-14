package repository

import (
	"context"
	"time"

	"github.com/thienkb1123/go-clean-arch/internal/models"
	"github.com/thienkb1123/go-clean-arch/internal/news"
	"github.com/thienkb1123/go-clean-arch/pkg/cache/redis"
	"github.com/thienkb1123/go-clean-arch/pkg/converter"
	"github.com/thienkb1123/go-clean-arch/pkg/errors"
)

// News redis repository
type newsRedisRepo struct {
	rdb redis.Client
}

// News redis repository constructor
func NewNewsRedisRepo(rdb redis.Client) news.RedisRepository {
	return &newsRedisRepo{rdb: rdb}
}

// Get new by id
func (n *newsRedisRepo) GetNewsByIDCtx(ctx context.Context, key string) (*models.NewsBase, error) {
	newsBytes, err := n.rdb.Get(ctx, key)
	if err != nil {
		return nil, errors.WithMessage(err, "newsRedisRepo.GetNewsByIDCtx.redisClusterClient.Get")
	}
	newsBase := &models.NewsBase{}
	if err = converter.BytesToAny(newsBytes, newsBase); err != nil {
		return nil, errors.WithMessage(err, "newsRedisRepo.GetNewsByIDCtx.json.Unmarshal")
	}

	return newsBase, nil
}

// Cache news item
func (n *newsRedisRepo) SetNewsCtx(ctx context.Context, key string, seconds int, news *models.NewsBase) error {
	newsBytes, err := converter.AnyToBytes(news)
	if err != nil {
		return errors.WithMessage(err, "newsRedisRepo.SetNewsCtx.json.Marshal")
	}
	if err = n.rdb.Set(ctx, key, newsBytes, time.Second*time.Duration(seconds)); err != nil {
		return errors.WithMessage(err, "newsRedisRepo.SetNewsCtx.redisClusterClient.Set")
	}
	return nil
}

// Delete new item from cache
func (n *newsRedisRepo) DeleteNewsCtx(ctx context.Context, key string) error {
	if err := n.rdb.Del(ctx, key); err != nil {
		return errors.WithMessage(err, "newsRedisRepo.DeleteNewsCtx.redisClusterClient.Del")
	}
	return nil
}
