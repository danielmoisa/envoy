package main

import (
	"os"

	"github.com/danielmoisa/envoy/src/router"
	"github.com/danielmoisa/envoy/src/utils/config"
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

func (server *Server) Start() {
	server.logger.Infow("Starting illa-builder-backend...")

	// init
	gin.SetMode(server.config.ServerMode)

	// init cors
	// server.engine.Use(gin.CustomRecovery(recovery.CorsHandleRecovery))
	// server.engine.Use(cors.Cors())
	server.router.RegisterRouters(server.engine)

	// run
	err := server.engine.Run(server.config.ServerHost + ":" + server.config.ServerPort)
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
