package controller

import (
	"github.com/shplume/Modou/core"
)

func Ping(c *core.Context) {
	c.Logger().Info("recv ping")

	c.JSON(200, successResponse("pong"))
}
