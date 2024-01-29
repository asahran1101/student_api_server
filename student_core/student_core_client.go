package student_core

import (
	"net/http"
	"strconv"

	"example.com/api/db"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
)

type StudentCoreClient struct {
	DbClient db.DatabaseInterface
}

func NewStudentCoreClient(dbClient db.DatabaseInterface) *StudentCoreClient {
	return &StudentCoreClient{
		DbClient: dbClient,
	}
}

func (s *StudentCoreClient) CreateStudent(context *gin.Context) {
	var student models.Student
	err := context.ShouldBindJSON(&student)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request body. Check if the request is missing a required field.",
			"error":   err.Error(),
		})

		return
	}

	student1, err := s.DbClient.Insert(&student)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server could not register the student in the database.",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Student was registered on the server.",
		"student": student1,
	})
}

func (s *StudentCoreClient) DeleteStudent(context *gin.Context) {
	rollNo, err := strconv.Atoi(context.Param("rollNo"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not find the end point. Check if the url is missing some parameter.",
			"error":   err.Error(),
		})

		return
	}

	err = s.DbClient.DeleteStudentByRollNo(rollNo)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete the student.",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Student has been deleted.",
	})
}

func (s *StudentCoreClient) GetAllStudents(context *gin.Context) {
	students, err := s.DbClient.SelectAllStudents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the student details.",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":  "Fetched the students",
		"students": students,
	})
}

func (s *StudentCoreClient) GetStudentByRollNo(context *gin.Context) {
	rollNo, err := strconv.Atoi(context.Param("rollNo"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not find the end point. Check if the url is missing some parameter.",
			"error":   err.Error(),
		})

		return
	}

	student, err := s.DbClient.SelectStudentByRollNo(rollNo)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the student details.",
			"error":   err.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Fetched the student.",
		"student": student,
	})
}

func (s *StudentCoreClient) UpdateStudent(context *gin.Context) {
	rollNo, err := strconv.Atoi(context.Param("rollNo"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not find the end point. Check if the url is missing some parameter.",
			"error":   err.Error(),
		})

		return
	}

	var updatedStudent models.Student
	err = context.ShouldBindJSON(&updatedStudent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request body. Check if the request is missing a required field.",
			"error":   err.Error(),
		})

		return
	}

	updatedStudent.RollNo = rollNo
	student, err := s.DbClient.UpdateStudentByRollNo(&updatedStudent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update the student.",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Student has been updated.",
		"student": student,
	})
}
