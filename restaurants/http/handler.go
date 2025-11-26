package http

import (
	"cookdie/restaurants"
	sql "cookdie/restaurants/sql/db/query_gen"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RestaurantService interface {
	CreateRestaurant(restaurant *restaurants.Restaurant) (*sql.Restaurant, error)
	GetRestaurantById(restaurantId uuid.UUID) (*sql.Restaurant, error)
}

//	type RecipeService interface {
//		CreateRecipe(input *)
//	}
func (rh *routeHandler) CreateRestaurant(c *gin.Context) {
	var input *restaurants.Restaurant

	if err := c.BindJSON(&input); err != nil {
		logger.Error("invalid input creating restaurants")
	}

	newRestaurant, err := rh.RestaurantService.CreateRestaurant(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(200, newRestaurant)
}

func (rh *routeHandler) GetDishById(c *gin.Context) {
	restaurantId, err := uuid.Parse(c.Param("restaurantId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	newDish, err := rh.RestaurantService.GetRestaurantById(restaurantId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	}
	c.JSON(200, newDish)
}
