package api

import (
	"fmt"

	db "github.com/DV-Lab/zuni-backend/db/sqlc"
	"github.com/DV-Lab/zuni-backend/token"
	"github.com/DV-Lab/zuni-backend/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}
	
	server.setupGinDateValidation()
	
	server.setupRouter()
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/api/student/create", server.createStudent)
	router.POST("/api/admin/create", server.createAdmin)
	router.POST("/auth/login", server.loginAccount)

	// TODO: this api need authentication + authorization
	router.POST("/api/degree/create", server.createDegree)



	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/api/account/:account_id", server.getAccount)
	authRoutes.PATCH("/api/account/update/password", server.updatePasswordAccount)

	authRoutes.GET("/api/degree/:citizenID", server.getDegree)
	authRoutes.PATCH("/api/degree/update/info", server.updateDegreeInfo)
	authRoutes.DELETE("/api/degree", server.deleteDegree)
	server.router = router
}
