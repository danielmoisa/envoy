package router

import (
	"github.com/danielmoisa/envoy/src/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Controller *controller.Controller
}

func NewRouter(controller *controller.Controller) *Router {
	return &Router{
		Controller: controller,
	}
}

func (r *Router) RegisterRoutes(engine *gin.Engine) {
	// Config
	engine.UseRawPath = true

	// Init routers
	routerGroup := engine.Group("/api/v1")

	// builderRouter := routerGroup.Group("/teams/:teamID/builder")
	// appRouter := routerGroup.Group("/teams/:teamID/apps")
	// appsRouter := routerGroup.Group("/apps")
	// publicAppRouter := routerGroup.Group("/teams/byIdentifier/:teamIdentifier/publicApps")
	// resourceRouter := routerGroup.Group("/teams/:teamID/resources")
	// actionRouter := routerGroup.Group("/teams/:teamID/apps/:appID/actions")
	// publicActionRouter := routerGroup.Group("/teams/byIdentifier/:teamIdentifier/apps/:appID/publicActions")
	// internalActionRouter := routerGroup.Group("/teams/:teamID/apps/:appID/internalActions")
	// roomRouter := routerGroup.Group("/teams/:teamID/room")
	// oauth2Router := routerGroup.Group("/oauth2")
	// flowActionRouter := routerGroup.Group("/teams/:teamID/workflow/:workflowID/flowActions")
	healthRouter := routerGroup.Group("/health")
	usersRouter := routerGroup.Group("/users")

	// Users routes
	usersRouter.GET("/:teamId", r.Controller.GetAllUsers)
	usersRouter.GET("/teamId/:userId", r.Controller.GetUser)
	usersRouter.POST("/", r.Controller.CreateUser)

	// Health Routes
	healthRouter.GET("", r.Controller.GetHealth)
}
