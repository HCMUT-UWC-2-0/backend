package api

import (
	"fmt"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/util"
	"github.com/HCMUT-UWC-2-0/backend/token"
	"github.com/gin-gonic/gin"

    "github.com/gin-contrib/cors"
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
	router.Use(cors.Default())
	// router.POST("/api/admin/create", server.createAdmin)
	router.POST("/auth/login", server.loginAccount)

	// TODO: this api need authentication + authorization
	// router.POST("/api/degree/create", server.createDegree)



	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	server.router = router
}
