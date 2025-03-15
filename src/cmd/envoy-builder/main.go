package main

import (
	"os"

	"github.com/danielmoisa/envoy/src/cache"
	"github.com/danielmoisa/envoy/src/controller"
	"github.com/danielmoisa/envoy/src/drive"
	"github.com/danielmoisa/envoy/src/drive/postgres"
	"github.com/danielmoisa/envoy/src/drive/redis"
	"github.com/danielmoisa/envoy/src/drive/s3"
	"github.com/danielmoisa/envoy/src/repository"
	"github.com/danielmoisa/envoy/src/router"
	"github.com/danielmoisa/envoy/src/utils/config"
	"github.com/danielmoisa/envoy/src/utils/cors"
	"github.com/danielmoisa/envoy/src/utils/logger"
	"github.com/danielmoisa/envoy/src/utils/recovery"
	"github.com/danielmoisa/envoy/src/utils/swagger"
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

func initStorage(globalConfig *config.Config, logger *zap.SugaredLogger) *repository.Repository {
	postgresDriver, err := postgres.NewPostgresConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		logger.Errorw("Error in startup, storage init failed.")
	}
	return repository.NewRepository(postgresDriver, logger)
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
	appStorage := initStorage(globalConfig, sugaredLogger)
	appCache := initCache(globalConfig, sugaredLogger)
	appDrive := initDrive(globalConfig, sugaredLogger)

	// Initialize Swagger UI and documentation
	swagger.InitSwagger(engine)

	// init attribute group
	// attrg, errInNewAttributeGroup := accesscontrol.NewRawAttributeGroup()
	// if errInNewAttributeGroup != nil {
	// 	return nil, errInNewAttributeGroup
	// }

	// Init controller
	c := controller.NewControllerForBackend(appStorage, appCache, appDrive)
	appRouter := router.NewRouter(c)
	server := NewServer(globalConfig, engine, appRouter, sugaredLogger)
	return server, nil

}

func (server *Server) Start() {
	server.logger.Infow("Starting envoy...")

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

// @title Envoy API
// @version 1.0
// @description This is a sample API
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com
func main() {
	server, err := initServer()

	if err != nil {
		server.logger.Errorw("Error initializing server: %v", err)
	}

	server.Start()
}
