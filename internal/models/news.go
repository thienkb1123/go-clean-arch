package models

import (
	"time"

	"github.com/google/uuid"
)

// News base model
type News struct {
	NewsID    uuid.UUID `json:"news_id" gorm:"column:news_id" validate:"omitempty"`
	AuthorID  uuid.UUID `json:"author_id,omitempty" gorm:"column:author_id" validate:"required"`
	Title     string    `json:"title" gorm:"column:title" validate:"required"`
	Content   string    `json:"content" gorm:"column:content" validate:"required"`
	ImageURL  *string   `json:"image_url,omitempty" gorm:"column:image_url"`
	Category  *string   `json:"category,omitempty" gorm:"column:category"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (*News) TableName() string {
	return "news"
}

// All News response
type NewsList struct {
	TotalCount int64   `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	News       []*News `json:"news"`
}

// News base
type NewsBase struct {
	NewsID    uuid.UUID `json:"news_id" gorm:"column:news_id" validate:"omitempty"`
	AuthorID  uuid.UUID `json:"author_id" gorm:"column:author_id" validate:"omitempty,uuid"`
	Title     string    `json:"title" gorm:"column:title" validate:"required"`
	Content   string    `json:"content" gorm:"column:content" validate:"required"`
	ImageURL  *string   `json:"image_url,omitempty" gorm:"column:image_url"`
	Category  *string   `json:"category,omitempty" gorm:"column:category"`
	Author    string    `json:"author" gorm:"column:author"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
