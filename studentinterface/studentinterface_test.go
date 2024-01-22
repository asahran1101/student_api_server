package studentinterface

import (
	"errors"
	"testing"

	"example.com/api/models"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	t.Run("Delete - Success", func(t *testing.T) {
		studentClient := &StudentIOClient{}

		student := &models.Student{
			Name:         "John Doe",
			GuardianName: "Jane Doe",
			Address:      "123 Main St",
			ContactNo:    "123-456-7890",
			EmailID:      "john.doe@example.com",
		}

		studentClient.Save(student)
		err := studentClient.Delete(student.RollNo)
		assert.NoError(t, err)
		retrievedStudent, err := studentClient.GetStudentByRollNo(student.RollNo)
		assert.Error(t, err)
		assert.Nil(t, retrievedStudent)
		assert.Equal(t, errors.New("Couldn't find the student with the mentioned roll number in the database."), err)
	})

	t.Run("Delete - Not Found", func(t *testing.T) {
		studentClient := &StudentIOClient{}
		err := studentClient.Delete(1)
		assert.Error(t, err)
		assert.Equal(t, errors.New("Couldn't find the student with the mentioned roll number in the database."), err)
	})
}

func TestGetStudentByRollNo(t *testing.T) {
	t.Run("Get Student By Roll No - Success", func(t *testing.T) {
		studentClient := &StudentIOClient{}

		student := &models.Student{
			Name:         "John Doe",
			GuardianName: "Jane Doe",
			Address:      "123 Main St",
			ContactNo:    "123-456-7890",
			EmailID:      "john.doe@example.com",
		}

		studentClient.Save(student)
		retrievedStudent, err := studentClient.GetStudentByRollNo(student.RollNo)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedStudent)
		assert.Equal(t, student, retrievedStudent)
	})

	t.Run("GetStudentByRollNo - Not Found", func(t *testing.T) {
		studentClient := &StudentIOClient{}
		retrievedStudent, err := studentClient.GetStudentByRollNo(1)
		assert.Error(t, err)
		assert.Nil(t, retrievedStudent)
		assert.Equal(t, errors.New("Couldn't find the student with the mentioned roll number in the database."), err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update - Success", func(t *testing.T) {
		studentClient := &StudentIOClient{}
		student := &models.Student{
			Name:         "John Doe",
			GuardianName: "Jane Doe",
			Address:      "123 Main St",
			ContactNo:    "123-456-7890",
			EmailID:      "john.doe@example.com",
		}

		studentClient.Save(student)

		updatedStudent := models.Student{
			RollNo:       student.RollNo,
			Name:         "Updated Name",
			GuardianName: "Updated Guardian",
			Address:      "Updated Address",
			ContactNo:    "Updated Contact",
			EmailID:      "updated.email@example.com",
		}

		err := studentClient.Update(updatedStudent)
		assert.NoError(t, err)
		retrievedStudent, err := studentClient.GetStudentByRollNo(student.RollNo)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedStudent)
		assert.Equal(t, updatedStudent, *retrievedStudent)
	})

	t.Run("Update - Not Found", func(t *testing.T) {
		studentClient := &StudentIOClient{}

		updatedStudent := models.Student{
			RollNo:       1,
			Name:         "Updated Name",
			GuardianName: "Updated Guardian",
			Address:      "Updated Address",
			ContactNo:    "Updated Contact",
			EmailID:      "updated.email@example.com",
		}

		err := studentClient.Update(updatedStudent)
		assert.Error(t, err)
		assert.Equal(t, errors.New("Couldn't find the student with the mentioned roll number in the database."), err)
	})
}
