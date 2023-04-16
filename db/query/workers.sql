-- name: CreateWorker :one
INSERT INTO
  "Workers" (
    "ssn",
    "name",
    "phone",
    "age",
    "worker_type",
    "gender",
    "date_of_birth",
    "place_of_birth"
  )
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: ListAllWorkers :many
SELECT * FROM "Workers" 
WHERE "worker_type" = $1
ORDER BY id;