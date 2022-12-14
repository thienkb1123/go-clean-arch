package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/thienkb1123/go-clean-arch/config"
	"github.com/thienkb1123/go-clean-arch/pkg/cache/redis"
	"github.com/thienkb1123/go-clean-arch/pkg/logger"
	"gorm.io/gorm"
)

// Server struct
type Server struct {
	fiber  *fiber.App
	cfg    *config.Config
	db     *gorm.DB
	redis  redis.Client
	logger logger.Logger
}

// NewServer New Server constructor
func NewServer(cfg *config.Config, db *gorm.DB, opts ...Option) *Server {
	s := &Server{
		fiber: fiber.New(),
		cfg:   cfg,
		db:    db,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Run() error {
	if err := s.MapHandlers(); err != nil {
		return err
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		ln, _ := net.Listen("tcp", ":"+s.cfg.Server.Port)
		err := s.fiber.Listener(ln)
		if err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	s.logger.Info("Server Exited Properly")
	return s.fiber.Shutdown()
}
