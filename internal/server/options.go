package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thienkb1123/go-clean-arch/pkg/cache/redis"
	"github.com/thienkb1123/go-clean-arch/pkg/logger"
)

// Option -.
type Option func(*Server)

func FiberEngine(gin *gin.Engine) Option {
	return func(s *Server) {
		s.gin = gin
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
