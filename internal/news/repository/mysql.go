package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/thienkb1123/go-clean-arch/internal/models"
	"github.com/thienkb1123/go-clean-arch/internal/news"
	"github.com/thienkb1123/go-clean-arch/pkg/utils"
	"gorm.io/gorm"
)

// News Repository
type newsRepo struct {
	db *gorm.DB
}

// News repository constructor
func NewNewsRepository(db *gorm.DB) news.Repository {
	return &newsRepo{db: db}
}

// Create news
func (r *newsRepo) Create(ctx context.Context, news *models.News) (*models.News, error) {
	err := r.db.Model(&models.News{}).Create(news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

// Update news item
func (r *newsRepo) Update(ctx context.Context, news *models.News) (*models.News, error) {
	db := r.db.Model(&models.News{})
	err := db.First(&news).Error
	if err != nil {
		return nil, err
	}

	err = db.Save(news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

// Get single news by id
func (r *newsRepo) GetNewsByID(ctx context.Context, newsID uuid.UUID) (*models.NewsBase, error) {
	result := &models.NewsBase{}
	err := r.db.Model(&models.User{}).Where("news_id", newsID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete news by id
func (r *newsRepo) Delete(ctx context.Context, newsID uuid.UUID) error {
	err := r.db.Model(&models.User{}).
		Where("news_id", newsID).
		Delete(&models.User{}).Error
	return err
}

// Get news
func (r *newsRepo) GetNews(ctx context.Context, pq *utils.PaginationQuery) (*models.NewsList, error) {
	totalCount := int64(0)
	db := r.db.WithContext(ctx).Model(&models.News{})
	db.Count(&totalCount)

	if totalCount == 0 {
		return &models.NewsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			News:       make([]*models.News, 0),
		}, nil
	}

	newsList := make([]*models.News, 0, pq.GetSize())
	err := db.Offset(pq.GetOffset()).Limit(pq.GetLimit()).Find(&newsList).Error
	if err != nil {
		return nil, err
	}
	return &models.NewsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		News:       newsList,
	}, nil
}
