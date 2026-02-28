package middleware

import (
	"go_basic_framework_template/internal/response"

	"github.com/gin-gonic/gin"
)

func RateLimit(limit int, window int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.Unauthorized(c, "missing authorization token")
			c.Abort()
			return
		}

		c.Set("user_id", "example_user")
		c.Next()
	}
}

func GetUserID(c *gin.Context) string {
	if userID, exists := c.Get("user_id"); exists {
		if id, ok := userID.(string); ok {
			return id
		}
	}
	return ""
}
