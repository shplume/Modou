package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/Modou/pkg/config"
)

type ConfigPriovider struct {
	config config.ConfigReader
}

func NewConfigProvider(config config.ConfigReader) *ConfigPriovider {
	return &ConfigPriovider{config: config}
}

func (p *ConfigPriovider) Provide(c *gin.Context) {
	c.Set("config", p.config)
}
