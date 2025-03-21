package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/Modou/pkg/logger"
)

type LoggerProvider struct {
	logger logger.Logger
}

func NewLoggerProvider(logger logger.Logger) *LoggerProvider {
	return &LoggerProvider{logger: logger}
}

func (p *LoggerProvider) Provide(c *gin.Context) {
	c.Set("logger", p.logger)
}
