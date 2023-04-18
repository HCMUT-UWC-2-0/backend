package api

import (
	"errors"
	"net/http"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/token"
	"github.com/gin-gonic/gin"
)


type routeResponse struct {
	StartLocation  string `json:"startLocation"`
	EndLocation  string `json:"endLocation"`
	Distance  string `json:"distance"`
	EstimatedTime  string `json:"estimatedTime"`

}

func newRouteResponse (route db.Route) routeResponse {
	return routeResponse{
		StartLocation: route.StartLocation,
		EndLocation: route.EndLocation,
		Distance: route.Distance,
		EstimatedTime: route.EstimatedTime,
	}
}

func (server *Server) listAllRoutes(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	_, err := server.store.GetBackOfficer(ctx, authPayload.BackOfficerInfo.Email)
	if err != nil {
		err := errors.New("no right to get all routes")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}

	routes, err := server.store.ListAllRoutes(ctx)
	rsp := make([]routeResponse, len(routes))
	for i, route := range routes {
		rsp[i] = newRouteResponse(route)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rsp)

}
