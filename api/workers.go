package api

import (
	"errors"
	"net/http"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/token"
	"github.com/gin-gonic/gin"
)

// LIST
type listAllWorkersRequest struct {
	Type db.WorkerType `form:"type" binding:"required,min=1"`
}

func (server *Server) listAllWorkers(ctx *gin.Context) {
	var req listAllWorkersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	_, err := server.store.GetBackOfficer(ctx, authPayload.BackOfficerInfo.Email)
	if err != nil {
		err := errors.New("no right to get all workers")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}

	arg := req.Type
	
	workers, err := server.store.ListAllWorkers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, workers)

}
