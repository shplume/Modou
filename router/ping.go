package router

import (
	"github.com/shplume/Modou/controller"
	"github.com/shplume/Modou/core"
)

func initPingRouter(router *core.RouterGroup) {
	router.GET("/ping", controller.Ping)
}
