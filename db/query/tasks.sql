-- name: CreateTask :one
INSERT INTO
  "Tasks" (
   "start_time",
    "end_time",
    "janitor_id",
    "collector_id",
    "vehicle_id" ,
    "mcp_id" ,
    "route_id"
  )
VALUES
  ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;


-- name: ListAllCurrentTasks :many
SELECT * FROM "Tasks"
WHERE "end_time" > NOW()
ORDER BY "created_at" DESC;


