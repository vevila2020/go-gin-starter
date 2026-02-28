package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
