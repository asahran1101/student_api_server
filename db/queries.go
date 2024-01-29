package db

const (
	deleteStudentByRollNoQuery = "DELETE FROM students WHERE rollNo = $1"
	insertStudentQuery         = "INSERT INTO students(Name, GuardianName, Address, ContactNo, EmailID) VALUES ($1, $2, $3, $4, $5) RETURNING RollNo"
	selectAllStudentsQuery     = "SELECT * FROM students ORDER BY rollNo ASC"
	selectStudentByRollNoQuery = "SELECT name, guardianName, address, contactNo, emailId FROM students WHERE rollNo = $1"
	updateStudentByRollNoQuery = "UPDATE students SET name = $1, guardianName = $2, address = $3, contactNo = $4, emailId = $5 WHERE rollNo = $6 RETURNING rollNo"
)
