package restaurants

import "time"

type Restaurant struct {
	ID          int32
	Name        *string
	Address     *string
	City        *string
	State       *string
	Star        int64
	PostalCode  *string
	Description *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
}
