package models

type Student struct {
	RollNo       int
	Name         string `binding:"required"`
	GuardianName string `binding:"required"`
	Address      string `binding:"required"`
	ContactNo    string `binding:"required"`
	EmailID      string `binding:"required"`
}
