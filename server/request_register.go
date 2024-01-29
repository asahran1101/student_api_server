package server

import (
	"example.com/api/db"
	"example.com/api/student_core"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	dbClient := db.New()

	var studentClient student_core.StudentCoreInterface
	studentClient = student_core.New(dbClient)

	server.DELETE(oneStudentWithRollNo, studentClient.DeleteStudent)

	server.GET(allStudents, studentClient.GetAllStudents)
	server.GET(oneStudentWithRollNo, studentClient.GetStudentByRollNo)

	server.POST(allStudents, studentClient.CreateStudent)

	server.PUT(oneStudentWithRollNo, studentClient.UpdateStudent)
}
