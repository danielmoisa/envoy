package router

import (
	"github.com/danielmoisa/envoy/src/controller"
	"github.com/danielmoisa/envoy/src/middleware"
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

	// Init base router group
	routerGroup := engine.Group("/api/v1")

	// Public routes
	healthRouter := routerGroup.Group("/health")
	authRouter := routerGroup.Group("/auth")

	// Health routes (public)
	healthRouter.GET("", r.Controller.GetHealth)

	// Auth routes (public)
	authRouter.POST("/login", r.Controller.Login)
	authRouter.POST("/logout", r.Controller.Logout)

	// Protected routes
	protected := routerGroup.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Users routes (protected)
		usersRouter := protected.Group("/users")
		{
			usersRouter.GET("/", r.Controller.GetAllUsers)
			usersRouter.GET("/:userId", r.Controller.GetUser)
			usersRouter.POST("/", r.Controller.CreateUser)
			usersRouter.PUT("/:userId", r.Controller.UpdateUser)
			usersRouter.DELETE("/:userId", r.Controller.DeleteUser)
		}

		// Companies routes (protected)
		companiesRouter := protected.Group("/companies")
		{
			companiesRouter.GET("", r.Controller.GetAllCompanies)
			companiesRouter.GET("/:companyId", r.Controller.GetCompany)
			companiesRouter.POST("", r.Controller.CreateCompany)
			companiesRouter.PUT("/:companyId", r.Controller.UpdateCompany)
			companiesRouter.DELETE("/:companyId", r.Controller.DeleteCompany)
		}

	}
}
