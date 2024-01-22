package studentinterface

import (
	"errors"

	"example.com/api/models"
)

type StudentIOInterface interface {
	Delete(rollNo int64) error
	GetAllStudents() []models.Student
	GetStudentByRollNo(rollNo int64) (*models.Student, error)
	Save(student *models.Student)
	Update(updatedStudent models.Student) error
}

type StudentIOClient struct {
	students     []models.Student
	noOfStudents int
}

func (s *StudentIOClient) Delete(rollNo int64) error {
	for studentInd, student := range s.students {
		if student.RollNo == rollNo {
			s.students = append(s.students[:studentInd], s.students[studentInd+1:]...)
			return nil
		}
	}

	return errors.New("Couldn't find the student with the mentioned roll number in the database.")
}

func (s *StudentIOClient) GetAllStudents() []models.Student {
	return s.students
}

func (s *StudentIOClient) GetStudentByRollNo(rollNo int64) (*models.Student, error) {
	for _, student := range s.students {
		if student.RollNo == rollNo {
			return &student, nil
		}
	}

	return nil, errors.New("Couldn't find the student with the mentioned roll number in the database.")
}

func (s *StudentIOClient) Save(student *models.Student) {
	s.noOfStudents++
	student.RollNo = int64(s.noOfStudents)
	s.students = append(s.students, *student)
}

func (s *StudentIOClient) Update(updatedStudent models.Student) error {
	for studentInd, student := range s.students {
		if student.RollNo == updatedStudent.RollNo {
			s.students[studentInd] = updatedStudent
			return nil
		}
	}

	return errors.New("Couldn't find the student with the mentioned roll number in the database.")
}
