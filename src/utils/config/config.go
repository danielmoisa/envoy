package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

const (
	DRIVE_TYPE_AWS   = "aws"
	DRIVE_TYPE_DO    = "do"
	DRIVE_TYPE_MINIO = "minio"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct{}

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: .env file not found or error loading it: %v", err)
		}
	})
	return instance
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Server Config Getters
func (c *Config) GetServerHost() string {
	return getEnv("ENVOY_SERVER_HOST", "0.0.0.0")
}

func (c *Config) GetServerPort() string {
	return getEnv("ENVOY_SERVER_PORT", "8001")
}

func (c *Config) GetServerMode() string {
	return getEnv("ENVOY_SERVER_MODE", "debug")
}

func (c *Config) GetSecretKey() string {
	return getEnv("ENVOY_SECRET_KEY", "8xEMrWkBARcDDYQ")
}

// ID Converter Getter
func (c *Config) GetRandomKey() string {
	return getEnv("ENVOY_RANDOM_KEY", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

// Postgres Config Getters
func (c *Config) GetPostgresAddr() string {
	return getEnv("ENVOY_PG_ADDR", "localhost")
}

func (c *Config) GetPostgresPort() string {
	return getEnv("ENVOY_PG_PORT", "5432")
}

func (c *Config) GetPostgresUser() string {
	return getEnv("ENVOY_PG_USER", "user")
}

func (c *Config) GetPostgresPassword() string {
	return getEnv("ENVOY_PG_PASSWORD", "password")
}

func (c *Config) GetPostgresDatabase() string {
	return getEnv("ENVOY_PG_DATABASE", "envoy")
}

func (c *Config) GetPostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.GetPostgresAddr(),
		c.GetPostgresPort(),
		c.GetPostgresUser(),
		c.GetPostgresPassword(),
		c.GetPostgresDatabase(),
	)
}

// Redis Config Getters
func (c *Config) GetRedisAddr() string {
	return getEnv("ENVOY_REDIS_ADDR", "localhost")
}

func (c *Config) GetRedisPort() string {
	return getEnv("ENVOY_REDIS_PORT", "6379")
}

func (c *Config) GetRedisPassword() string {
	return getEnv("ENVOY_REDIS_PASSWORD", "")
}

func (c *Config) GetRedisDatabase() int {
	val := getEnv("ENVOY_REDIS_DATABASE", "0")
	db, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return db
}

func (c *Config) GetRedisConnectionString() string {
	return fmt.Sprintf("%s:%s", c.GetRedisAddr(), c.GetRedisPort())
}

// Drive Config Getters
func (c *Config) GetDriveType() string {
	return getEnv("ENVOY_DRIVE_TYPE", "")
}

func (c *Config) GetDriveAccessKeyID() string {
	return getEnv("ENVOY_DRIVE_ACCESS_KEY_ID", "")
}

func (c *Config) GetDriveAccessKeySecret() string {
	return getEnv("ENVOY_DRIVE_ACCESS_KEY_SECRET", "")
}

func (c *Config) GetDriveRegion() string {
	return getEnv("ENVOY_DRIVE_REGION", "")
}

func (c *Config) GetDriveEndpoint() string {
	return getEnv("ENVOY_DRIVE_ENDPOINT", "")
}

func (c *Config) GetDriveSystemBucketName() string {
	return getEnv("ENVOY_DRIVE_SYSTEM_BUCKET_NAME", "envoy-cloud")
}

func (c *Config) GetDriveTeamBucketName() string {
	return getEnv("ENVOY_DRIVE_TEAM_BUCKET_NAME", "envoy-cloud-team")
}

func (c *Config) GetDriveUploadTimeout() time.Duration {
	val := getEnv("ENVOY_DRIVE_UPLOAD_TIMEOUT", "30s")
	duration, err := time.ParseDuration(val)
	if err != nil {
		return 30 * time.Second
	}
	return duration
}

// Drive Type Helpers
func (c *Config) IsAWSTypeDrive() bool {
	driveType := c.GetDriveType()
	return driveType == DRIVE_TYPE_AWS || driveType == DRIVE_TYPE_DO
}

func (c *Config) IsMINIODrive() bool {
	return c.GetDriveType() == DRIVE_TYPE_MINIO
}

// AWS/S3 Compatibility Methods
func (c *Config) GetAWSS3Endpoint() string         { return c.GetDriveEndpoint() }
func (c *Config) GetAWSS3AccessKeyID() string      { return c.GetDriveAccessKeyID() }
func (c *Config) GetAWSS3AccessKeySecret() string  { return c.GetDriveAccessKeySecret() }
func (c *Config) GetAWSS3Region() string           { return c.GetDriveRegion() }
func (c *Config) GetAWSS3SystemBucketName() string { return c.GetDriveSystemBucketName() }
func (c *Config) GetAWSS3TeamBucketName() string   { return c.GetDriveTeamBucketName() }
func (c *Config) GetAWSS3Timeout() time.Duration   { return c.GetDriveUploadTimeout() }

// MinIO Compatibility Methods
func (c *Config) GetMINIOAccessKeyID() string      { return c.GetDriveAccessKeyID() }
func (c *Config) GetMINIOAccessKeySecret() string  { return c.GetDriveAccessKeySecret() }
func (c *Config) GetMINIOEndpoint() string         { return c.GetDriveEndpoint() }
func (c *Config) GetMINIOSystemBucketName() string { return c.GetDriveSystemBucketName() }
func (c *Config) GetMINIOTeamBucketName() string   { return c.GetDriveTeamBucketName() }
func (c *Config) GetMINIOTimeout() time.Duration   { return c.GetDriveUploadTimeout() }
