// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: task.sql

package db

import (
	"context"
	"time"
)

const createTask = `-- name: CreateTask :one
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
RETURNING id, start_time, end_time, janitor_id, collector_id, vehicle_id, mcp_id, route_id, status, created_at, updated_at
`

type CreateTaskParams struct {
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	JanitorID   int32     `json:"janitor_id"`
	CollectorID int32     `json:"collector_id"`
	VehicleID   int32     `json:"vehicle_id"`
	McpID       int32     `json:"mcp_id"`
	RouteID     int32     `json:"route_id"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.StartTime,
		arg.EndTime,
		arg.JanitorID,
		arg.CollectorID,
		arg.VehicleID,
		arg.McpID,
		arg.RouteID,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.StartTime,
		&i.EndTime,
		&i.JanitorID,
		&i.CollectorID,
		&i.VehicleID,
		&i.McpID,
		&i.RouteID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}