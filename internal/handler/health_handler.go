package handler

import (
	"go_basic_framework_template/internal/response"

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
	response.Success(c, gin.H{
		"ready": true,
	})
}
