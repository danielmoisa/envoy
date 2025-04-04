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
	routerGroup := engine.Group("/api/v1")

	// Protected middleware
	protected := routerGroup.Group("")
	protected.Use(middleware.AuthMiddleware())

	// Register routes
	r.RegisterHealthRoutes(routerGroup)
	r.registerAuthRoutes(routerGroup)
	r.registerUserRoutes(protected)
	r.registerCompanyRoutes(protected)
	r.registerJobRoutes(protected)
	r.registerApplicationRoutes(protected)
	r.registerCandidateRoutes(protected)
}

func (r *Router) RegisterHealthRoutes(group *gin.RouterGroup) {
	healthRouter := group.Group("/health")
	healthRouter.GET("", r.Controller.GetHealth)
}

func (r *Router) registerUserRoutes(group *gin.RouterGroup) {
	usersRouter := group.Group("/users")
	usersRouter.GET("/", r.Controller.GetAllUsers)
	usersRouter.GET("/:userId", r.Controller.GetUser)
	usersRouter.POST("/", r.Controller.CreateUser)
	usersRouter.PUT("/:userId", r.Controller.UpdateUser)
	usersRouter.DELETE("/:userId", r.Controller.DeleteUser)
}

func (r *Router) registerCompanyRoutes(group *gin.RouterGroup) {
	companiesRouter := group.Group("/companies")
	companiesRouter.GET("", r.Controller.GetAllCompanies)
	companiesRouter.GET("/:companyId", r.Controller.GetCompany)
	companiesRouter.POST("", r.Controller.CreateCompany)
	companiesRouter.PUT("/:companyId", r.Controller.UpdateCompany)
	companiesRouter.DELETE("/:companyId", r.Controller.DeleteCompany)
}

func (r *Router) registerAuthRoutes(group *gin.RouterGroup) {
	authRouter := group.Group("/auth")
	authRouter.POST("/login", r.Controller.Login)
	authRouter.POST("/logout", r.Controller.Logout)
}

func (r *Router) registerJobRoutes(group *gin.RouterGroup) {
	jobsRouter := group.Group("/jobs")
	jobsRouter.GET("", r.Controller.GetAllJobs)
	jobsRouter.GET("/:jobId", r.Controller.GetJob)
	jobsRouter.POST("", r.Controller.CreateJob)
	jobsRouter.PUT("/:jobId", r.Controller.UpdateJob)
	jobsRouter.DELETE("/:jobId", r.Controller.DeleteJob)
}

func (r *Router) registerApplicationRoutes(group *gin.RouterGroup) {
	applicationsRouter := group.Group("/applications")
	applicationsRouter.GET("", r.Controller.GetAllApplications)
	applicationsRouter.GET("/:applicationId", r.Controller.GetApplication)
	applicationsRouter.POST("", r.Controller.CreateApplication)
	applicationsRouter.PUT("/:applicationId", r.Controller.UpdateApplication)
	applicationsRouter.DELETE("/:applicationId", r.Controller.DeleteApplication)
}

func (r *Router) registerCandidateRoutes(group *gin.RouterGroup) {
	candidatesRouter := group.Group("/candidates")
	candidatesRouter.GET("", r.Controller.GetAllCandidates)
	candidatesRouter.GET("/:candidateId", r.Controller.GetCandidate)
	candidatesRouter.POST("", r.Controller.CreateCandidate)
	candidatesRouter.PUT("/:candidateId", r.Controller.UpdateCandidate)
	candidatesRouter.DELETE("/:candidateId", r.Controller.DeleteCandidate)
}
