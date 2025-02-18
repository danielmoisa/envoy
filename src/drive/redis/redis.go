package redis

import (
	"github.com/danielmoisa/envoy/src/utils/config"
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Config struct {
	Addr     string `env:"ENVOY_REDIS_ADDR" envDefault:"localhost"`
	Port     string `env:"ENVOY_REDIS_PORT" envDefault:"6379"`
	Password string `env:"ENVOY_REDIS_PASSWORD" envDefault:"envoypass"`
	Database int    `env:"ENVOY_REDIS_DATABASE" envDefault:"0"`
}

func NewRedisConnectionByGlobalConfig(config *config.Config, logger *zap.SugaredLogger) (*redis.Client, error) {
	redisConfig := &Config{
		Addr:     config.GetRedisAddr(),
		Port:     config.GetRedisPort(),
		Password: config.GetRedisPassword(),
		Database: config.GetRedisDatabase(),
	}
	return NewRedisConnection(redisConfig, logger)
}

func NewRedisConnection(config *Config, logger *zap.SugaredLogger) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Database,
	})

	logger.Infow("connected with redis", "redis", config)

	return rdb, nil
}
