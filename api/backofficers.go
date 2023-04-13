package api

// import (
// 	"net/http"

// 	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
// 	"github.com/gin-gonic/gin"
// )

// type createAdminRequest struct {
// 	AdminID             string `json:"admin_id" binding:"required"`
// 	Password            string `json:"password" binding:"required,min=6"`
// 	PublicKey           string `json:"public_key"  binding:"required"`
// 	EncryptedPrivateKey string `json:"encrypted_private_key"  binding:"required"`
// 	Name                string `json:"name" binding:"required"`
// }

// type adminResponse struct {
// 	AdminID string `json:"admin_id"`
// 	Name    string `json:"name"`
// }

// func newAdminResponse(admin db.Admin) adminResponse {
// 	return adminResponse{
// 		AdminID: admin.AdminID,
// 		Name:    admin.Name,
// 	}
// }

// func (server *Server) createAdmin(ctx *gin.Context) {
// 	var req createAdminRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.InsertAdminTxParams{
// 		AdminID:             req.AdminID,
// 		Password:            req.Password,
// 		PublicKey:           req.PublicKey,
// 		EncryptedPrivateKey: req.EncryptedPrivateKey,
// 		Name:                req.Name,
// 	}

// 	result, err := server.store.InsertAdminTx(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	rsp := (result)
// 	ctx.JSON(http.StatusOK, rsp)
// }
