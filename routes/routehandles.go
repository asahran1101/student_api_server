package routes

import (
	"net/http"
	"strconv"

	"example.com/api/iointerface"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

func CreateStudent(context *gin.Context) {
	var student models.Student
	err := context.ShouldBindJSON(&student)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data. Check if the request is missing a required field.",
		})

		return
	}

	student.RollNo = 1
	iointerface.Save(&student)

	context.JSON(http.StatusCreated, gin.H{
		"message": "Student was registered on the server.",
		"student": student,
	})
}

func DeleteStudent(context *gin.Context) {
	rollNo, err := strconv.ParseInt(context.Param("rollNo"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data. Check if the url is missing some parameter.",
		})

		return
	}

	err = iointerface.Delete(rollNo)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not find the student with the mentioned roll number.",
		})

		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"message": "Student has been deleted.",
	})
}

func GetStudent(context *gin.Context) {
	rollNo, err := strconv.ParseInt(context.Param("rollNo"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data. Check if the url is missing some parameter.",
		})

		return
	}

	student, err := iointerface.GetStudentByRollNo(rollNo)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not find the student with the mentioned roll number.",
		})

		return
	}

	context.JSON(http.StatusOK, student)
}

func GetStudents(context *gin.Context) {
	students := iointerface.GetAllStudents()
	context.JSON(http.StatusOK, students)
}

func UpdateStudent(context *gin.Context) {
	rollNo, err := strconv.ParseInt(context.Param("rollNo"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data. Check if the url is missing some parameter.",
		})

		return
	}

	var updatedStudent models.Student
	updatedStudent.RollNo = rollNo
	err = context.ShouldBindJSON(&updatedStudent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data. Check if the request is missing a required field.",
		})

		return
	}

	err = iointerface.Update(updatedStudent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not find the student with the mentioned roll number.",
		})

		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"message": "Student has been updated.",
	})
}
