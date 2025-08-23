CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- for gen_random_uuid()

CREATE TABLE dish (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  restaurant_id UUID NOT NULL,
  name TEXT NOT NULL,
  images TEXT[] DEFAULT '{}',
  rating DOUBLE,
  created_by UUID NOT NULL,
  updated_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT NOW() NOT NULL
);


-- name: GetMenu :one
SELECT * FROM dish WHERE id = $1;


-- name: UpdateMenuRating :exec
UPDATE dish
SET rating = $1, updated_at = NOW()
WHERE id = $2;
