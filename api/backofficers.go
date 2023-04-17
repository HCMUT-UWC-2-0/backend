package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/HCMUT-UWC-2-0/backend/token"
	"github.com/HCMUT-UWC-2-0/backend/util"
	"github.com/gin-gonic/gin"
)

type loginAccountRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginAccountResponse struct {
	AccessToken string    `json:"accessToken"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

func (server *Server) loginAccount(ctx *gin.Context) {
	var req loginAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	bo, err := server.store.GetBackOfficer(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, bo.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// using email + ssn for generating accessToken

	boAccessTokenInfo := token.BackOfficerInfo{
		Email: bo.Email,
		Ssn:   bo.Ssn,
	}

	accessToken, err := server.tokenMaker.CreateToken(
		boAccessTokenInfo,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var rsp = loginAccountResponse{
		AccessToken: accessToken,
		Name:        bo.Name,
		DateOfBirth: bo.DateOfBirth,
		Phone:       bo.Phone,
		Email:       bo.Email,
		ExpiredAt:   time.Now().Add(server.config.AccessTokenDuration),
	}
	ctx.JSON(http.StatusOK, rsp)
}
