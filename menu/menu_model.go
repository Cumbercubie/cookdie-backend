package menu

import (
	"time"

	"github.com/google/uuid"
)

type Dish struct {
	RestaurantId *uuid.UUID
	Name         string
	Images       []string
	Rating       float64
	CreatedBy    *uuid.UUID
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
