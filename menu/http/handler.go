package http

import (
	"cookdie/menu"
	sql "cookdie/menu/sql/db/query_gen"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MenuService interface {
	CreateDish(dish *menu.Dish) (*sql.Dish, error)
	GetDishById(dishId uuid.UUID) (*sql.Dish, error)
}

//	type RecipeService interface {
//		CreateRecipe(input *)
//	}

func (rh *routeHandler) registerMenuRoutes(r *gin.Engine) {
	apiV1 := r.Group("/v1")

	apiV1.GET("/api/restaurants/menu/dishes", rh.GetDishById)
	apiV1.POST("/api/restaurants/menu/dishes", rh.CreateDish)
}

func (rh *routeHandler) CreateDish(c *gin.Context) {
	var input *menu.Dish

	if err := c.BindJSON(&input); err != nil {
		logger.Error("unable to bind json")
	}

	newDish, err := rh.MenuService.CreateDish(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	}
	c.JSON(200, newDish)
}

func (rh *routeHandler) GetDishById(c *gin.Context) {
	dishId, err := uuid.Parse(c.Param("disId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	newDish, err := rh.MenuService.GetDishById(dishId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	}
	c.JSON(200, newDish)
}
