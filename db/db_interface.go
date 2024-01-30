package db

import "example.com/api/models"

type DatabaseInterface interface {
	Insert(s *models.Student) (*models.Student, error)
	DeleteStudentByRollNo(rollNo int) error
	SelectAllStudents() ([]models.Student, error)
	SelectStudentByRollNo(rollNo int) (*models.Student, error)
	UpdateStudentByRollNo(s *models.Student) (*models.Student, error)
}
