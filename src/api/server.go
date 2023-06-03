package api

import (
	"github.com/gin-gonic/gin"
	db "shopping-cart/src/db/sqlc"
	"shopping-cart/src/util"
)

// Server to serve the RestAPI
type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) *Server {
	server := &Server{
		config: config,
		store:  store,
		router: gin.Default(),
	}

	return server
}

// Start to run the RestAPI
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
