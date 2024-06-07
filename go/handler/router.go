package handler

import "github.com/gin-gonic/gin"

type Router interface {
	Routes() []Route
}

type Route struct {
	method   string
	path     string
	handlers []gin.HandlerFunc
}

func NewRoute(method, path string, handlers ...gin.HandlerFunc) Route {
	return Route{
		method:   method,
		path:     path,
		handlers: handlers,
	}
}

type rootRouter struct {
	*gin.Engine
	routers []Router
}

func NewRootRouter(engine *gin.Engine, routers ...Router) *rootRouter {
	r := &rootRouter{routers: routers}
	r.Init(engine)
	return r
}

func (g *rootRouter) Init(engine *gin.Engine) {
	g.Engine = engine
	g.setRoutes()
}

func (g *rootRouter) setRoutes() {
	for _, router := range g.routers {
		for _, route := range router.Routes() {
			g.Engine.Handle(route.method, route.path, route.handlers...)
		}
	}
}

func (g *rootRouter) Handle(r Router) {
	g.routers = append(g.routers, r)
}
