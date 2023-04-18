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

-- name: GetVehicle :one
SELECT * FROM "Vehicles" 
WHERE id = $1 LIMIT 1;


-- name: CreateVehicleStatus :one
INSERT INTO
  "VehicleStatus" (
    "vehicle_id"
  )
VALUES
  ($1)
RETURNING *;



-- name: UpdateVehicleStatus :one
UPDATE "VehicleStatus"
SET "status" = $2
WHERE "vehicle_id" = $1
RETURNING *;