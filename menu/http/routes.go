package http

import (
	"cookdie/menu"

	"github.com/gin-gonic/gin"
)

type routeHandler struct {
	MenuService MenuService
}

type RouteHandlerConfig struct {
	MenuService MenuService
}

func NewRouteHandler(config *RouteHandlerConfig) *routeHandler {

	return &routeHandler{
		MenuService: config.MenuService,
	}

}

func (rh *routeHandler) RegisterRoutes(r *gin.Engine) {

	rh.registerMenuRoutes(r)

}

var logger menu.Logger
