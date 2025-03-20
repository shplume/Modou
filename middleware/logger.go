package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/Modou/pkg/logger"
)

func LoggerMiddleware(logger logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	}
}
