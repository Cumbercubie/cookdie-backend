package restaurants

import (
	"context"
	"time"

	restaurants "cookdie/restaurants/models"
	sql "cookdie/restaurants/sql/db/query_gen"

	"github.com/google/uuid"
)

type restaurantStore interface {
	CreateRestaurant(restaurant *restaurants.Restaurant) (*sql.Restaurant, error)
	GetRestaurantById(restaurantId uuid.UUID) (*sql.Restaurant, error)
}

type RestaurantService struct {
	store sql.Queries
}

type RestaurantServiceConfig struct {
	Store sql.Queries
}

func NewRestaurantService(cfg *RestaurantServiceConfig) *RestaurantService {
	return &RestaurantService{
		store: cfg.Store,
	}
}

func (s *RestaurantService) GetRestaurantById(restaurantId uuid.UUID) (*sql.Restaurant, error) {
	dish, err := s.store.GetRestaurantById(context.Background(), restaurantId)

	if err != nil {
		return nil, err
	}

	return &dish, nil
}
func (s *RestaurantService) CreateRestaurant(input *Restaurant) (*sql.Restaurant, error) {
	err := validateRestaurantInput(input)

	if err != nil {
		logger.Error("Restaurant input not valid")
		return nil, err
	}

	t := time.Now()

	restaurantParams := sql.CreateRestaurantParams{
		Name:        input.Name,
		Address:     input.Address,
		Star:        input.Star,
		City:        input.City,
		State:       input.State,
		PostalCode:  input.PostalCode,
		Description: input.Description,
	}

	newRestaurant, err := s.store.CreateRestaurant(context.Background(), restaurantParams)

	if err != nil {
		return nil, err
	}

	return &newRestaurant, nil
}

func validateRestaurantInput(input *Restaurant) error {
	return nil
}
