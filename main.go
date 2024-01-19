package main

import (
	"net/http"

	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/students", listAllStudents)

	server.POST("/students", createStudent)

	server.Run(":8080")
}

func listAllStudents(context *gin.Context) {
	students := models.GetAllStudents()
	context.JSON(http.StatusOK, students)
}

func createStudent(context *gin.Context) {
	var student models.Student
	err := context.ShouldBindJSON(&student)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data. Check if the request is missing a required field.",
		})

		return
	}

	student.RollNo = 1
	student.Subjects = make([]string, 0, 6)
	student.Marks = make(map[string]float64, 6)

	context.JSON(http.StatusCreated, gin.H{
		"message": "Student was registered on the server.",
		"student": student,
	})
	student.Save()
}
