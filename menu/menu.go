package menu

import (
	"context"
	"time"

	sql "cookdie/menu/sql/db/query_gen"
)

type dishStore interface {
	CreateDish(dish *Dish) (*Dish, error)
}

type MenuService struct {
	queries sql.Queries
}

func (s *MenuService) CreateDish(input *Dish) (*sql.Dish, error) {
	err := validateDishInput(input)

	if err != nil {
		logger.Error("dish input not valid")
		return nil, err
	}

	dishParams := sql.CreateMenuParams{
		RestaurantID: *input.RestaurantId,
		Name:         input.Name,
		Images:       input.Images,
		Rating:       &input.Rating,
		CreatedBy:    *input.CreatedBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	newDish, err := s.queries.CreateMenu(context.Background(), dishParams)

	if err != nil {
		return nil, err
	}

	return &newDish, nil
}

func validateDishInput(input *Dish) error {
	return nil
}
