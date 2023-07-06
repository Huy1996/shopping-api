package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "shopping-cart/src/db/sqlc"
	"shopping-cart/src/token"
	"shopping-cart/src/util"
)

// Server to serve the HTTP requests for the RestAPI
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup rooting
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

// Start to run the RestAPI
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse will handle how gin going to format error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// setupRouter will set up routing in HTTP server
func (server *Server) setupRouter() {
	router := gin.Default()

	// routing
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	server.router = router
}
