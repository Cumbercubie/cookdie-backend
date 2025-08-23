-- name: CreateDish :one
INSERT INTO dish (
  id, restaurant_id, name, images, rating, created_by, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetDishById :one
SELECT *
FROM dish
WHERE id = $1;