package server

import (
	"github.com/asahran1101/student_api_server/constants"
	"github.com/asahran1101/student_api_server/services/db"
	"github.com/asahran1101/student_api_server/services/student_core"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	dbClient := db.New()

	var studentClient student_core.StudentCoreInterface
	studentClient = student_core.NewStudentCoreClient(dbClient)

	server.DELETE(constants.OneStudentWithRollNoRoute, studentClient.DeleteStudent)

	server.GET(constants.AllStudentsRoute, studentClient.GetAllStudents)
	server.GET(constants.OneStudentWithRollNoRoute, studentClient.GetStudentByRollNo)

	server.POST(constants.AllStudentsRoute, studentClient.CreateStudent)

	server.PUT(constants.OneStudentWithRollNoRoute, studentClient.UpdateStudent)
}
