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
	// Client router
	router.POST("/clients", server.createClient)
	router.GET("/clients/:id", server.getClient)
	router.GET("/clients", server.listClients)
	router.DELETE("/clients/:id", server.deleteClient)
	router.PATCH("/clients", server.updateClient)
	router.GET("/clients/count", server.countClients)
	router.GET("/clients/search", server.searchClientsByName)

	// User router
	router.POST("/users", server.createUser)

	// Hospital router
	router.POST("/hospitals", server.createHospital)
	router.GET("/hospitals/:id", server.getHospital)
	router.GET("/hospitals", server.listHospital)
	router.PATCH("/hospitals", server.updateHospital)
	router.DELETE("/hospitals/:id", server.deleteHospital)
	router.GET("/hospitals/count", server.countHospitals)
	router.GET("/hospitals/search/name", server.searchHospitalsByName)
	router.GET("/hospitals/search/location", server.searchHospitalsByLocation)

	// Speciality router
	router.POST("/specialities", server.createSpeciality)
	router.GET("/specialities/:id", server.getSpeciality)
	router.GET("/specialities", server.listSpecialities)
	router.DELETE("/specialities/:id", server.deleteSpeciality)
	router.PATCH("/specialities", server.updateSpeciality)
	router.GET("/specialities/count", server.countSpecialites)
	router.GET("/specialities/search", server.searchSpecialitiesByName)

	// CheckUp Time router
	router.POST("/checkuptimes", server.createCheckUpTime)
	router.GET("/checkuptimes/:id", server.getCheckUpTime)
	router.GET("/checkuptimes", server.listCheckUpTime)
	router.PATCH("/checkuptimes", server.updateCheckUpTime)
	router.DELETE("/checkuptimes/:id", server.deleteCheckUpTime)

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
