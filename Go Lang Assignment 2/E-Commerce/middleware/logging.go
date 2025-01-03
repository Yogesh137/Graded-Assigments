package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddlewareGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Printf("Started %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
		fmt.Printf("Completed %s in %v\n", c.Request.URL.Path, time.Since(start))
	}
}
