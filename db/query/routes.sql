-- name: CreateRoute :one
INSERT INTO
  "Routes" (
    "start_location",
    "end_location",
    "distance",
    "estimated_time"
  )
VALUES
  ($1, $2, $3, $4)
RETURNING *;

-- name: GetRoute :one
SELECT * FROM "Routes" 
WHERE id = $1 LIMIT 1;

-- name: ListAllRoutes :many
SELECT * FROM "Routes" 
ORDER BY id;
