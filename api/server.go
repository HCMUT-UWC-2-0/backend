package api

import (
	"fmt"
	"net/http"

	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"
	"github.com/HCMUT-UWC-2-0/backend/token"
	"github.com/HCMUT-UWC-2-0/backend/util"
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
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
	// router.POST("/api/admin/create", server.createAdmin)
	router.POST("/auth/login", server.loginAccount)

	// TODO: this api need authentication + authorization

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/api/workers", server.listAllWorkers)
	authRoutes.GET("/api/vehicles", server.listAllVehicles)
	authRoutes.GET("/api/mcps", server.listAllMCPs)
	authRoutes.GET("/api/routes", server.listAllRoutes)
	authRoutes.POST("/api/tasks", server.createTask)
	authRoutes.GET("/api/tasks/current", server.listAllCurrentTasks)

	server.router = router
}
