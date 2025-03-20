package core

import (
	"github.com/gin-gonic/gin"
)

type Handler func(*Context)

func WrapHandler(handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{Context: c}
		handler(ctx)
	}
}

type RouterGroup struct {
	group *gin.RouterGroup
}

func (g *RouterGroup) GET(path string, handler Handler) {
	g.group.GET(path, WrapHandler(handler))
}

func (g *RouterGroup) POST(path string, handler Handler) {
	g.group.POST(path, WrapHandler(handler))
}

func (g *RouterGroup) PUT(path string, handler Handler) {
	g.group.PUT(path, WrapHandler(handler))
}

func (g *RouterGroup) DELETE(path string, handler Handler) {
	g.group.DELETE(path, WrapHandler(handler))
}

func (g *RouterGroup) PATCH(path string, handler Handler) {
	g.group.PATCH(path, WrapHandler(handler))
}

func (g *RouterGroup) OPTIONS(path string, handler Handler) {
	g.group.OPTIONS(path, WrapHandler(handler))
}

func (g *RouterGroup) HEAD(path string, handler Handler) {
	g.group.HEAD(path, WrapHandler(handler))
}

func (g *RouterGroup) Any(path string, handler Handler) {
	g.group.Any(path, WrapHandler(handler))
}

func (g *RouterGroup) USE(middleware ...Handler) {
	for _, m := range middleware {
		g.group.Use(WrapHandler(m))
	}
}

func (g *RouterGroup) Group(path string, handlers ...Handler) *RouterGroup {
	group := g.group.Group(path)
	for _, handler := range handlers {
		group.Use(WrapHandler(handler))
	}
	return &RouterGroup{group: group}
}

func (g *RouterGroup) Handle(method, path string, handler Handler) {
	g.group.Handle(method, path, WrapHandler(handler))
}

func (g *RouterGroup) Match(method []string, path string, handler Handler) {
	g.group.Match(method, path, WrapHandler(handler))
}

func (s *Server) Group(path string, handlers ...Handler) *RouterGroup {
	group := s.Engine.Group(path)
	for _, handler := range handlers {
		group.Use(WrapHandler(handler))
	}
	return &RouterGroup{group: group}
}
