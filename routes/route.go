package routes

import "github.com/gin-gonic/gin"

type RouteHandlerFunc func(ctx *gin.Context) (interface{}, error)

type Route interface {
	GetMethod() string
	GetPath() string
	GetMiddleware() []gin.HandlerFunc
	GetHandler() RouteHandlerFunc
}

type routeData struct {
	method     string
	path       string
	middleware []gin.HandlerFunc
	handler    RouteHandlerFunc
}

func (r *routeData) GetMethod() string {
	return r.method
}

func (r *routeData) GetPath() string {
	return r.path
}

func (r *routeData) GetMiddleware() []gin.HandlerFunc {
	return r.middleware
}

func (r *routeData) GetHandler() RouteHandlerFunc {
	return r.handler
}
