package store

const (
	CREATE_DISH_QUERY = "INSERT INTO dish(restaurant_id, name, images, rating, created_by, created_at, updated_at) value ($1, $2, $3, $4, $5, $6, $7) returning restaurant_id, name, images, rating, created_by, created_at, updated_at"
)