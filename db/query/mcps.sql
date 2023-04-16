-- name: CreateMCP :one
INSERT INTO
  "MCPs" (
    "location",
    "capacity"
  )
VALUES
  ($1, $2)
RETURNING *;

-- name: ListAllMCPs :many
SELECT * FROM "MCPs" 
ORDER BY id;