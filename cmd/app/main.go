package main

import (
	"log"

	"github.com/thienkb1123/go-clean-arch/config"
	"github.com/thienkb1123/go-clean-arch/internal/server"
	"github.com/thienkb1123/go-clean-arch/pkg/cache/redis"
	"github.com/thienkb1123/go-clean-arch/pkg/database/mysql"
	"github.com/thienkb1123/go-clean-arch/pkg/logger"
)

func main() {
	log.Println("Starting api server")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	// token, _ := utils.GenerateJWTToken(&models.User{
	// 	UserID: uuid.New(),
	// }, cfg)

	// fmt.Println("token: ", token)

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	// Repository
	mysqlDB, err := mysql.New(&cfg.MySQL)
	if err != nil {
		appLogger.Fatalf("MySQL init: %s", err)
	}

	rdb, err := redis.NewClient(&cfg.Redis)
	if err != nil {
		appLogger.Fatalf("RedisCluster init: %s", err)
	}

	s := server.NewServer(
		cfg,
		mysqlDB,
		server.Redis(rdb),
		server.Logger(appLogger),
	)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
