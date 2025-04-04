package api

import (
	"CodeManager/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

type Handler struct{
	usecases *usecases.Usecase
}

func NewHandler(usecases *usecases.Usecase) *Handler {
	return &Handler{
		usecases: usecases,
	}
}


// RegisterRoutes registers all routes for the API
func (h *Handler) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default()) 

	health := router.Group("/health")
	{
		health.GET("/ping", h.Ping)
	}

	domain := router.Group("/api/v1")
	{
		domain.POST("/run-and-analyze", h.RunAndAnalyzeHandler)
		domain.GET("/runtimes", h.GetRuntimesHandler)
	}

	swagger := router.Group("")
	{
		swagger.StaticFile("/doc.json", "./docs/swagger.json")
	}
	
	url := ginSwagger.URL("http://localhost:8001/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
