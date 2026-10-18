package middleware

import (
	"runtime/debug"

	"go-gin-starter/internal/logger"
	"go-gin-starter/internal/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic recovered",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("stack", string(debug.Stack())),
				)

				response.ServerError(c, "internal server error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
