package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"

	"example.com/api/constants"
	"example.com/api/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DatabaseClient struct {
	Db *sql.DB
}

var once sync.Once
var dbClient DatabaseClient

func New() DatabaseClient {
	once.Do(func() {
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
	})

	return dbClient
}

func (DbClient DatabaseClient) Insert(s *models.Student) (*models.Student, error) {
	err := DbClient.Db.QueryRow(
		constants.InsertStudentQuery,
		s.Name,
		s.GuardianName,
		s.Address,
		s.ContactNo,
		s.EmailID,
	).Scan(&s.RollNo)

	if err != nil {
		return nil, fmt.Errorf("Could not register the student due to internal server error. %v", err)
	}

	return s, nil
}

func (DbClient DatabaseClient) DeleteStudentByRollNo(rollNo int) error {
	_, err := DbClient.SelectStudentByRollNo(rollNo)

	if err != nil {
		return err
	}

	_, err = DbClient.Db.Exec(
		constants.DeleteStudentByRollNoQuery,
		rollNo,
	)

	if err != nil {
		return fmt.Errorf("Could not delete the student. %v", err)
	}

	return nil
}

func (DbClient DatabaseClient) SelectAllStudents() ([]models.Student, error) {
	rows, err := DbClient.Db.Query(
		constants.SelectAllStudentsQuery,
	)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("Could not fetch the student details due to internal server error. %v", err)
	}

	var students = []models.Student{}
	var student models.Student

	if !rows.Next() {
		return nil, fmt.Errorf("No students exist in the database.")
	}

	for rows.Next() {
		var rollNo int
		var name, guardianName, address, contactNo, emailId string
		err := rows.Scan(&rollNo, &name, &guardianName, &address, &contactNo, &emailId)

		if err != nil {
			return nil, fmt.Errorf("Could not fetch the student details due to internal server error. %v", err)
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
		return nil, fmt.Errorf("Could not fetch the student details due to internal server error. %v", err)
	}

	return students, nil
}

func (DbClient DatabaseClient) SelectStudentByRollNo(rollNo int) (*models.Student, error) {
	row, err := DbClient.Db.Query(
		constants.SelectStudentByRollNoQuery,
		rollNo,
	)

	if err != nil {
		return nil, fmt.Errorf("Could not fetch the student details due to internal server error. %v", err)
	}

	if !row.Next() {
		return nil, fmt.Errorf("Student with the mentioned roll number does not exist.")
	}

	defer row.Close()
	var s models.Student
	s.RollNo = rollNo
	err = row.Scan(&s.Name, &s.GuardianName, &s.Address, &s.ContactNo, &s.EmailID)

	if err != nil {
		return nil, fmt.Errorf("Could not fetch the student details due to internal server error. %v", err)
	}

	return &s, nil
}

func (DbClient DatabaseClient) UpdateStudentByRollNo(s *models.Student) (*models.Student, error) {
	var rollNo int
	err := DbClient.Db.QueryRow(
		constants.UpdateStudentByRollNoQuery,
		s.Name,
		s.GuardianName,
		s.Address,
		s.ContactNo,
		s.EmailID,
		s.RollNo,
	).Scan(&rollNo)

	if err != nil {
		return nil, fmt.Errorf("Could not update the student details. %v", err)
	}

	s.RollNo = rollNo
	return s, nil
}
