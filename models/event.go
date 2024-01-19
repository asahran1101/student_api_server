package models

type Student struct {
	RollNo       int
	Name         string `binding:"required"`
	GuardianName string `binding:"required"`
	Address      string `binding:"required"`
	ContactNo    string `binding:"required"`
	EmailID      string `binding:"required"`
	Subjects     []string
	Marks        map[string]float64
}

var students = []Student{}

func (s Student) Save() {
	students = append(students, s)
}

func GetAllStudents() []Student {
	return students
}
