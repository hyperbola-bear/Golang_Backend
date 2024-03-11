package api

import (
	db "example.com/golang_backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.POST("/api/register", server.addTeacherStudents)
	router.GET("/api/commonstudents", server.getTeachersStudents)
	router.POST("/api/suspend", server.addSuspension)
	router.GET("/api/suspend", server.getSuspendedStudents)
	router.POST("/api/retrievefornotifications", server.getRecipientStudents)

	server.router = router
	return server
}