package server

import (
	"github.com/gofiber/fiber/v2/middleware/requestid"
	apiMiddlewares "github.com/thienkb1123/go-clean-arch/internal/middleware"
	newsHttp "github.com/thienkb1123/go-clean-arch/internal/news/delivery/http"
	newsRepository "github.com/thienkb1123/go-clean-arch/internal/news/repository"
	newsUseCase "github.com/thienkb1123/go-clean-arch/internal/news/usecase"
	"github.com/thienkb1123/go-clean-arch/pkg/metric"
)

// Map Server Handlers
func (s *Server) MapHandlers() error {
	metrics, err := metric.CreateMetrics(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	if err != nil {
		s.logger.Errorf("CreateMetrics Error: %s", err)
	}
	s.logger.Info(
		"Metrics available URL: %s, ServiceName: %s",
		s.cfg.Metrics.URL,
		s.cfg.Metrics.ServiceName,
	)

	metrics.SetSkipPath([]string{"readiness"})

	// Init repositories
	nRepo := newsRepository.NewNewsRepository(s.db)
	newsRedisRepo := newsRepository.NewNewsRedisRepo(s.redis)

	// Init useCases
	newsUC := newsUseCase.NewNewsUseCase(s.cfg, nRepo, newsRedisRepo, s.logger)

	// Init handlers
	newsHandlers := newsHttp.NewNewsHandlers(s.cfg, newsUC, s.logger)

	mw := apiMiddlewares.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)

	// s.gin.Use(metric.RecordMetrics("go-clean-arch", []string{"metrics", "readiness"}))

	// v1 := s.gin.Group("/api/v1")

	// newsGroup := v1.Group("/news")

	s.fiber.Use(requestid.New())
	s.fiber.Use(mw.MetricsMiddleware(metrics))

	v1 := s.fiber.Group("/api/v1")
	newsGroup := v1.Group("/news")

	newsHttp.MapNewsRoutes(newsGroup, newsHandlers, mw)

	return nil
}
