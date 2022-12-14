package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thienkb1123/go-clean-arch/internal/middleware"
	"github.com/thienkb1123/go-clean-arch/internal/news"
)

// Map news routes
func MapNewsRoutes(newsGroup fiber.Router, h news.Handlers, mw *middleware.MiddlewareManager) {
	newsGroup.Use(mw.AuthJWTMiddleware())
	newsGroup.Post("/create", h.Create)
	newsGroup.Put("/:news_id", h.Update)
	newsGroup.Delete("/:news_id", h.Delete)
	newsGroup.Get("/:news_id", h.GetByID)
	newsGroup.Get("", h.GetNews)
}
