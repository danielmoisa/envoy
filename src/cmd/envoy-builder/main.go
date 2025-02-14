package main

import (
	"os"

	"github.com/danielmoisa/envoy/src/cache"
	"github.com/danielmoisa/envoy/src/controller"
	"github.com/danielmoisa/envoy/src/drive"
	"github.com/danielmoisa/envoy/src/drive/postgres"
	"github.com/danielmoisa/envoy/src/drive/redis"
	"github.com/danielmoisa/envoy/src/drive/s3"
	"github.com/danielmoisa/envoy/src/router"
	"github.com/danielmoisa/envoy/src/storage"
	"github.com/danielmoisa/envoy/src/utils/config"
	"github.com/danielmoisa/envoy/src/utils/cors"
	"github.com/danielmoisa/envoy/src/utils/logger"
	"github.com/danielmoisa/envoy/src/utils/recovery"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router *router.Router
	engine *gin.Engine
	logger *zap.SugaredLogger
	config *config.Config
}

func NewServer(config *config.Config, engine *gin.Engine, router *router.Router, logger *zap.SugaredLogger) *Server {
	return &Server{
		router: router,
		engine: engine,
		config: config,
		logger: logger,
	}
}

func initStorage(globalConfig *config.Config, logger *zap.SugaredLogger) *storage.Storage {
	postgresDriver, err := postgres.NewPostgresConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		logger.Errorw("Error in startup, storage init failed.")
	}
	return storage.NewStorage(postgresDriver, logger)
}

func initCache(globalConfig *config.Config, logger *zap.SugaredLogger) *cache.Cache {
	redisDriver, err := redis.NewRedisConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		logger.Errorw("Error in startup, cache init failed.")
	}
	return cache.NewCache(redisDriver, logger)
}

func initDrive(globalConfig *config.Config, logger *zap.SugaredLogger) *drive.Drive {
	if globalConfig.IsAWSTypeDrive() {
		teamAWSConfig := s3.NewTeamAwsConfigByGlobalConfig(globalConfig)
		teamDriveS3Instance := s3.NewS3Drive(teamAWSConfig)
		return drive.NewDrive(teamDriveS3Instance, logger)
	}

	// Failed
	logger.Errorw("Error in startup, drive init failed.")
	return nil
}

func initServer() (*Server, error) {
	globalConfig := config.GetInstance()
	engine := gin.New()
	sugaredLogger := logger.NewSugardLogger()

	// init validator
	// validator := tokenvalidator.NewRequestTokenValidator()

	// Init driver
	storage := initStorage(globalConfig, sugaredLogger)
	cache := initCache(globalConfig, sugaredLogger)
	drive := initDrive(globalConfig, sugaredLogger)

	// init attribute group
	// attrg, errInNewAttributeGroup := accesscontrol.NewRawAttributeGroup()
	// if errInNewAttributeGroup != nil {
	// 	return nil, errInNewAttributeGroup
	// }

	// Init controller
	c := controller.NewControllerForBackend(storage, cache, drive)
	router := router.NewRouter(c)
	server := NewServer(globalConfig, engine, router, sugaredLogger)
	return server, nil

}

func (server *Server) Start() {
	server.logger.Infow("Starting envoy-builder...")

	// Init
	gin.SetMode(server.config.GetServerMode())

	// Init cors
	server.engine.Use(gin.CustomRecovery(recovery.CorsHandleRecovery))
	server.engine.Use(cors.Cors())
	server.router.RegisterRoutes(server.engine)

	// Run
	err := server.engine.Run(server.config.GetServerHost() + ":" + server.config.GetServerPort())
	if err != nil {
		server.logger.Errorw("Error in startup", "err", err)
		os.Exit(2)
	}
}

func main() {
	server, err := initServer()

	if err != nil {

	}

	server.Start()
}
