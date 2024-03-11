package api

import (
	"context"
	"net/http"
	"strings"

	db "example.com/golang_backend/db/sqlc"
	"github.com/gin-gonic/gin"
)


type AddTeacherStudentsRequest struct {
	Teacher  string   `json:"teacher" binding:"required"`
	Students []string `json:"students" binding:"required"`
}

func (server *Server) addTeacherStudents(ctx *gin.Context) {
	var req AddTeacherStudentsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if (len(req.Students) == 0) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "students array cannot be empty"})
		return
	}
	arg := db.AddTeacherStudentsParams{
		Teacher:  req.Teacher,
		Students: req.Students,
	}
	err := server.store.AddTeacherStudents(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "students added successfully"})
}

// type GetTeacherStudentsRequest struct {
// 	Teacher string `form:"teacher" binding:"required"`
// }

func (server *Server) getTeachersStudents(ctx *gin.Context) {
	// var req GetTeacherStudentsRequest
	teacherEmails := ctx.QueryArray("teacher")
	
	// if err := ctx.ShouldBindQuery(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }
	students, err := server.store.GetTeachersStudents(context.Background(), teacherEmails)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if students.Students == nil{
		ctx.JSON(http.StatusOK, gin.H{"students": []string{}})
		return
	}
	ctx.JSON(http.StatusOK, students)
}

type getRecipientStudentsParams struct {
	Teacher string `json:"teacher" binding:"required"`
	Notification string `json:"notification" binding:"required"`
}

func (server *Server) getRecipientStudents(ctx *gin.Context) {
	var req getRecipientStudentsParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	teacherEmail := req.Teacher
	
	students, err := server.store.GetTeacherStudents(context.Background(), teacherEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	notification := req.Notification
	whitespaceSplit := strings.Fields(notification)
	for _, word := range whitespaceSplit {
		if strings.HasPrefix(word, "@") {
			student := word[1:]
			students = append(students, student)
		}
	}
	suspendedStudents, err := server.store.GetSuspendedStudents(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	suspensionMap := make(map[string]bool)
	for _, student := range suspendedStudents {
		suspensionMap[student] = true
	}
	var finalStudents []string
	for _, student := range students {
		_, ok := suspensionMap[student]
		if  !ok {
			finalStudents = append(finalStudents, student)
			suspensionMap[student] = true
		} 
	}
	ctx.JSON(http.StatusOK, gin.H{"recipients": finalStudents})
}