package core

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/Modou/pkg/config"
	"github.com/shplume/Modou/pkg/logger"
)

type Context struct {
	*gin.Context
}

type ContextManager struct {
	providers []Provider
}

type Provider interface {
	Provide(c *gin.Context)
}

func NewContextManager() *ContextManager {
	return &ContextManager{
		providers: make([]Provider, 0),
	}
}

func (m *ContextManager) Use(provider Provider) {
	m.providers = append(m.providers, provider)
}

func (m *ContextManager) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, provider := range m.providers {
			provider.Provide(c)
		}
		c.Next()
	}
}

func (c *Context) Logger() logger.Logger {
	log, exists := c.Get("logger")
	if !exists {
		return logger.GetDefaultLogger()
	}

	return log.(logger.Logger)
}

func (c *Context) Config() config.ConfigReader {
	conf, exists := c.Get("config")
	if !exists {
		return config.GetDefaultConfigReader()
	}

	return conf.(config.ConfigReader)
}
