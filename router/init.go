package router

import "github.com/shplume/Modou/core"

func Init(r *core.Server) {
	base := r.Group("/api/v1")

	initPingRouter(base)
}
