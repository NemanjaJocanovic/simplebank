package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nemanjaj/simplebank/db/sqlc"
)

// Server servest HTTP requests to our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// New Server creates a new HTTP server and set up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to trouter
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
