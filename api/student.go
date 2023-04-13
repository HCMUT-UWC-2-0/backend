package api

import (
	"net/http"
	"time"

	db "github.com/DV-Lab/zuni-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createStudentRequest struct {
	StudentID           string        `json:"student_id" binding:"required"`
	Password            string        `json:"password" binding:"required,min=6"`
	PublicKey           string        `json:"public_key"  binding:"required"`
	EncryptedPrivateKey string        `json:"encrypted_private_key"  binding:"required"`
	CitizenID           string        `json:"citizen_id" binding:"required,min=6"`
	Name                string        `json:"name" binding:"required"`
	Gender              db.GenderType `json:"gender" binding:"required"`
	DateOfBirth         string        `json:"date_of_birth" binding:"required,date"`
	PlaceOfBirth        string        `json:"place_of_birth" binding:"required"`
	Class               string        `json:"class" binding:"required"`
	Department          string        `json:"department" binding:"required"`
	TimeOfTraining      int32         `json:"time_of_training" binding:"required"`
	FormOfTraining      string        `json:"form_of_training" binding:"required"`
}

type studentResponse struct {
	StudentID      string        `json:"student_id"`
	CitizenID      string        `json:"citizen_id"`
	Name           string        `json:"name"`
	Gender         db.GenderType `json:"gender"`
	DateOfBirth    time.Time     `json:"date_of_birth"`
	PlaceOfBirth   string        `json:"place_of_birth"`
	Class          string        `json:"class"`
	Department     string        `json:"department"`
	TimeOfTraining int32         `json:"time_of_training"`
	FormOfTraining string        `json:"form_of_training"`
}

func newStudentResponse(student db.Student) studentResponse {
	return studentResponse{
		StudentID:      student.StudentID,
		CitizenID:      student.CitizenID,
		Name:           student.Name,
		Gender:         student.Gender,
		DateOfBirth:    student.DateOfBirth,
		PlaceOfBirth:   student.PlaceOfBirth,
		Class:          student.Class,
		Department:     student.Department,
		TimeOfTraining: student.TimeOfTraining,
		FormOfTraining: student.FormOfTraining,
	}
}

func (server *Server) createStudent(ctx *gin.Context) {
	var req createStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg:= db.InsertStudentTxParams{
		StudentID: req.StudentID,
		Password:req.Password,
		PublicKey:req.PublicKey,
		EncryptedPrivateKey:req.EncryptedPrivateKey,
		CitizenID:req.CitizenID,
		Name:req.Name,
		Gender:req.Gender,
		DateOfBirth:req.DateOfBirth,
		PlaceOfBirth:req.PlaceOfBirth,
		Class:req.Class,
		Department:req.Department,
		TimeOfTraining:req.TimeOfTraining,
		FormOfTraining:req.FormOfTraining,
	}

	result, err := server.store.InsertStudentTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	rsp := (result)
	ctx.JSON(http.StatusOK, rsp)
}

// // List Customer

// type listCustomerRequest struct {
// 	PageID   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
// }

// func (server *Server) listCustomers(ctx *gin.Context) {
// 	var req listCustomerRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	arg := db.ListCustomersParams{
// 		Limit:  req.PageSize,
// 		Offset: (req.PageID - 1) * req.PageSize,
// 	}
// 	customers, err := server.store.ListCustomers(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	rsp := make([]customerResponse, req.PageSize)
// 	for index, value := range customers {
// 		rsp[index] = newCustomerResponse(value)
// 	}
// 	ctx.JSON(http.StatusOK, rsp)"student_id": 1911072,
// "password": "secret",

// //

// //DELETE

// type deleteCustomerRequest struct {
// 	Username string `form:"d_username" binding:"required,min=1"`
// }

// func (server *Server) deleteCustomer(ctx *gin.Context) {
// 	var req deleteCustomerRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	err := server.store.DeleteCustomer(ctx, req.Username)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, req.Username)
// }

// type updateInfoCustomerRequest struct {
// 	Username string `json:"username" binding:"required,alphanum"`
// 	Sex      string `json:"sex"`
// 	Dob      string `json:"dob"`
// 	Phone    string `json:"phone"`
// 	Email    string `json:"email"`
// }

// func (server *Server) updateInfoCustomer(ctx *gin.Context) {
// 	var req updateInfoCustomerRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	time, err := util.MapStringToNullTime(req.Dob)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	arg := db.UpdateInfoCustomerParams{
// 		Username: req.Username,
// 		Sex:      util.MapStringToNullString(req.Sex),
// 		Dob:      time,
// 		Phone:    util.MapStringToNullString(req.Phone),
// 		Email:    util.MapStringToNullString(req.Email),
// 	}

// 	customer, err := server.store.UpdateInfoCustomer(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	rsp := newCustomerResponse(customer)
// 	ctx.JSON(http.StatusOK, rsp)

// }

// type updatePasswordCustomerRequest struct {
// 	Username string `json:"username" binding:"required,alphanum"`
// 	Password string `json:"password"`
// }

// func (server *Server) updatePasswordCustomer(ctx *gin.Context) {
// 	var req updatePasswordCustomerRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	hashedPassword, err := util.HashPassword(req.Password)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	arg := db.UpdatePasswordCustomerParams{
// 		Username:          req.Username,
// 		HashedPassword:    hashedPassword,
// 		PasswordChangedAt: time.Now(),
// 	}

// 	customer, err := server.store.UpdatePasswordCustomer(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	rsp := newCustomerResponse(customer)
// 	ctx.JSON(http.StatusOK, rsp)

// }
