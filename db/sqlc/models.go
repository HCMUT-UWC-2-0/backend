// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type GenderType string

const (
	GenderTypeMALE   GenderType = "MALE"
	GenderTypeFEMALE GenderType = "FEMALE"
)

func (e *GenderType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = GenderType(s)
	case string:
		*e = GenderType(s)
	default:
		return fmt.Errorf("unsupported scan type for GenderType: %T", src)
	}
	return nil
}

type NullGenderType struct {
	GenderType GenderType
	Valid      bool // Valid is true if GenderType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGenderType) Scan(value interface{}) error {
	if value == nil {
		ns.GenderType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.GenderType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGenderType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.GenderType), nil
}

type TasksStatus string

const (
	TasksStatusOPENED TasksStatus = "OPENED"
	TasksStatusDONE   TasksStatus = "DONE"
)

func (e *TasksStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TasksStatus(s)
	case string:
		*e = TasksStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for TasksStatus: %T", src)
	}
	return nil
}

type NullTasksStatus struct {
	TasksStatus TasksStatus
	Valid       bool // Valid is true if TasksStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTasksStatus) Scan(value interface{}) error {
	if value == nil {
		ns.TasksStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TasksStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTasksStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TasksStatus), nil
}

type VehicleStatusType string

const (
	VehicleStatusTypeAVAILABLE VehicleStatusType = "AVAILABLE"
	VehicleStatusTypeUSINGE    VehicleStatusType = "USINGE"
)

func (e *VehicleStatusType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VehicleStatusType(s)
	case string:
		*e = VehicleStatusType(s)
	default:
		return fmt.Errorf("unsupported scan type for VehicleStatusType: %T", src)
	}
	return nil
}

type NullVehicleStatusType struct {
	VehicleStatusType VehicleStatusType
	Valid             bool // Valid is true if VehicleStatusType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVehicleStatusType) Scan(value interface{}) error {
	if value == nil {
		ns.VehicleStatusType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VehicleStatusType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVehicleStatusType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VehicleStatusType), nil
}

type WorkerStatusType string

const (
	WorkerStatusTypeAVAILABLE WorkerStatusType = "AVAILABLE"
	WorkerStatusTypeWORKING   WorkerStatusType = "WORKING"
)

func (e *WorkerStatusType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkerStatusType(s)
	case string:
		*e = WorkerStatusType(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkerStatusType: %T", src)
	}
	return nil
}

type NullWorkerStatusType struct {
	WorkerStatusType WorkerStatusType
	Valid            bool // Valid is true if WorkerStatusType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkerStatusType) Scan(value interface{}) error {
	if value == nil {
		ns.WorkerStatusType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WorkerStatusType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkerStatusType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WorkerStatusType), nil
}

type WorkerType string

const (
	WorkerTypeJANITOR   WorkerType = "JANITOR"
	WorkerTypeCOLLECTOR WorkerType = "COLLECTOR"
)

func (e *WorkerType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkerType(s)
	case string:
		*e = WorkerType(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkerType: %T", src)
	}
	return nil
}

type NullWorkerType struct {
	WorkerType WorkerType
	Valid      bool // Valid is true if WorkerType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkerType) Scan(value interface{}) error {
	if value == nil {
		ns.WorkerType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WorkerType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkerType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WorkerType), nil
}

type BackOfficer struct {
	ID             int64       `json:"id"`
	Email          string      `json:"email"`
	Ssn            string      `json:"ssn"`
	HashedPassword string      `json:"hashed_password"`
	Name           string      `json:"name"`
	Phone          string      `json:"phone"`
	Age            int32       `json:"age"`
	Gender         interface{} `json:"gender"`
	DateOfBirth    time.Time   `json:"date_of_birth"`
	PlaceOfBirth   string      `json:"place_of_birth"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

type MCP struct {
	ID        int64     `json:"id"`
	Location  string    `json:"location"`
	Capacity  string    `json:"capacity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MCPStatus struct {
	ID               int64     `json:"id"`
	McpID            string    `json:"mcp_id"`
	Location         string    `json:"location"`
	Capacity         string    `json:"capacity"`
	CurrentLevelFill string    `json:"current_level_fill"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Route struct {
	ID            int64     `json:"id"`
	StartLocation string    `json:"start_location"`
	EndLocation   string    `json:"end_location"`
	Distance      string    `json:"distance"`
	EstimatedTime string    `json:"estimated_time"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Task struct {
	ID        int64       `json:"id"`
	StartTime time.Time   `json:"start_time"`
	EndTime   time.Time   `json:"end_time"`
	WorkerID  int32       `json:"worker_id"`
	VehicleID int32       `json:"vehicle_id"`
	McpID     int32       `json:"mcp_id"`
	RouteID   int32       `json:"route_id"`
	Status    interface{} `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Vehicle struct {
	ID              int64     `json:"id"`
	MakeBy          string    `json:"make_by"`
	Model           string    `json:"model"`
	Capacity        string    `json:"capacity"`
	FuelConsumption string    `json:"fuel_consumption"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type VehicleStatus struct {
	ID          int64       `json:"id"`
	VehicleID   string      `json:"vehicle_id"`
	Status      interface{} `json:"status"`
	CurrentFuel string      `json:"current_fuel"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Worker struct {
	ID           int64       `json:"id"`
	Ssn          string      `json:"ssn"`
	Name         string      `json:"name"`
	Phone        string      `json:"phone"`
	Age          int32       `json:"age"`
	WorkerType   interface{} `json:"worker_type"`
	Gender       interface{} `json:"gender"`
	DateOfBirth  time.Time   `json:"date_of_birth"`
	PlaceOfBirth string      `json:"place_of_birth"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type WorkerStatus struct {
	ID        int64       `json:"id"`
	WorkerID  string      `json:"worker_id"`
	Status    interface{} `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}