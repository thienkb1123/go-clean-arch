package http

import (
	"github.com/gin-gonic/gin"
	"github.com/thienkb1123/go-clean-arch/internal/middleware"
	"github.com/thienkb1123/go-clean-arch/internal/news"
)

// Map news routes
func MapNewsRoutes(newsGroup *gin.RouterGroup, h news.Handlers, mw *middleware.MiddlewareManager) {
	newsGroup.Use(mw.AuthJWTMiddleware())
	newsGroup.POST("/create", h.Create)
	newsGroup.POST("/:newsId", h.Update)
	newsGroup.POST("/:newsId", h.Delete)
	newsGroup.POST("/:newsId", h.GetByID)
	newsGroup.POST("", h.GetNews)
}
