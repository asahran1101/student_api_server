package constants

const (
	//Endpoint Routes
	AllStudentsRoute          = "/students"
	OneStudentWithRollNoRoute = "/students/:rollNo"

	//Database Queries
	DeleteStudentByRollNoQuery = "DELETE FROM students WHERE rollNo = $1"
	InsertStudentQuery         = "INSERT INTO students(Name, GuardianName, Address, ContactNo, EmailID) VALUES ($1, $2, $3, $4, $5) RETURNING RollNo"
	SelectAllStudentsQuery     = "SELECT * FROM students ORDER BY rollNo ASC"
	SelectStudentByRollNoQuery = "SELECT name, guardianName, address, contactNo, emailId FROM students WHERE rollNo = $1"
	UpdateStudentByRollNoQuery = "UPDATE students SET name = $1, guardianName = $2, address = $3, contactNo = $4, emailId = $5 WHERE rollNo = $6 RETURNING rollNo"
)
