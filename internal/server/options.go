package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thienkb1123/go-clean-arch/pkg/cache/redis"
	"github.com/thienkb1123/go-clean-arch/pkg/logger"
)

// Option -.
type Option func(*Server)

func FiberEngine(fiber *fiber.App) Option {
	return func(s *Server) {
		s.fiber = fiber
	}
}

func Redis(rdb redis.Client) Option {
	return func(s *Server) {
		s.redis = rdb
	}
}

func Logger(logger logger.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}
