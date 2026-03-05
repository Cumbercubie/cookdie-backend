package http

import (
	events "cookdie/events/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (rh *routeHandler) registerEventRoutes(r *gin.Engine) {
	apiV1 := r.Group("/v1")

	apiV1.POST("/api/event/checkin", rh.CheckinRestaurantHandler)
	apiV1.POST("/api/event/review", rh.ReviewRestaurant)
}

type EventService interface {
	CheckinRestaurant(restaurantId string, userId string, date time.Time) (*events.Checkin, error)
	PostReviewRestaurant(review *events.Review) (*events.Review, error)
}

func (rh *routeHandler) ReviewRestaurant(c *gin.Context) {
	var input *events.Review

	if err := c.BindJSON(&input); err != nil {
		logger.Error("unable to bind json")
	}
	review, err := rh.EventService.PostReviewRestaurant(input)

	if err != nil {
		logger.Error("unable to checkin")
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, review)
	}
}

func (rh *routeHandler) CheckinRestaurantHandler(c *gin.Context) {
	var input *events.Checkin

	if err := c.BindJSON(&input); err != nil {
		logger.Error("unable to bind json")
	}
	dateTime, err := time.Parse("2006-01-02 15:04:05", input.Date)

	if err != nil {
		logger.Error("unable to parse checkin time")
		c.JSON(http.StatusBadRequest, err.Error())
	}
	newCheckin, err := rh.EventService.CheckinRestaurant(input.BusinessId, input.UserId, dateTime)

	if err != nil {
		logger.Error("unable to checkin")
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, newCheckin)
	}

}
