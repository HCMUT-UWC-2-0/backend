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


