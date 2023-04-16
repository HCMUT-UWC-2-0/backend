-- name: CreateVehicle :one
INSERT INTO
  "Vehicles" (
    "make_by",
    "model",
    "capacity",
    "fuel_consumption"
  )
VALUES
  ($1, $2, $3, $4)
RETURNING *;

-- name: ListAllVehicles :many
SELECT * FROM "Vehicles" 
ORDER BY id;