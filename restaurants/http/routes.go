package http

import (
	"cookdie/restaurants"

	"github.com/gin-gonic/gin"
)

type routeHandler struct {
	RestaurantService RestaurantService
}

type RouteHandlerConfig struct {
	RestaurantService RestaurantService
}

func NewRouteHandler(config *RouteHandlerConfig) *routeHandler {

	return &routeHandler{
		RestaurantService: config.RestaurantService,
	}

}

func (rh *routeHandler) RegisterRoutes(r *gin.Engine) {

	rh.registerRestaurantRoutes(r)

}

var logger restaurants.Logger
