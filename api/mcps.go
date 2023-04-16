package api

import (
	"errors"
	"net/http"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/token"
	"github.com/gin-gonic/gin"
)


type mcpResponse struct {
	Location  string `json:"location"`
	Capacity string `json:"capacity"`

}

func newMcpResponse (mcp db.MCP) mcpResponse {
	return mcpResponse{
		Location: mcp.Location,
		Capacity: mcp.Capacity,
	}
}

func (server *Server) listAllMCPs(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	_, err := server.store.GetBackOfficer(ctx, authPayload.BackOfficerInfo.Email)
	if err != nil {
		err := errors.New("no right to get all workers")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}

	mcps, err := server.store.ListAllMCPs(ctx)
	rsp := make([]mcpResponse, len(mcps))
	for i, mcp := range mcps {
		rsp[i] = newMcpResponse(mcp)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rsp)

}
