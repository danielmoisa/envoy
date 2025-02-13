package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/caarlos0/env"
)

const DRIVE_TYPE_AWS = "aws"
const DRIVE_TYPE_DO = "do"
const DRIVE_TYPE_MINIO = "minio"

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		var err error
		if instance == nil {
			instance, err = getConfig() // not thread safe
			if err != nil {
				panic(err)
			}
		}
	})
	return instance
}

type Config struct {
	// Server config
	ServerHost string `env:"ENVOY_SERVER_HOST" envDefault:"0.0.0.0"`
	ServerPort string `env:"ENVOY_SERVER_PORT" envDefault:"8001"`
	ServerMode string `env:"ENVOY_SERVER_MODE" envDefault:"debug"`
	SecretKey  string `env:"ENVOY_SECRET_KEY" envDefault:"8xEMrWkBARcDDYQ"`

	// Key for idconvertor
	RandomKey string `env:"ENVOY_RANDOM_KEY"  envDefault:"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"`

	// Storage config
	PostgresAddr     string `env:"ENVOY_PG_ADDR" envDefault:"localhost"`
	PostgresPort     string `env:"ENVOY_PG_PORT" envDefault:"5432"`
	PostgresUser     string `env:"ENVOY_PG_USER" envDefault:"envoy_builder"`
	PostgresPassword string `env:"ENVOY_PG_PASSWORD" envDefault:"71De5JllWSetLYU"`
	PostgresDatabase string `env:"ENVOY_PG_DATABASE" envDefault:"envoy_builder"`

	// Cache config
	RedisAddr     string `env:"ENVOY_REDIS_ADDR" envDefault:"localhost"`
	RedisPort     string `env:"ENVOY_REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"ENVOY_REDIS_PASSWORD" envDefault:""`
	RedisDatabase int    `env:"ENVOY_REDIS_DATABASE" envDefault:"0"`

	// Drive config
	DriveType             string `env:"ENVOY_DRIVE_TYPE" envDefault:""`
	DriveAccessKeyID      string `env:"ENVOY_DRIVE_ACCESS_KEY_ID" envDefault:""`
	DriveAccessKeySecret  string `env:"ENVOY_DRIVE_ACCESS_KEY_SECRET" envDefault:""`
	DriveRegion           string `env:"ENVOY_DRIVE_REGION" envDefault:""`
	DriveEndpoint         string `env:"ENVOY_DRIVE_ENDPOINT" envDefault:""`
	DriveSystemBucketName string `env:"ENVOY_DRIVE_SYSTEM_BUCKET_NAME" envDefault:"envoy-cloud"`
	DriveTeamBucketName   string `env:"ENVOY_DRIVE_TEAM_BUCKET_NAME" envDefault:"envoy-cloud-team"`
	DriveUploadTimeoutRaw string `env:"ENVOY_DRIVE_UPLOAD_TIMEOUT" envDefault:"30s"`
	DriveUploadTimeout    time.Duration
}

func getConfig() (*Config, error) {
	// Fetch
	cfg := &Config{}
	err := env.Parse(cfg)

	// Process data
	var errInParseDuration error
	cfg.DriveUploadTimeout, errInParseDuration = time.ParseDuration(cfg.DriveUploadTimeoutRaw)
	if errInParseDuration != nil {
		return nil, errInParseDuration
	}

	// Ok
	fmt.Printf("----------------\n")
	fmt.Printf("run by following config: %+v\n", cfg)
	fmt.Printf("parse config error info: %+v\n", err)

	return cfg, err
}

func (c *Config) GetSecretKey() string {
	return c.SecretKey
}

func (c *Config) GetRandomKey() string {
	return c.RandomKey
}

func (c *Config) GetPostgresAddr() string {
	return c.PostgresAddr
}

func (c *Config) GetPostgresPort() string {
	return c.PostgresPort
}

func (c *Config) GetPostgresUser() string {
	return c.PostgresUser
}

func (c *Config) GetPostgresPassword() string {
	return c.PostgresPassword
}

func (c *Config) GetPostgresDatabase() string {
	return c.PostgresDatabase
}

func (c *Config) GetRedisAddr() string {
	return c.RedisAddr
}

func (c *Config) GetRedisPort() string {
	return c.RedisPort
}

func (c *Config) GetRedisPassword() string {
	return c.RedisPassword
}

func (c *Config) GetRedisDatabase() int {
	return c.RedisDatabase
}

func (c *Config) GetDriveType() string {
	return c.DriveType
}

func (c *Config) IsAWSTypeDrive() bool {
	if c.DriveType == DRIVE_TYPE_AWS || c.DriveType == DRIVE_TYPE_DO {
		return true
	}
	return false
}

func (c *Config) IsMINIODrive() bool {
	return c.DriveType == DRIVE_TYPE_MINIO
}

func (c *Config) GetAWSS3Endpoint() string {
	return c.DriveEndpoint
}

func (c *Config) GetAWSS3AccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetAWSS3AccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetAWSS3Region() string {
	return c.DriveRegion
}

func (c *Config) GetAWSS3SystemBucketName() string {
	return c.DriveSystemBucketName
}

func (c *Config) GetAWSS3TeamBucketName() string {
	return c.DriveTeamBucketName
}

func (c *Config) GetAWSS3Timeout() time.Duration {
	return c.DriveUploadTimeout
}

func (c *Config) GetMINIOAccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetMINIOAccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetMINIOEndpoint() string {
	return c.DriveEndpoint
}

func (c *Config) GetMINIOSystemBucketName() string {
	return c.DriveSystemBucketName
}

func (c *Config) GetMINIOTeamBucketName() string {
	return c.DriveTeamBucketName
}

func (c *Config) GetMINIOTimeout() time.Duration {
	return c.DriveUploadTimeout
}
