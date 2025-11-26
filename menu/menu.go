package menu

import (
	"context"
	"time"

	sql "cookdie/menu/sql/db/query_gen"

	"github.com/google/uuid"
)

type dishStore interface {
	CreateDish(dish *Dish) (*Dish, error)
}

type MenuService struct {
	store sql.Queries
}

type MenuServiceConfig struct {
	Store sql.Queries
}

func NewMenuService(cfg *MenuServiceConfig) *MenuService {
	return &MenuService{
		store: cfg.Store,
	}
}

func (s *MenuService) GetDishById(dishId uuid.UUID) (*sql.Dish, error) {
	dish, err := s.store.GetDishById(context.Background(), dishId)

	if err != nil {
		return nil, err
	}

	return &dish, nil
}
func (s *MenuService) CreateDish(input *Dish) (*sql.Dish, error) {
	err := validateDishInput(input)

	if err != nil {
		logger.Error("dish input not valid")
		return nil, err
	}

	t := time.Now()

	dishParams := sql.CreateDishParams{
		RestaurantID: *input.RestaurantId,
		Name:         input.Name,
		Images:       input.Images,
		Rating:       &input.Rating,
		CreatedBy:    *input.CreatedBy,
		CreatedAt:    t,
		UpdatedAt:    &t,
	}

	newDish, err := s.store.CreateDish(context.Background(), dishParams)

	if err != nil {
		return nil, err
	}

	return &newDish, nil
}

func validateDishInput(input *Dish) error {
	return nil
}
