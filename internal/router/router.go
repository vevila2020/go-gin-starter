package router

import (
	"time"

	"go-gin-starter/internal/config"
	"go-gin-starter/internal/handler"
	"go-gin-starter/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	timeout := 30 * time.Second
	if cfg := config.GetConfig(); cfg != nil && cfg.Server.RequestTimeout > 0 {
		timeout = time.Duration(cfg.Server.RequestTimeout) * time.Second
	}
	r.Use(middleware.Timeout(timeout))

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