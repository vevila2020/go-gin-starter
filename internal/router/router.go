package router

import (
	"go_basic_framework_template/internal/handler"
	"go_basic_framework_template/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	healthHandler := handler.NewHealthHandler()
	r.GET("/health", healthHandler.Health)
	r.GET("/ready", healthHandler.Ready)

	api := r.Group("/api/v1")
	{
		userHandler := handler.NewUserHandler()
		users := api.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.List)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}
}
