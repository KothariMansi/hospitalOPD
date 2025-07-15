package api

import (
	db "github.com/KothariMansi/hospitalOPD/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serve http request for our hospital OPD service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// New Server create a new HTTP server and create routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Add route to router
	router.POST("/clients", server.createClient)
	router.GET("/clients/:id", server.getClient)
	router.GET("/clients", server.ListClients)
	router.DELETE("/clients/:id", server.DeleteClient)
	router.PATCH("/clients", server.DeleteClient)

	server.router = router
	return server
}

// Start run the HTTP server on specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
