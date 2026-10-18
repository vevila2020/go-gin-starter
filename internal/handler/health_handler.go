package handler

import (
	"net/http"

	"go-gin-starter/internal/database"
	"go-gin-starter/internal/redis"
	"go-gin-starter/internal/response"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	response.Success(c, gin.H{
		"status": "ok",
	})
}

func (h *HealthHandler) Ready(c *gin.Context) {
	checks := gin.H{
		"database": "disabled",
		"redis":    "disabled",
	}

	if database.IsEnabled() {
		if err := database.DB.Exec("SELECT 1").Error; err != nil {
			c.JSON(http.StatusServiceUnavailable, response.Response{
				Code:    response.CodeServerError,
				Message: "database not ready",
			})
			return
		}
		checks["database"] = "ok"
	}

	if redis.IsEnabled() {
		if err := redis.Client.Ping(c.Request.Context()).Err(); err != nil {
			c.JSON(http.StatusServiceUnavailable, response.Response{
				Code:    response.CodeServerError,
				Message: "redis not ready",
			})
			return
		}
		checks["redis"] = "ok"
	}

	response.Success(c, gin.H{
		"ready":  true,
		"checks": checks,
	})
}