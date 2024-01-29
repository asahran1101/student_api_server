package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"example.com/api/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DatabaseInterface interface {
	Insert(s *models.Student) (*models.Student, error)
	DeleteStudentByRollNo(rollNo int) error
	SelectAllStudents() ([]models.Student, error)
	SelectStudentByRollNo(rollNo int) (*models.Student, error)
	UpdateStudentByRollNo(s *models.Student) (*models.Student, error)
}

type DatabaseClient struct {
	Db *sql.DB
}

func New() DatabaseClient {
	var dbClient DatabaseClient
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := "postgres"
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", err)
		panic(err)
	} else {
		dbClient.Db = db
		fmt.Println("Successfully connected to database!")
	}
	return dbClient
}

func (DbClient DatabaseClient) Insert(s *models.Student) (*models.Student, error) {
	err := DbClient.Db.QueryRow(
		insertStudentQuery,
		s.Name,
		s.GuardianName,
		s.Address,
		s.ContactNo,
		s.EmailID,
	).Scan(&s.RollNo)

	return s, err
}

func (DbClient DatabaseClient) DeleteStudentByRollNo(rollNo int) error {
	_, err := DbClient.Db.Exec(
		deleteStudentByRollNoQuery,
		rollNo,
	)

	return err
}

func (DbClient DatabaseClient) SelectAllStudents() ([]models.Student, error) {
	rows, err := DbClient.Db.Query(
		selectAllStudentsQuery,
	)
	defer rows.Close()

	if err != nil {
		return nil, errors.New("Could not fetch the student details due to internal server error.")
	}

	var students = []models.Student{}
	var student models.Student

	for rows.Next() {
		var rollNo int
		var name, guardianName, address, contactNo, emailId string
		err := rows.Scan(&rollNo, &name, &guardianName, &address, &contactNo, &emailId)

		if err != nil {
			return nil, errors.New("Could not fetch the student details due to internal server error.")
		}

		student = models.Student{
			RollNo:       rollNo,
			Name:         name,
			GuardianName: guardianName,
			Address:      address,
			ContactNo:    contactNo,
			EmailID:      emailId,
		}

		students = append(students, student)
	}

	if err != nil {
		return nil, errors.New("Could not fetch the student details due to internal server error.")
	}

	return students, nil
}

func (DbClient DatabaseClient) SelectStudentByRollNo(rollNo int) (*models.Student, error) {
	row, err := DbClient.Db.Query(
		selectStudentByRollNoQuery,
		rollNo,
	)

	if err != nil {
		return nil, errors.New("Could not fetch the student details due to internal server error.")
	}

	if !row.Next() {
		return nil, errors.New("Student with the mentioned roll number does not exist.")
	}

	defer row.Close()
	var s models.Student
	s.RollNo = rollNo
	err = row.Scan(&s.Name, &s.GuardianName, &s.Address, &s.ContactNo, &s.EmailID)

	if err != nil {
		return nil, errors.New("Could not fetch the student details due to internal server error.")
	}

	return &s, nil
}

func (DbClient DatabaseClient) UpdateStudentByRollNo(s *models.Student) (*models.Student, error) {
	var rollNo int
	err := DbClient.Db.QueryRow(
		updateStudentByRollNoQuery,
		s.Name,
		s.GuardianName,
		s.Address,
		s.ContactNo,
		s.EmailID,
		s.RollNo,
	).Scan(&rollNo)

	if err != nil {
		return nil, errors.New("Could not update the student details due to internal server error.")
	}

	s.RollNo = rollNo
	return s, err
}
