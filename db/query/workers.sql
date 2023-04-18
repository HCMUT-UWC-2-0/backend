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


-- name: CreateWorkerStatus :one
INSERT INTO
  "WorkerStatus" (
    "worker_id"
  )
VALUES
  ($1)
RETURNING *;


-- name: UpdateWorkerStatus :one
UPDATE "WorkerStatus"
SET "status" = $2
WHERE "worker_id" = $1
RETURNING *;


-- name: GetWorker :one
SELECT * FROM "Workers" 
WHERE id = $1 LIMIT 1;
