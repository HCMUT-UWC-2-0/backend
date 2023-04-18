// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: vehicles.sql

package db

import (
	"context"
)

const createVehicle = `-- name: CreateVehicle :one
INSERT INTO
  "Vehicles" (
    "make_by",
    "model",
    "capacity",
    "fuel_consumption"
  )
VALUES
  ($1, $2, $3, $4)
RETURNING id, make_by, model, capacity, fuel_consumption, created_at, updated_at
`

type CreateVehicleParams struct {
	MakeBy          string `json:"make_by"`
	Model           string `json:"model"`
	Capacity        string `json:"capacity"`
	FuelConsumption string `json:"fuel_consumption"`
}

func (q *Queries) CreateVehicle(ctx context.Context, arg CreateVehicleParams) (Vehicle, error) {
	row := q.db.QueryRowContext(ctx, createVehicle,
		arg.MakeBy,
		arg.Model,
		arg.Capacity,
		arg.FuelConsumption,
	)
	var i Vehicle
	err := row.Scan(
		&i.ID,
		&i.MakeBy,
		&i.Model,
		&i.Capacity,
		&i.FuelConsumption,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createVehicleStatus = `-- name: CreateVehicleStatus :one
INSERT INTO
  "VehicleStatus" (
    "vehicle_id"
  )
VALUES
  ($1)
RETURNING id, vehicle_id, status, current_fuel, created_at, updated_at
`

func (q *Queries) CreateVehicleStatus(ctx context.Context, vehicleID int32) (VehicleStatus, error) {
	row := q.db.QueryRowContext(ctx, createVehicleStatus, vehicleID)
	var i VehicleStatus
	err := row.Scan(
		&i.ID,
		&i.VehicleID,
		&i.Status,
		&i.CurrentFuel,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getVehicle = `-- name: GetVehicle :one
SELECT id, make_by, model, capacity, fuel_consumption, created_at, updated_at FROM "Vehicles" 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetVehicle(ctx context.Context, id int64) (Vehicle, error) {
	row := q.db.QueryRowContext(ctx, getVehicle, id)
	var i Vehicle
	err := row.Scan(
		&i.ID,
		&i.MakeBy,
		&i.Model,
		&i.Capacity,
		&i.FuelConsumption,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllVehicles = `-- name: ListAllVehicles :many
SELECT id, make_by, model, capacity, fuel_consumption, created_at, updated_at FROM "Vehicles" 
ORDER BY id
`

func (q *Queries) ListAllVehicles(ctx context.Context) ([]Vehicle, error) {
	rows, err := q.db.QueryContext(ctx, listAllVehicles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vehicle{}
	for rows.Next() {
		var i Vehicle
		if err := rows.Scan(
			&i.ID,
			&i.MakeBy,
			&i.Model,
			&i.Capacity,
			&i.FuelConsumption,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateVehicleStatus = `-- name: UpdateVehicleStatus :one
UPDATE "VehicleStatus"
SET "status" = $2
WHERE "vehicle_id" = $1
RETURNING id, vehicle_id, status, current_fuel, created_at, updated_at
`

type UpdateVehicleStatusParams struct {
	VehicleID int32             `json:"vehicle_id"`
	Status    VehicleStatusType `json:"status"`
}

func (q *Queries) UpdateVehicleStatus(ctx context.Context, arg UpdateVehicleStatusParams) (VehicleStatus, error) {
	row := q.db.QueryRowContext(ctx, updateVehicleStatus, arg.VehicleID, arg.Status)
	var i VehicleStatus
	err := row.Scan(
		&i.ID,
		&i.VehicleID,
		&i.Status,
		&i.CurrentFuel,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
