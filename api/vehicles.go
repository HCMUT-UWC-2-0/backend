package api

import (
	"errors"
	"net/http"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/token"
	"github.com/gin-gonic/gin"
)


type vehicleResponse struct {
	MakeBy string `json:"makeBy"`
	Model  string `json:"model"`
	Capacity string `json:"capacity"`
	FuelConsumption string `json:"fuelConsumption"`

}

func newVehicleResponse (vehicle db.Vehicle) vehicleResponse {
	return vehicleResponse{
		MakeBy: vehicle.MakeBy,
		Model: vehicle.Model,
		Capacity: vehicle.Capacity,
		FuelConsumption: vehicle.FuelConsumption,
	}
}

func (server *Server) listAllVehicles(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	_, err := server.store.GetBackOfficer(ctx, authPayload.BackOfficerInfo.Email)
	if err != nil {
		err := errors.New("no right to get all workers")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}

	vehicles, err := server.store.ListAllVehicles(ctx)
	rsp := make([]vehicleResponse, len(vehicles))
	for i, vehicle := range vehicles {
		rsp[i] = newVehicleResponse(vehicle)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rsp)

}
