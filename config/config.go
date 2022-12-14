package config

import (
	"github.com/caarlos0/env/v6"
)

// App config struct
type Config struct {
	Server  ServerConfig
	MySQL   MySQLConfig
	Redis   RedisConfig
	Logger  Logger
	Metrics Metrics
}

// Server config struct
type ServerConfig struct {
	AppVersion        string `env:"APP_VERSION"`
	Port              string `env:"PORT"`
	Mode              string `env:"MODE"`
	JwtSecretKey      string `env:"JWT_SECRET_KEY"`
	ReadTimeout       int    `env:"READ_TIMEOUT"`
	WriteTimeout      int    `env:"WRITE_TIMEOUT"`
	CtxDefaultTimeout int    `env:"CTX_DEFAULT_TIMEOUT"`
	Debug             bool   `env:"DEBUG"`
}

// Metrics config
type Metrics struct {
	URL         string `env:"METRICS_URL"`
	ServiceName string `env:"METRICS_SERVICE_NAME"`
}

// Logger config
type Logger struct {
	Development       bool   `env:"LOGGER_DEVELOPMENT"`
	DisableCaller     bool   `env:"LOGGER_DISABLE_CALLER"`
	DisableStacktrace bool   `env:"LOGGER_DISABLE_STACKTRACE"`
	Encoding          string `env:"LOGGER_ENCODING"`
	Level             string `env:"LOGGER_LEVEL"`
}

// Postgresql config
type MySQLConfig struct {
	URI             string `env:"MYSQL_URI"`
	MaxIdleConns    int    `env:"MYSQL_MAX_IDLE_CONNS"`
	MaxOpenConns    int    `env:"MYSQL_MAX_OPEN_CONNS"`
	ConnMaxLifeTime int    `env:"MYSQL_CON_MAX_LIFE_TIME"`
	Timeout         int    `env:"MYSQL_TIMEOUT"`
	Debug           bool   `env:"MYSQL_DEBUG"`
}

// Redis config

type RedisConfig struct {
	Mode       string `env:"REDIS_MODE"`
	Cluster    RedisCluster
	Standalone RedisClient
}

type RedisCluster struct {
	Addrs        string `env:"REDIS_CLUSTER_ADDRS"`
	DialTimeout  int    `env:"REDIS_CLUSTER_DIAL_TIMEOUT"`
	ReadTimeout  int    `env:"REDIS_CLUSTER_READ_TIMEOUT"`
	WriteTimeout int    `env:"REDIS_CLUSTER_WRITE_TIMEOUT"`
	PoolSize     int    `env:"REDIS_CLUSTER_POOL_SIZE"`
	PoolTimeout  int    `env:"REDIS_CLUSTER_POOL_TIMEOUT"`
}

type RedisClient struct {
	RedisAddr    string `env:"REDIS_CLIENT_ADDR"`
	MinIdleConns int    `env:"REDIS_CLIENT_MIN_IDLE_CONNS"`
	PoolSize     int    `env:"REDIS_CLIENT_POOL_SIZE"`
	PoolTimeout  int    `env:"REDIS_CLIENT_POOL_TIMEOUT"`
}

// Load config file from given path
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
