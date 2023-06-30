package api

import (
	"github.com/gin-gonic/gin"
	db "shopping-cart/src/db/sqlc"
	"shopping-cart/src/util"
)

// Server to serve the HTTP requests for the RestAPI
type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup rooting
func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
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

	server.router = router
}
