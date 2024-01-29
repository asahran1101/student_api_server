package student_core

import (
	"testing"

	"example.com/api/db"
	"example.com/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestCreateStudent(t *testing.T) {
	t.Run("Create Student - Success", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()
		mockUser := db.NewMockDatabaseInterface(mockCtl)

		gomock.InOrder(
			mockUser.EXPECT().Insert(&models.Student{
				Name:         "Aniket Sahran",
				GuardianName: "Rajesh Sahran",
				Address:      "Gurgaon",
				ContactNo:    "8700161082",
				EmailID:      "asahran@cloudera.com",
			}).Return(
				gomock.AssignableToTypeOf(&models.Student{}),
				nil,
			),
		)

		context := &gin.Context{}
		student := New(mockUser)

		student.CreateStudent(context)

		//server.POST("/students", CreateStudent)

		// student := models.Student{
		// 	Name:         "Aniket Sahran",
		// 	GuardianName: "Rajesh Sahran",
		// 	Address:      "Gurugram",
		// 	ContactNo:    "9876543210",
		// 	EmailID:      "asahran@cloudera.com",
		// }

		// jsonData, err := json.Marshal(student)
		// assert.NoError(t, err)

		// request, err := http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
		// assert.NoError(t, err)

		// request.Header.Set("Content-Type", "application/json")
		// response := httptest.NewRecorder()
		// server.ServeHTTP(response, request)
		// assert.Equal(t, http.StatusCreated, response.Code)
		// var result map[string]interface{}

		// if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		// 	t.Fatal(err)
		// }

		// assert.Equal(t, "Student was registered on the server.", result["message"])
	})

	// t.Run("Create Student - Missing Field", func(t *testing.T) {
	// 	server := gin.Default()
	// 	server.POST("/students", CreateStudent)

	// 	student := models.Student{
	// 		Name:         "Aniket Sahran",
	// 		GuardianName: "Rajesh Sahran",
	// 		ContactNo:    "9876543210",
	// 		EmailID:      "asahran@cloudera.com",
	// 	}

	// 	jsonData, err := json.Marshal(student)
	// 	assert.NoError(t, err)

	// 	request, err := http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
	// 	assert.NoError(t, err)

	// 	request.Header.Set("Content-Type", "application/json")
	// 	response := httptest.NewRecorder()
	// 	server.ServeHTTP(response, request)

	// 	assert.Equal(t, http.StatusBadRequest, response.Code)

	// 	var result map[string]interface{}
	// 	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	assert.Equal(t, "Could not parse request data. Check if the request is missing a required field.", result["message"])
	// })
}

// func TestDeleteStudent(t *testing.T) {
// 	t.Run("Delete Student - Success", func(t *testing.T) {
// 		server1 := gin.Default()
// 		server1.POST("/students", CreateStudent)

// 		student := models.Student{
// 			Name:         "Aniket Sahran",
// 			GuardianName: "Rajesh Sahran",
// 			Address:      "Gurugram",
// 			ContactNo:    "9876543210",
// 			EmailID:      "asahran@cloudera.com",
// 		}

// 		jsonData, err := json.Marshal(student)
// 		assert.NoError(t, err)

// 		request, err := http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
// 		assert.NoError(t, err)

// 		request.Header.Set("Content-Type", "application/json")
// 		response := httptest.NewRecorder()
// 		server1.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusCreated, response.Code)

// 		var result map[string]interface{}
// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Student was registered on the server.", result["message"])

// 		server2 := gin.Default()
// 		server2.DELETE("/students/:rollNo", DeleteStudent)
// 		request, err = http.NewRequest("DELETE", "/students/1", nil)
// 		assert.NoError(t, err)

// 		response = httptest.NewRecorder()
// 		server2.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusAccepted, response.Code)

// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Student has been deleted.", result["message"])
// 	})

// 	t.Run("Delete Student - Student not found", func(t *testing.T) {
// 		server := gin.Default()
// 		server.DELETE("/students/:rollNo", DeleteStudent)
// 		request, err := http.NewRequest("DELETE", "/students/1", nil)
// 		assert.NoError(t, err)

// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusBadRequest, response.Code)

// 		var result map[string]interface{}
// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Could not find the student with the mentioned roll number.", result["message"])
// 	})
// }

// func TestGetStudent(t *testing.T) {
// 	t.Run("Get Student - Success", func(t *testing.T) {
// 		gin.SetMode(gin.TestMode)
// 		server1 := gin.Default()
// 		server2 := gin.Default()
// 		server1.POST("/students", CreateStudent)
// 		server2.GET("/students/:rollNo", GetStudent)

// 		student := models.Student{
// 			Name:         "Aniket Sahran",
// 			GuardianName: "Rajesh Sahran",
// 			Address:      "Gurugram",
// 			ContactNo:    "9876543210",
// 			EmailID:      "asahran@cloudera.com",
// 		}

// 		jsonData, err := json.Marshal(student)
// 		assert.NoError(t, err)

// 		request, err := http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
// 		assert.NoError(t, err)

// 		request.Header.Set("Content-Type", "application/json")
// 		response := httptest.NewRecorder()
// 		server1.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusCreated, response.Code)

// 		var result map[string]interface{}
// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Student was registered on the server.", result["message"])

// 		request, err = http.NewRequest("GET", "/students/1", nil)
// 		assert.NoError(t, err)

// 		server2.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusCreated, response.Code)

// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Could not find the student with the mentioned roll number.", result["message"])
// 	})

// 	t.Run("Get Student - Student not found", func(t *testing.T) {
// 		gin.SetMode(gin.TestMode)
// 		server := gin.Default()
// 		server.GET("/students/:rollNo", GetStudent)
// 		request, err := http.NewRequest("GET", "/students/1", nil)
// 		assert.NoError(t, err)

// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusBadRequest, response.Code)

// 		var result map[string]interface{}
// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Could not find the student with the mentioned roll number.", result["message"])
// 	})
// }

// func TestGetStudents(t *testing.T) {
// 	t.Run("Get Students - Success", func(t *testing.T) {
// 		server1 := gin.Default()
// 		server1.POST("/students", CreateStudent)

// 		student := models.Student{
// 			Name:         "Aniket Sahran",
// 			GuardianName: "Rajesh Sahran",
// 			Address:      "Gurugram",
// 			ContactNo:    "9876543210",
// 			EmailID:      "asahran@cloudera.com",
// 		}

// 		jsonData, err := json.Marshal(student)
// 		assert.NoError(t, err)

// 		request, err := http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
// 		assert.NoError(t, err)

// 		request.Header.Set("Content-Type", "application/json")
// 		response := httptest.NewRecorder()
// 		server1.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusCreated, response.Code)

// 		var result map[string]interface{}
// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Student was registered on the server.", result["message"])

// 		server2 := gin.Default()
// 		server2.GET("/students", GetStudents)
// 		request, err = http.NewRequest("GET", "/students", nil)
// 		assert.NoError(t, err)

// 		response = httptest.NewRecorder()
// 		server2.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusOK, response.Code)
// 	})
// }

// func TestUpdateStudent(t *testing.T) {
// 	t.Run("Update Student - Success", func(t *testing.T) {
// 		server1 := gin.Default()
// 		server1.POST("/students", CreateStudent)

// 		student := models.Student{
// 			Name:         "Aniket Sahran",
// 			GuardianName: "Rajesh Sahran",
// 			Address:      "Gurugram",
// 			ContactNo:    "9876543210",
// 			EmailID:      "asahran@cloudera.com",
// 		}

// 		jsonData, err := json.Marshal(student)
// 		assert.NoError(t, err)

// 		request, err := http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
// 		assert.NoError(t, err)

// 		request.Header.Set("Content-Type", "application/json")
// 		response := httptest.NewRecorder()
// 		server1.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusCreated, response.Code)

// 		var result map[string]interface{}
// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Student was registered on the server.", result["message"])

// 		server2 := gin.Default()
// 		server2.PUT("/students/:rollNo", UpdateStudent)

// 		student = models.Student{
// 			Name:         "Aniket Sahran",
// 			GuardianName: "Sushma Sahran",
// 			Address:      "Gurugram",
// 			ContactNo:    "9876543210",
// 			EmailID:      "asahran@cloudera.com",
// 		}

// 		jsonData, err = json.Marshal(student)
// 		assert.NoError(t, err)

// 		request, err = http.NewRequest("PUT", "/students/1", bytes.NewBuffer(jsonData))
// 		assert.NoError(t, err)

// 		response = httptest.NewRecorder()
// 		server2.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusBadRequest, response.Code)

// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Could not find the student with the mentioned roll number.", result["message"])
// 	})

// 	t.Run("Update Student - Student not found", func(t *testing.T) {
// 		server := gin.Default()
// 		server.PUT("/students/:rollNo", UpdateStudent)

// 		student := models.Student{
// 			Name:         "Aniket Sahran",
// 			GuardianName: "Rajesh Sahran",
// 			Address:      "Gurugram",
// 			ContactNo:    "9876543210",
// 			EmailID:      "asahran@cloudera.com",
// 		}

// 		jsonData, err := json.Marshal(student)
// 		assert.NoError(t, err)

// 		request, err := http.NewRequest("PUT", "/students/1", bytes.NewBuffer(jsonData))
// 		assert.NoError(t, err)

// 		response := httptest.NewRecorder()
// 		server.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusBadRequest, response.Code)

// 		var result map[string]interface{}
// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Could not find the student with the mentioned roll number.", result["message"])
// 	})

// 	t.Run("Update Student - Missing field", func(t *testing.T) {
// 		server1 := gin.Default()
// 		server1.POST("/students", CreateStudent)

// 		student := models.Student{
// 			Name:         "Aniket Sahran",
// 			GuardianName: "Rajesh Sahran",
// 			Address:      "Gurugram",
// 			ContactNo:    "9876543210",
// 			EmailID:      "asahran@cloudera.com",
// 		}

// 		jsonData, err := json.Marshal(student)
// 		assert.NoError(t, err)

// 		request, err := http.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
// 		assert.NoError(t, err)

// 		request.Header.Set("Content-Type", "application/json")
// 		response := httptest.NewRecorder()
// 		server1.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusCreated, response.Code)

// 		var result map[string]interface{}
// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Student was registered on the server.", result["message"])

// 		server2 := gin.Default()
// 		server2.PUT("/students/:rollNo", UpdateStudent)

// 		updatedStudent := models.Student{
// 			Name:      "Aniket Sahran",
// 			Address:   "Gurugram",
// 			ContactNo: "9876543210",
// 			EmailID:   "asahran@cloudera.com",
// 		}

// 		jsonData, err = json.Marshal(updatedStudent)
// 		assert.NoError(t, err)

// 		request, err = http.NewRequest("PUT", "/students/"+strconv.Itoa(int(student.RollNo)), bytes.NewBuffer(jsonData))
// 		assert.NoError(t, err)

// 		response = httptest.NewRecorder()
// 		server2.ServeHTTP(response, request)
// 		assert.Equal(t, http.StatusBadRequest, response.Code)

// 		err = json.NewDecoder(response.Body).Decode(&result)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Could not parse request data. Check if the request is missing a required field.", result["message"])
// 	})
// }

// func GetTestGinContext() *gin.Context {
// 	gin.SetMode(gin.TestMode)

// 	w := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(w)
// 	ctx.Request = &http.Request{
// 		Header: make(http.Header),
// 	}

// 	return ctx
// }
