package iointerface

import (
	"errors"

	"example.com/api/models"
)

var students = []models.Student{}

func GetAllStudents() []models.Student {
	return students
}

func GetStudentByRollNo(rollNo int64) (*models.Student, error) {
	for _, student := range students {
		if student.RollNo == rollNo {
			return &student, nil
		}
	}

	return nil, errors.New("Couldn't find the student with the mentioned roll number in the database.")
}

func Save(student models.Student) {
	students = append(students, student)
}

func Update(updatedStudent models.Student) error {
	for studentInd, student := range students {
		if student.RollNo == updatedStudent.RollNo {
			students[studentInd] = updatedStudent
			return nil
		}
	}

	return errors.New("Couldn't find the student with the mentioned roll number in the database.")
}

func Delete(rollNo int64) error {
	for studentInd, student := range students {
		if student.RollNo == rollNo {
			students = append(students[:studentInd], students[studentInd+1:]...)
			return nil
		}
	}

	return errors.New("Couldn't find the student with the mentioned roll number in the database.")
}
