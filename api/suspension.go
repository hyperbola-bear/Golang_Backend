package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddSuspensionRequest struct {
	Student string `json:"student" binding:"required"`
}

func (server *Server) addSuspension(ctx *gin.Context) {
	var req AddSuspensionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.AddSuspension(context.Background(), req.Student)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "student suspended successfully"})
}

func (server *Server) getSuspendedStudents(ctx *gin.Context) {
	students, err := server.store.GetSuspendedStudents(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, students)
}