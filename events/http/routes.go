package http

import (
	"cookdie/restaurants"

	"github.com/gin-gonic/gin"
)

type routeHandler struct {
	EventService EventService
}

type RouteHandlerConfig struct {
	EventService EventService
}

func NewRouteHandler(config *RouteHandlerConfig) *routeHandler {

	return &routeHandler{
		EventService: config.EventService,
	}

}

func (rh *routeHandler) RegisterRoutes(r *gin.Engine) {

	rh.registerEventRoutes(r)

}

var logger restaurants.Logger
