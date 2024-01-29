package student_core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"example.com/api/db"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	t.Run("Create Student - Success", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		studentData := models.Student{
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		studentDataResponse := models.Student{
			RollNo:       1,
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		gomock.InOrder(
			mockDbClient.EXPECT().Insert(&studentData).Return(
				&studentDataResponse,
				nil,
			),
		)

		jsonData, err := json.Marshal(studentData)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response)
		context.Request, err = http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.CreateStudent(context)
		assert.Equal(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Student was registered on the server.", result["message"])
		assert.NotNil(t, result["student"])
	})

	t.Run("Create Student - Missing Field", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		studentData := models.Student{
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
		}

		jsonData, err := json.Marshal(studentData)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response)
		context.Request, err = http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.CreateStudent(context)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Could not parse request body. Check if the request is missing a required field.", result["message"])
		assert.NotNil(t, result["error"])
	})

	t.Run("Create Student - Server Error", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		studentData := models.Student{
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		gomock.InOrder(
			mockDbClient.EXPECT().Insert(&studentData).Return(
				nil,
				fmt.Errorf("Could not register the student due to internal server error."),
			),
		)

		jsonData, err := json.Marshal(studentData)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response)
		context.Request, err = http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.CreateStudent(context)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Server could not register the student in the database.", result["message"])
		assert.NotNil(t, result["error"])
	})
}

func TestDeleteStudent(t *testing.T) {
	t.Run("Delete Student - Success", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		studentData := models.Student{
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		studentDataResponse := models.Student{
			RollNo:       1,
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		gomock.InOrder(
			mockDbClient.EXPECT().Insert(&studentData).Return(
				&studentDataResponse,
				nil,
			),
			mockDbClient.EXPECT().DeleteStudentByRollNo(studentDataResponse.RollNo).Return(
				nil,
			),
		)

		jsonData, err := json.Marshal(studentData)
		assert.NoError(t, err)

		response1 := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response1)
		context.Request, err = http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.CreateStudent(context)
		assert.Equal(t, http.StatusOK, response1.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response1.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Student was registered on the server.", result["message"])
		assert.NotNil(t, result["student"])

		response2 := httptest.NewRecorder()
		context, _ = gin.CreateTestContext(response2)
		context.Request, err = http.NewRequest("DELETE", fmt.Sprintf("/students/%d", studentDataResponse.RollNo), nil)
		context.Params = []gin.Param{
			{
				Key:   "rollNo",
				Value: strconv.Itoa(studentDataResponse.RollNo),
			},
		}
		assert.NoError(t, err)

		studentClient.DeleteStudent(context)
		assert.Equal(t, http.StatusOK, response2.Code)

		err = json.NewDecoder(response2.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Student has been deleted.", result["message"])
	})

	t.Run("Delete Student - Server Error", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		rollNo := 1

		gomock.InOrder(
			mockDbClient.EXPECT().DeleteStudentByRollNo(rollNo).Return(
				fmt.Errorf("Student with the mentioned roll number does not exist."),
			),
		)

		response := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response)
		var err error
		context.Request, err = http.NewRequest("DELETE", fmt.Sprintf("/students/%d", rollNo), nil)
		context.Params = []gin.Param{
			{
				Key:   "rollNo",
				Value: strconv.Itoa(rollNo),
			},
		}
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.DeleteStudent(context)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Could not delete the student.", result["message"])
		assert.NotNil(t, result["error"])
	})
}

func TestGetAllStudents(t *testing.T) {
	t.Run("Get Students - Success", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		studentData := []models.Student{
			{
				Name:         "Aniket Sahran",
				GuardianName: "Rajesh Sahran",
				Address:      "Gurgaon",
				ContactNo:    "8700161082",
				EmailID:      "asahran@cloudera.com",
			},
			{
				Name:         "Aniket Sahran",
				GuardianName: "Rajesh Sahran",
				Address:      "Gurgaon",
				ContactNo:    "8700161082",
				EmailID:      "asahran@cloudera.com",
			},
		}

		gomock.InOrder(
			mockDbClient.EXPECT().SelectAllStudents().Return(
				studentData,
				nil,
			),
		)

		response := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response)
		var err error
		context.Request, err = http.NewRequest("GET", "/students", nil)
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.GetAllStudents(context)
		assert.Equal(t, http.StatusOK, response.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Fetched the students", result["message"])
		assert.NotNil(t, result["students"])
	})

	t.Run("Get Students - Server Error", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		gomock.InOrder(
			mockDbClient.EXPECT().SelectAllStudents().Return(
				nil,
				fmt.Errorf("Could not fetch the student details due to internal server error."),
			),
		)

		response := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response)
		var err error
		context.Request, err = http.NewRequest("GET", "/students", nil)
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.GetAllStudents(context)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Could not fetch the student details.", result["message"])
		assert.NotNil(t, result["error"])
	})
}

func TestGetStudentByRollNo(t *testing.T) {
	t.Run("Get Student - Success", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		studentData := models.Student{
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		studentDataResponse := models.Student{
			RollNo:       1,
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		gomock.InOrder(
			mockDbClient.EXPECT().Insert(&studentData).Return(
				&studentDataResponse,
				nil,
			),
			mockDbClient.EXPECT().SelectStudentByRollNo(studentDataResponse.RollNo).Return(
				&studentDataResponse,
				nil,
			),
		)

		jsonData, err := json.Marshal(studentData)
		assert.NoError(t, err)

		response1 := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response1)
		context.Request, err = http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.CreateStudent(context)
		assert.Equal(t, http.StatusOK, response1.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response1.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Student was registered on the server.", result["message"])
		assert.NotNil(t, result["student"])

		response2 := httptest.NewRecorder()
		context, _ = gin.CreateTestContext(response2)
		context.Request, err = http.NewRequest("GET", fmt.Sprintf("/students/%d", studentDataResponse.RollNo), nil)
		context.Params = []gin.Param{
			{
				Key:   "rollNo",
				Value: strconv.Itoa(studentDataResponse.RollNo),
			},
		}
		assert.NoError(t, err)

		studentClient.GetStudentByRollNo(context)
		assert.Equal(t, http.StatusOK, response2.Code)

		err = json.NewDecoder(response2.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Fetched the student.", result["message"])
		assert.NotNil(t, result["student"])
	})

	t.Run("Get Student - Server Error", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)
		var rollNo int

		gomock.InOrder(
			mockDbClient.EXPECT().SelectStudentByRollNo(rollNo).Return(
				nil,
				fmt.Errorf("Could not fetch the student details due to internal server error."),
			),
		)

		response := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response)
		var err error
		context.Request, err = http.NewRequest("GET", fmt.Sprintf("/students/%d", rollNo), nil)
		context.Params = []gin.Param{
			{
				Key:   "rollNo",
				Value: strconv.Itoa(rollNo),
			},
		}
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.GetStudentByRollNo(context)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Could not fetch the student details.", result["message"])
		assert.NotNil(t, result["error"])
	})
}

func TestUpdateStudent(t *testing.T) {
	t.Run("Update Student - Success", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		studentData := models.Student{
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		studentDataResponse := models.Student{
			RollNo:       0,
			Name:         "Aniket Sahran",
			GuardianName: "Rajesh Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		updatedStudentData := models.Student{
			RollNo:       studentData.RollNo,
			Name:         "Aniket Sahran",
			GuardianName: "Sushma Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		gomock.InOrder(
			mockDbClient.EXPECT().Insert(&studentData).Return(
				&studentDataResponse,
				nil,
			),
			mockDbClient.EXPECT().UpdateStudentByRollNo(&updatedStudentData).Return(
				&updatedStudentData,
				nil,
			),
		)

		jsonData, err := json.Marshal(studentData)
		assert.NoError(t, err)

		response1 := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response1)
		context.Request, err = http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.CreateStudent(context)
		assert.Equal(t, http.StatusOK, response1.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response1.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Student was registered on the server.", result["message"])
		assert.NotNil(t, result["student"])

		jsonData, err = json.Marshal(updatedStudentData)
		assert.NoError(t, err)

		response2 := httptest.NewRecorder()
		context, _ = gin.CreateTestContext(response2)
		context.Request, err = http.NewRequest("PUT", fmt.Sprintf("/students/%d", studentDataResponse.RollNo), bytes.NewBuffer(jsonData))
		context.Params = []gin.Param{
			{
				Key:   "rollNo",
				Value: strconv.Itoa(studentDataResponse.RollNo),
			},
		}
		assert.NoError(t, err)

		studentClient.UpdateStudent(context)
		assert.Equal(t, http.StatusOK, response2.Code)

		err = json.NewDecoder(response2.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Student has been updated.", result["message"])
		assert.NotNil(t, result["student"])
	})

	t.Run("Update Student - Server Error", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		updatedStudentData := models.Student{
			RollNo:       1,
			Name:         "Aniket Sahran",
			GuardianName: "Sushma Sahran",
			Address:      "Gurgaon",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		gomock.InOrder(
			mockDbClient.EXPECT().UpdateStudentByRollNo(&updatedStudentData).Return(
				nil,
				fmt.Errorf("Could not update the student details."),
			),
		)

		jsonData, err := json.Marshal(updatedStudentData)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response)
		context.Request, err = http.NewRequest("PUT", fmt.Sprintf("/students/%d", updatedStudentData.RollNo), bytes.NewBuffer(jsonData))
		context.Params = []gin.Param{
			{
				Key:   "rollNo",
				Value: strconv.Itoa(updatedStudentData.RollNo),
			},
		}
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.UpdateStudent(context)
		assert.Equal(t, http.StatusInternalServerError, response.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Could not update the student.", result["message"])
		assert.NotNil(t, result["error"])
	})

	t.Run("Update Student - Missing field", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockDbClient := db.NewMockDatabaseInterface(mockCtl)

		updatedStudentData := models.Student{
			RollNo:       1,
			Name:         "Aniket Sahran",
			GuardianName: "Sushma Sahran",
			ContactNo:    "8700161082",
			EmailID:      "asahran@cloudera.com",
		}

		jsonData, err := json.Marshal(updatedStudentData)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(response)
		context.Request, err = http.NewRequest("PUT", fmt.Sprintf("/students/%d", updatedStudentData.RollNo), bytes.NewBuffer(jsonData))
		context.Params = []gin.Param{
			{
				Key:   "rollNo",
				Value: strconv.Itoa(updatedStudentData.RollNo),
			},
		}
		assert.NoError(t, err)

		studentClient := NewStudentCoreClient(mockDbClient)
		studentClient.UpdateStudent(context)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Could not parse request body. Check if the request is missing a required field.", result["message"])
		assert.NotNil(t, result["error"])
	})
}
