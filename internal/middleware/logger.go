package middleware

import (
	"github.com/gofiber/fiber/v2"
	pkgLogger "github.com/thienkb1123/go-clean-arch/pkg/logger"
	"github.com/thienkb1123/go-clean-arch/pkg/utils"
)

// LoggerMiddleware set the logger with some fields inside the logger.
func (mw *MiddlewareManager) LoggerMiddleware(l pkgLogger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		ctx = l.WithFields(ctx, pkgLogger.Fields{
			"METHOD":     c.Method(),
			"PATH":       c.Path(),
			"REQUEST_ID": utils.GetRequestID(c),
		})
		c.SetUserContext(ctx)
		return c.Next()
	}
}
