package api

import (
	"errors"
	"fmt"
	"net/http"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/token"
	"github.com/HCMUT-UWC-2-0/backend/util"
	"github.com/gin-gonic/gin"
)

type taskRequest struct {
	StartTime   string `json:"startTime" binding:"required,date"`
	EndTime     string `json:"endTime" binding:"required,date"`
	JanitorID   int64  `json:"janitorId" binding:"required"`
	CollectorID int64  `json:"collectorId" binding:"required"`
	VehicleID   int64  `json:"vehicleId" binding:"required"`
	McpID       int64  `json:"mcpId" binding:"required"`
	RouteID     int64  `json:"routeId" binding:"required"`
}
type taskResponse struct {
	ID     int64             `json:"id"`
	Status db.TaskStatusType `json:"status"`
}

func newTaskResponse(task db.Task) taskResponse {
	return taskResponse{
		ID:     task.ID,
		Status: (task.Status),
	}
}

func (server *Server) createTask(ctx *gin.Context) {

	var req taskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	_, err := server.store.GetBackOfficer(ctx, authPayload.BackOfficerInfo.Email)
	if err != nil {
		err := errors.New("no right to create Task")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}

	start_time, err := util.MapStringToTime(req.StartTime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	end_time, err := util.MapStringToTime(req.EndTime)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var args db.CreateTaskParams = db.CreateTaskParams{
		StartTime:   start_time,
		EndTime:     end_time,
		JanitorID:   int32(req.JanitorID),
		CollectorID: int32(req.CollectorID),
		VehicleID:   int32(req.VehicleID),
		McpID:       int32(req.McpID),
		RouteID:     int32(req.RouteID),
	}

	task, err := server.store.InsertTaskTx(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newTaskResponse(task))

}

type currentTaskResponse struct {
	Janitor   string `json:"janitor"`
	Collector string `json:"collector"`
	Vehicle   string `json:"vehicle"`
	Route     string `json:"route"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Status    db.TaskStatusType `json:"status"`
}

func (server *Server) listAllCurrentTasks(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	_, err := server.store.GetBackOfficer(ctx, authPayload.BackOfficerInfo.Email)
	if err != nil {
		err := errors.New("no right to create Task")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}

	tasks, err := server.store.ListAllCurrentTasks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var currentTasks []currentTaskResponse
	for _, task := range tasks {
		janitor, err := server.store.GetWorker(ctx, int64(task.JanitorID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		collector, err := server.store.GetWorker(ctx, int64(task.CollectorID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		vehicle, err := server.store.GetVehicle(ctx, int64(task.VehicleID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		route, err := server.store.GetRoute(ctx, int64(task.RouteID))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		currentTasks = append(currentTasks, currentTaskResponse{
			Janitor:   janitor.Name,
			Collector: collector.Name,
			Vehicle:   fmt.Sprintf("%s@%d", vehicle.Model, vehicle.ID),
			Route:     fmt.Sprintf("%s - %s", route.StartLocation, route.EndLocation),
			StartTime: task.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:   task.EndTime.Format("2006-01-02 15:04:05"),
			Status:    task.Status,
		})
	}

	ctx.JSON(http.StatusOK, currentTasks)

}
