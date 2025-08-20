package menu

type restaurantStore interface {
}
type RestaurantService struct {
	store restaurantStore
}

type RestaurantServiceConfig struct {
	Store restaurantStore
}

func NewRestaurantService(cfg *RestaurantServiceConfig) *RestaurantService {
	return &RestaurantService{
		store: cfg.Store,
	}
}
