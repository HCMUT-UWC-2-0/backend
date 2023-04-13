package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	db "github.com/DV-Lab/zuni-backend/db/sqlc"
	"github.com/DV-Lab/zuni-backend/token"
	"github.com/DV-Lab/zuni-backend/util"
	"github.com/gin-gonic/gin"
)

type AccountResponse struct {
	Role        db.RoleType `json:"role"`
	PublicKey   string      `json:"public_key"`
	AccountInfo interface{} `json:"account_info"`
	CreatedAt   time.Time   `json:"create_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type loginAccountRequest struct {
	AccountId string `json:"account_id" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

type loginAccountResponse struct {
	AccessToken string          `json:"access_token"`
	Account     AccountResponse `json:"account_info"`
}

func newAccountResponse(account db.Account, accountInfo interface{}) AccountResponse {
	return AccountResponse{
		Role:        account.Role,
		PublicKey:   account.PublicKey,
		AccountInfo: accountInfo,
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   account.UpdatedAt,
	}
}

func (server *Server) loginAccount(ctx *gin.Context) {
	var req loginAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.AccountId)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, account.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// using student's name for generating accessToken

	accountTokenInfo := token.AccountInfo{
		AccountId: account.AccountID,
		Role:      account.Role,
	}

	accessToken, err := server.tokenMaker.CreateToken(
		accountTokenInfo,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var rsp loginAccountResponse
	if account.Role == db.RoleTypeADMIN {
		// ADMIN case
		admin, err := server.store.GetAdmin(ctx, account.AccountID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		rsp = loginAccountResponse{
			AccessToken: accessToken,
			Account:     newAccountResponse(account, newAdminResponse(admin)),
		}
	} else {
		// STUDENT case
		student, err := server.store.GetStudent(ctx, account.AccountID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		rsp = loginAccountResponse{
			AccessToken: accessToken,
			Account:     newAccountResponse(account, newStudentResponse(student)),
		}
	}

	ctx.JSON(http.StatusOK, rsp)
}

// Get Account

type getAccountRequest struct {
	AccountID string `uri:"account_id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	//========SHOULD BIND URI===========
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := server.store.GetAccount(ctx, req.AccountID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.AccountID != authPayload.AccountInfo.AccountId || account.Role != authPayload.AccountInfo.Role {
		err := errors.New("account doesn't belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var rsp AccountResponse
	if account.Role == db.RoleTypeADMIN {
		// ADMIN case
		admin, err := server.store.GetAdmin(ctx, account.AccountID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		rsp = newAccountResponse(account, newAdminResponse(admin))
	} else {
		// STUDENT case
		student, err := server.store.GetStudent(ctx, account.AccountID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		rsp = newAccountResponse(account, newStudentResponse(student))
	}

	ctx.JSON(http.StatusOK, rsp)
}

type updatePasswordStudentRequest struct {
	AccountID string `json:"account_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (server *Server) updatePasswordAccount(ctx *gin.Context) {
	var req updatePasswordStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.AccountID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.AccountID != authPayload.AccountInfo.AccountId || account.Role != authPayload.AccountInfo.Role {
		err := errors.New("account doesn't belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.UpdatePasswordAccountParams{
		AccountID: req.AccountID,
		Password:  hashedPassword,
		UpdatedAt: time.Now(),
	}

	updated_account, err := server.store.UpdatePasswordAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newAccountResponse(updated_account, nil )
	ctx.JSON(http.StatusOK, rsp)

}
