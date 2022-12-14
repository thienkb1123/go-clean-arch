//go:generate mockgen -source mysql_repository.go -destination mock/mysql_repository_mock.go -package mock

package news

import (
	"context"

	"github.com/google/uuid"
	"github.com/thienkb1123/go-clean-arch/internal/models"
	"github.com/thienkb1123/go-clean-arch/pkg/utils"
)

// News Repository
type Repository interface {
	Create(ctx context.Context, news *models.News) (*models.News, error)
	Update(ctx context.Context, news *models.News) (*models.News, error)
	GetNewsByID(ctx context.Context, newsID uuid.UUID) (*models.NewsBase, error)
	Delete(ctx context.Context, newsID uuid.UUID) error
	GetNews(ctx context.Context, pq *utils.PaginationQuery) (*models.NewsList, error)
}
