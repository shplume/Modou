package core

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/Modou/pkg/logger"
)

type Context struct {
	*gin.Context
}

func (c *Context) GetLogger() logger.Logger {
	log, exists := c.Get("logger")
	if !exists {
		return logger.DefaultLoggerInstance
	}

	return log.(logger.Logger)
}
