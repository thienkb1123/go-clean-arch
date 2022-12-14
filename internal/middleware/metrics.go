package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/thienkb1123/go-clean-arch/pkg/metric"
)

// Prometheus metrics middleware
func (mw *MiddlewareManager) MetricsMiddleware(metrics metric.Metrics) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		method := c.Route().Method

		path := c.Route().Path
		if metrics.SkipPath(path) {
			return c.Next()
		}

		err := c.Next()
		status := fiber.StatusInternalServerError
		if err != nil {
			if e, ok := err.(*fiber.Error); ok {
				status = e.Code
			}
		} else {
			status = c.Response().StatusCode()
		}

		metrics.ObserveResponseTime(status, method, path, time.Since(start).Seconds())
		metrics.IncHits(status, method, path)
		return err
	}

}
