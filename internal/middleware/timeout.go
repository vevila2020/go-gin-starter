package middleware

import (
	"context"
	"net/http"
	"time"

	"go-gin-starter/internal/response"

	"github.com/gin-gonic/gin"
)

func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()

		if ctx.Err() == context.DeadlineExceeded && !c.Writer.Written() {
			response.Error(c, http.StatusGatewayTimeout, "request timeout")
			c.Abort()
		}
	}
}