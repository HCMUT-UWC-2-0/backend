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
	"github.com/lib/pq"
)

type createDegreeRequest struct {
	StudentID      string `json:"student_id" binding:"required"`
	Specialization string `json:"specialization" binding:"required"`
	DecisionNumber string `json:"decision_number" binding:"required"`
	DateOfIssue    string `json:"date_of_issue" binding:"required,date"`
}

type degreeResponse struct {
	ID             int64     `json:"id"`
	StudentID      string    `json:"student_id"`
	Specialization string    `json:"specialization"`
	DecisionNumber string    `json:"decision_number"`
	DateOfIssue    time.Time `json:"date_of_issue"`
}

func newDegreeResponse(degree db.Degree) degreeResponse {
	return degreeResponse{
		ID:             degree.ID,
		StudentID:      degree.StudentID,
		Specialization: degree.Specialization,
		DecisionNumber: degree.DecisionNumber,
		DateOfIssue:    degree.DateOfIssue,
	}
}

func (server *Server) createDegree(ctx *gin.Context) {
	var req createDegreeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: Authentication + Authorization

	dateOfIssue, err := util.MapStringToTime(req.DateOfIssue)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreteDegreeParams{
		StudentID:      req.StudentID,
		Specialization: req.Specialization,
		DecisionNumber: req.DecisionNumber,
		DateOfIssue:    dateOfIssue,
	}

	degree, err := server.store.CreteDegree(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newDegreeResponse(degree)
	ctx.JSON(http.StatusOK, rsp)
}

// Get certificate by citizen id
type getDegreeRequest struct {
	ID        int64  `json:"id" binding:"required"`
	StudentID string `json:"student_id" binding:"required"`
}

func (server *Server) getDegree(ctx *gin.Context) {
	var req getDegreeRequest
	//========SHOULD BIND URI===========
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	student, err := server.store.GetStudent(ctx, req.StudentID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if student.StudentID != authPayload.AccountInfo.AccountId || authPayload.AccountInfo.Role != db.RoleTypeSTUDENT {
		err := errors.New("account doesn't belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.GetDegreeParams{
		ID:        req.ID,
		StudentID: req.StudentID,
	}

	degree, err := server.store.GetDegree(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newDegreeResponse(degree))
}

// Delete Degree

type deleteDegreeRequest struct {
	ID        int64  `json:"id" binding:"required"`
	AdminID   string `json:"admin_id" binding:"required"`
	StudentID string `json:"student_id" binding:"required"`
}

func (server *Server) deleteDegree(ctx *gin.Context) {
	var req deleteDegreeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	admin, err := server.store.GetAdmin(ctx, req.AdminID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.AccountInfo.Role != db.RoleTypeADMIN || admin.AdminID != authPayload.AccountInfo.AccountId {
		err := errors.New("account doesn't have right to delete degree")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err = server.store.DeleteDegree(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, req)
}

// Update Degree Info

type updateDegreeInfoRequest struct {
	ID             int64  `json:"id" binding:"required"`
	StudentID      string `json:"student_id" binding:"required"`
	Specialization string `json:"specialization" binding:"required"`
	DecisionNumber string `json:"decision_number" binding:"required"`
	DateOfIssue    string `json:"date_of_issue" binding:"required,date"`
}

func (server *Server) updateDegreeInfo(ctx *gin.Context) {
	var req updateDegreeInfoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	getDegreeArg := db.GetDegreeParams{
		ID:        req.ID,
		StudentID: req.StudentID,
	}

	degree, e := server.store.GetDegree(ctx, getDegreeArg)

	if e != nil {
		if e == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(e))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(e))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if degree.StudentID != authPayload.AccountInfo.AccountId || authPayload.AccountInfo.Role != db.RoleTypeADMIN {
		err := errors.New("account doesn't belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	dateOfIssue, err := util.MapStringToTime(req.DateOfIssue)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateDegreeInfoParams{
		ID:             req.ID,
		Specialization: req.Specialization,
		DecisionNumber: req.DecisionNumber,
		DateOfIssue:    dateOfIssue,
		UpdatedAt:      time.Now(),
	}

	updated_certificate, err := server.store.UpdateDegreeInfo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newDegreeResponse(updated_certificate)
	ctx.JSON(http.StatusOK, rsp)

}
