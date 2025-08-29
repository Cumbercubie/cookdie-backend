-- name: CreateRestaurant :one
INSERT INTO restaurants (
    name,
    address,
    city,
    state,
    star,
    postal_code,
    description
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;


-- name: GetRestaurantById :one
SELECT *
FROM restaurants
WHERE id = $1;
