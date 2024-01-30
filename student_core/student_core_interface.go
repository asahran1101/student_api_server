package student_core

import (
	"github.com/gin-gonic/gin"
)

type StudentCoreInterface interface {
	CreateStudent(context *gin.Context)
	DeleteStudent(context *gin.Context)
	GetStudentByRollNo(context *gin.Context)
	GetAllStudents(context *gin.Context)
	UpdateStudent(context *gin.Context)
}
