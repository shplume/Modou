package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/Modou/core"
)

func Ping(c *core.Context) {
	c.GetLogger().Info("recv ping")

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
